package main

import (
    "sync/atomic"
    "fmt"
)

func main() {
    var v atomic.Value
    v.Store("joker")
    fmt.Println(v.Load())    // joker
    v.Store(1)
    fmt.Println(v.Load())  // panic: sync/atomic: store of inconsistently typed value into Value
}

对于v.Store("joker")触发第一次写入，类型就固定为string了，后面的v.Store(1)尝试写入int类型报错，这也是概述里对任意类型加引号的原因


原子操作由底层硬件支持，而锁则由操作系统提供的 API 实现。
若实现相同的功能，前者通常会更有效率，并且更能利用计算机多核的优势。
所以，以后当我们想并发安全的更新一些变量的时候，我们应该优先选择用atomic.Value来实现。

使用规则：
不能用atomic.Value原子值存储nil
我们向原子值存储的第一个值，决定了它今后能且只能存储哪一个类型的值
建议：不要把内部使用的atomic.Value原子值暴露给外界，如果非要暴露也要通过API封装形式，做严格的check


https://blog.csdn.net/u010853261/article/details/103996679
atomic.Value实践以及原理
Store 实现
// Store sets the value of the Value to x.
// All calls to Store for a given Value must use values of the same concrete type.
// Store of an inconsistent type panics, as does Store(nil).
func (v *Value) Store(x interface{}) {
	if x == nil {
		panic("sync/atomic: store of nil value into Value")
	}
	// 转换*Value 为 *ifaceWords
	vp := (*ifaceWords)(unsafe.Pointer(v))
	// 转换要存储的value为(*ifaceWords)
	xp := (*ifaceWords)(unsafe.Pointer(&x))
	for {
		// 原子加载atomic.Value里面当前存储的变量类型
		typ := LoadPointer(&vp.typ)
		// type为空，表示第一次加载
		if typ == nil {
			// Attempt to start first store.
			// Disable preemption so that other goroutines can use
			// active spin wait to wait for completion; and so that
			// GC does not see the fake type accidentally.
			// 当前线程禁止抢占，GC也不会看到这个中间态
			runtime_procPin()
			// 设置类型为中间态
			if !CompareAndSwapPointer(&vp.typ, nil, unsafe.Pointer(^uintptr(0))) {
				//已经处于中间态了。
				runtime_procUnpin()
				continue
			}
			// Complete first store.
			StorePointer(&vp.data, xp.data)
			StorePointer(&vp.typ, xp.typ)
			runtime_procUnpin()
			return
		}
		// 正在第一次Store的中间过程中(也就是中间态)
		if uintptr(typ) == ^uintptr(0) {
			// First store in progress. Wait.
			// Since we disable preemption around the first store,
			// we can wait with active spinning.
			continue
		}
		// First store completed. Check type and overwrite data.
		if typ != xp.typ {
			panic("sync/atomic: store of inconsistently typed value into Value")
		}
		StorePointer(&vp.data, xp.data)
		return
	}
}

现在描述一下大致流程：
1.先把现有值和将要写入的值转换成ifaceWords类型，这样我们下一步就可以得到这两个interface{}的原始类型（typ）和真正的值（data）。
2.进入 一个无限 for 循环。配合CompareAndSwap食用，可以达到乐观锁的功效。
3.通过LoadPointer这个原子操作拿到当前Value中存储的类型。下面根据这个类型的不同，分3种情况处理：
	1.一个Value实例被初始化后，它的typ字段会被设置为指针的零值 nil，所以先判断如果typ是 nil 那就证明这个Value还未被写入过数据。那之后就是一段初始写入的操作：
		1.runtime_procPin()，它可以将一个goroutine死死占用当前使用的P(P-M-G中的processor)，不允许其它goroutine/M抢占, 使得它在执行当前逻辑的时候不被打断，以便可以尽快地完成工作，因为别人一直在等待它。另一方面，在禁止抢占期间，GC 线程也无法被启用，这样可以防止 GC 线程看到一个莫名其妙的指向^uintptr(0)的类型（这是赋值过程中的中间状态）。
		2.使用CAS操作，先尝试将typ设置为^uintptr(0)这个中间状态。如果失败，则证明已经有别的线程抢先完成了赋值操作，那它就解除抢占锁，然后重新回到 for 循环第一步。
		3.如果设置成功，那证明当前线程抢到了这个"乐观锁”，它可以安全的把v设为传入的新值了（19~23行）。注意，这里是先写data字段，然后再写typ字段。因为我们是以typ字段的值作为写入完成与否的判断依据的。
	2.第一次写入还未完成，如果看到 typ字段还是^uintptr(0)这个中间类型，证明刚刚的第一次写入还没有完成，所以它会继续循环，“忙等"到第一次写入完成。
	3.第一次写入已完成（第31行及之后） - 首先检查上一次写入的类型与这一次要写入的类型是否一致，如果不一致则抛出异常。反之，则直接把这一次要写入的值写入到data字段。
这个逻辑的主要思想就是，为了完成多个字段的原子性写入，我们可以抓住其中的一个字段，以它的状态来标志整个原子写入的状态。


Load
func (v *Value) Load() (x interface{}) {
	vp := (*ifaceWords)(unsafe.Pointer(v))
	typ := LoadPointer(&vp.typ)
	if typ == nil || uintptr(typ) == ^uintptr(0) {
		// First store not yet completed.
		return nil
	}
	data := LoadPointer(&vp.data)
	xp := (*ifaceWords)(unsafe.Pointer(&x))
	xp.typ = typ
	xp.data = data
	return
}

读取相对就简单很多了，它有两个分支：
	1.如果当前的typ是 nil 或者^uintptr(0)，那就证明第一次写入还没有开始，或者还没完成，那就直接返回 nil （不对外暴露中间状态）。
	2.否则，根据当前看到的typ和data构造出一个新的interface{}返回出去。
