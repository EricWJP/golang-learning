函数
  函数 与 方法可以重名
  不同类型接受者可以重名
  函数定义无序
  go函数参数不支持默认参数
  return与声明时返回值定义 必须完全匹配：数量、类型 其中 return单句：顺序必须也一致
	func funcname(p int) { return 0,0 } //无接收者
  // 接收者必须有括号
	func (p type) funcname(p int) { return } //无返回值
	func (p type) funcname(p,q int) (r,s int) { return 0,0 }
	func (p type) funcname(p int) (int,int) { return 0,0 }
	不允许函数嵌套
	可以利用匿名函数实现
  递归函数
  	func rec(i int) {
  		if i == 10 {
  			return
  		}
  		rec(i+1)
  		fmt.Print("%d ", i)
  	}
  全局变量：函数外定义
  局部变量：函数内定义 局部覆盖全局
  

  多值返回
    func f(arg int) (int, int) { ... }
    func f() { ... }
  	func (file *File) Write(b []byte) (n int, err error)
  	无错 err 为nil，出错 err 非nil
  命名返回值	return 没有值 会给返回值赋相应类型的nil值
    不具名返回 
      单个可以没有括号
        func f(arg int) int { ... }
      多个 必须有括号
        func f(arg int) (int, int) { ... }
    具名返回必须有括号
    func f(arg int) (i int) { ... }
  	func ReadFull(r Reader, buf []byte) (n int, err error) {
  		for len(buf) > 0 && err == nil {
  			var nr int
  			nr, err = r.Read(buf)
  			n += nr
  			buf = buf[nr:len(buf)]
  		}
  		return
  	}

  延迟代码
  	defer 指定的函数会在函数退出前调用
  	func ReadWrite() bool {
  		file.Open("file")
  		defer file.Close()
  		if failureX {
  			...
  			// defer fmt.Printf("%d ", i)
  		}
  	}
    多个 defer 按照后进先出--LIFO 的顺序执行
    defer 也可以修改返回值

    匿名参数必须有 ()
    defer func() {
      /* ... */
    }()

    defer func(x int) {
      /*...*/
    }(5)

    匿名参数可以访问任何命名的参数
    func f() (ret int) {
      defer func() {
        ret++
      }()
      return 0    //返回 1 而不返回 0
    }

    变参
    func myfunc(arg ...int) {}
    arg 是一个 int类型的slice
    函数符号 也就是被叫做闭包

    func myfunc(arg ...int) {
      myfunc2(arg...)       // 按原样传递
      myfunc2(arg[:2]...)   // 传递部分
    }

    函数作为值
    func main() {
      a := func() {   // 定义一个匿名函数，并且赋值给 a
        println("Hello")
      }
      a()         // 调用函数
      fmt.Printf("%T\n", a)   // 打印 a 的类型， 输出 func()
    }


    函数也可以作为 map 的值
    var xs = map[int]func() int{
      1: func() int { return 10 },
      2: func() int { return 20 },
      3: func() int { return 30 },   //必须有逗号
    }

    函数也可以作为函数的参数

    回调
    func callback(y int, f func(int)) {
      f(y)        // 调用回调函数 f 输入变量 y
    }

    恐慌Panic 和 恢复Recover  内建函数
      异常机制
      作为最后的手段被使用
    Panic 中断原有的控制流程，进入一个令人恐慌的流程。但是延迟函数会正常执行，然后 返回到F调用panic的地方，这一过程继续向上，直到程序崩溃时的所有goroutine返回
    Recover 让进入panic流程中的goroutine恢复过来。仅在延迟函数中有效
      正常执行调用recover 返回nil。 如果goroutine陷入恐慌，调用recover 可以捕获到panic的输入值，并且恢复正常执行
    func throwsPanic(f func()) (b bool) {
      defer func() {
        if x := recover(); x != nil {
          b = true
        }()
        f()
        return 
      }
    }
    定义了一个利用 recover 的 defer 函数。如果当前的 goroutine 产生了 panic， 这个 defer 函数能够发现。当 recover() 返回非 nil 值，设置 b 为 true;
    调用作为参数接收的函数 f
    返回 b 的值。由于 b 是命名返回值，无须指定 b

    func average(xs []float64) (avg float64) {
      sum := 0.0
      switch len(xs) {
      case 0:
        avg = 0
      default:
        for _, v := range xs {  // v只在for循环内有效
          sum += v
        }
        avg = sum / float64(len(x))  //为了除法正常运算，必须将值转换为 float64
      }
      return
    }





















