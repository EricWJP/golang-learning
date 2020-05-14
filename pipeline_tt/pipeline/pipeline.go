package pipeline

import (
	"sync"
)

type Pipeline struct {
	mtx         sync.Mutex
	once        sync.Once
	workerChans []chan struct{}
	prevThd     *pipelineThread
	ChanExit    chan struct{}
	continued   bool
}

type pipelineThread struct {
	sigs     []chan struct{}
	chanExit chan struct{}
	// interrupt SyncFlag
	// setInterrupt func()
	err error
}

// type SyncFlag interface {
// 	Wait()
// 	Chan() <-chan struct{}
// 	Done() bool
// }

// type syncFlag struct {
// 	once sync.once
// 	c chan struct{}
// }

func NewPipeline(workers ...int) *Pipeline {
	if len(workers) < 1 {
		panic("New pipeline need at least one argument")
	}

	workersChan := make([]chan struct{}, len(workers))
	for i := range workersChan {
		workersChan[i] = make(chan struct{}, workers[i])
	}

	prevThd := newPipelineThread(len(workers))
	for _, sig := range prevThd.sigs {
		close(sig)
	}
	close(prevThd.chanExit)

	chanExit := make(chan struct{})
	return &Pipeline{
		workerChans: workersChan,
		prevThd:     prevThd,
		ChanExit:    chanExit,
		continued:   true,
	}
}

func newPipelineThread(length int) *pipelineThread {
	p := &pipelineThread{
		sigs:     make([]chan struct{}, length),
		chanExit: make(chan struct{}),
	}
	// p.setInterrupt, p.interrupt = NewSyncFlag()

	for i := range p.sigs {
		p.sigs[i] = make(chan struct{})
	}
	return p
}

func (p *Pipeline) done() {
	p.once.Do(func() {
		close(p.ChanExit)
		p.continued = false
	})
}

// func NewSyncFlag() (done func(), flag SyncFlag) {
// 	f := &syncFlag{
// 		c: make(chan struct{}),
// 	}
// 	return f.done, f
// }

// func HasClosed(c <-chan struct{}) bool {
// 	select {
// 	case <-c:
// 		return true
// 	default:
// 		return false
// 	}
// }

// func (f *syncFlag) done() {
// 	f.once.Do(func(){
// 		close(f.c)
// 	})
// }

// func (f *syncflag) Wailt() {
// 	<-f.c
// }

// func (f *syncflag) Chan() <-chan struct{} {
// 	return f.c
// }

// func (f *fyncFlag) Done() bool {
// 	return HasClosed(f.c)
// }

func (p *Pipeline) Async(works ...func() error) bool {
	if len(works) != len(p.workerChans) {
		panic("Async: arguments number not matched to NewPipeline(....)")
	}

	p.mtx.Lock()
	// if p.prevThd.interrupt.Done() {
	// 	return false
	// }
	prevThd := p.prevThd
	thisThd := newPipelineThread(len(p.workerChans))
	p.prevThd = thisThd
	p.mtx.Unlock()

	lock := func(idx int) bool {
		select {
		// case <-prevThd.interrupt.Chan():
		// 	return false
		case <-prevThd.sigs[idx]:
		}
		select {
		// case <-prevThd.interrupt.Chan():
		// 	return false
		case p.workerChans[idx] <- struct{}{}:
		}
		return true
	}
	if !(p.continued && lock(0)) {
		// thisThd.setInterrupt()
		<-prevThd.chanExit
		thisThd.err = prevThd.err
		close(thisThd.chanExit)
		return false
	}
	go func() {
		select {
		// case <-prevThd.interrupt.Chan():
		// thisThd.setInterrupt()
		case <-thisThd.chanExit:
		}
	}()
	go func() {
		var err error
		for i, work := range works {
			close(thisThd.sigs[i])
			if work != nil {
				err = work()
			}
			if err != nil || (i+1 < len(works) && !(p.continued && lock(i+1))) {
				// thisThd.setInterrupt()
				p.done()
				break
			}
			<-p.workerChans[i]
		}
		<-prevThd.chanExit
		// if prevThd.interrupt.Done() {
		// 	thisThd.setInterrupt()
		// }
		if prevThd.err != nil {
			thisThd.err = prevThd.err
		} else {
			thisThd.err = err
		}
		close(thisThd.chanExit)
	}()
	return true
}

//等待流水线中所有任务执行完毕或失败，返回第一个错误，如果无错误则返回nil。
func (p *Pipeline) Wait() error {
	p.mtx.Lock()
	lastThd := p.prevThd
	p.mtx.Unlock()
	<-lastThd.chanExit
	return lastThd.err
}
