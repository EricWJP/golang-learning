并发
并行是性能相关的
并发是程序设计相关的
goroutine 是与其他goroutine并行执行的，有着相同地址空间的函数
轻量的，仅比分配占空间多一点点儿消耗
初始时栈是很小的，所以他们也是廉价的，并且随着需要在堆空间上分配和释放

go ready("Tea", 2)   ready() 作为 goroutine 运行

channel 可以与 Unix shell 中的双向管道做类比： 可以通过它发送或者接受值
	这些值只能是特定的类型：channel类型， 必须使用make 创建 channel

ci := make(chan int)		  // 用于发送和接收整数
cs := make(chan string)		  // 用于字符串
cf := make(chan interface{})。// 使用空接口满足各种类型

向channel发送接收 <-
	ci <- 1		// 发送整数 1 到channel ci
	<-ci 		// 从channel ci 接收整数
	i := <-ci 	// 从channel ci 接收整数，并保存到 i 中

select 可以监听 channel 上输入的数据

runtime.GOMAXPROCS(n)  设置goroutine 并行执行的数量 同时运行的CPU的最大数量

ch := make(chan bool) 创建channel时，bool型的无缓冲channel会被创建
	读取 value := <-ch 将会被阻塞，直到有数据接收。任何发送 ch<-5 将会被阻塞，直到数据被读出
	无缓冲channel是在多个goroutine之间同步很棒的工具

Go也允许指定channel的缓冲大小（存储多少元素）。 ch := make(chan bool, 4) 4个元素的bool型channel。
	前4个元素可以无阻塞写入。第五个元素时，代码将会阻塞，直到其他goroutine从channel中读取一些元素，腾出空间



								value==0	无缓冲
ch := make(chan type, value)
								value>0		缓冲value的元素

关闭channel
	x, ok = <-ch

ok 为 true 意味着channel尚未被关闭，同时可以读取数据。
ok 为 false 在这个情况下表示channel被关闭

只读或只写 channel














