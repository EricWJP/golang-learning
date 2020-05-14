go 编译器需要源码非.a临时文件，
	链接时使用的是 
		第三方源码：以最新源码编译生成的.a文件
		标准库：编译生成的.a 文件，即使修改了标准库文件也不会重新编译生成.a文件
import 
	目录路径下只能有一个包
	包名和路径名可以不同，惯例：目录名与包名要相同
	import "fmt"  目录路径名
	import fm "fmt"		fm：指定包名的别名  和  目录路径名

init() 不能被其他函数调用
在main()函数执行前先隐形调用所有的init()函数，且顺序固定
1、对同一个go文件的init()调用顺序是从上到下的
2、对同一个package中不同文件
	先按文件名字符串，逐个调用各个包的依赖包的init()
	再按文件名字符串比较“从小到大”顺序调用各文件中的init()函数
3、import 包，先调用该包依赖的包的init(),在调用该包的init()
4、相同的包只import第一次(所有依赖嵌套中发现的第一个)，调用第一次该包所有的init()

dep 	https://tonybai.com/2017/06/08/first-glimpse-of-dep/
	brew install dep
	dep init
	  做的事情
	  	利用gps分析当前代码包中的包依赖关系；
		将分析出的项目包的直接依赖(即main.go显式import的第三方包，direct dependency)约束(constraint)写入项目根目录下的Gopkg.toml文件中；
		将项目依赖的所有第三方包（包括直接依赖和传递依赖transitive dependency）在满足Gopkg.toml中约束范围内的最新version/branch/revision信息写入Gopkg.lock文件中；
		创建root vendor目录，并且以Gopkg.lock为输入，将其中的包（精确checkout 到revision）下载到项目root vendor下面。
	  生成文件
		├── Gopkg.lock
		├── Gopkg.toml
		├── main.go
		└── vendor/	
		vender 示例：
			vendor
				├── github.com
				│   └── beego
				└── go.uber.org
				    ├── atomic
				    └── zap
	  vendor flatten平坦化
	  	依赖包中存在带有冲突的约束，那么dep init必将失败
	  在
	dep status
	dep 
	dep ensure
	dep ensure  		  	install the project's dependencies

	dep ensure 'go.uber.org/zap@<1.4.0'  更新指定依赖，不会修改Gopkg.toml
 	dep ensure -update	  	update the locked versions of all dependencies
 		手动修改了 Gopkg.lock 文件之后需要执行这条命令
 	dep ensure -add github.com/pkg/errors 		add a dependency to the project
 	ep remove github.com/gorilla/mux

 	Gopkg.toml
	 	# 必需包
		required = ["github.com/gin-gonic/gin"]
		# 忽略包
		#ignored = []没有可以不写
		# 项目元数据
		#[metadata]
		 
		 
		# 约束条件
		[[constraint]]
		 # name = 
		 # 可选：版本
		 # version =
		 # 分支
		 # branch
		 # 修订
		 # revision
		 # 可选：指定来源
		 # source = "github.com/gin-gonic/gin"


不同类型的变量赋nil后 仍然是不相等的

go build	编译包和依赖
go install	编译和安装包和依赖 在 $GOPATH/bin目录生成 同名二进制可执行文件
go doc 		查看某个包或者symbol的文档
	go doc --http=:6060	在web浏览器查看帮助文档
go list		列出包和依赖
go run 		编译和运行 go 程序，从main包 函数main() 开始编译和运行
go test		测试包 包名 *_test.go package同名 import "testing"  函数名 Test*
go version 	go 版本号
go vet 		报告可能的错误
go get		下载和安装包和依赖
go env		Go 环境信息
go clean	移除目标文件和缓存文件
go bug		开始一个bug报告
go fix		自动格式化代码文件并保存

把包的源文件放到公用代码库根目录就好

Golang 精编100题
https://blog.csdn.net/itcastcpp/article/details/80462619
var array [5]int
var slice []int
mp := make(map[int]func() int) // func() 括号不能丢

var iii interface{} = 23	   // 所有类型都实现了 空接口 interface{}
	fmt.Printf("%T", iii)      // 输出 int
}

type II interface {
	get() int
	push(int) 
	put(int) string 
}



协程死锁
	阻塞是死锁的必要条件，但不是充分条件
	所有的goroutine阻塞才会死锁抛出panic


golang中分为值类型和引用类型
值类型分别有：int系列、float系列、bool、string、数组和结构体
引用类型有：指针、slice切片、管道channel、接口interface、map、函数等
值类型的特点是：变量直接存储值，内存通常在栈中分配
引用类型的特点是：变量存储的是一个地址，这个地址对应的空间里才是真正存储的值，内存通常在堆中分配

34. 【中级】 golang中的指针运算包括（BC）
A. 可以对指针进行自增或自减运算
B. 可以通过“&”取指针的地址
C. 可以通过“*”取指针指向的数据
D. 可以对指针进行下标运算
35. 【初级】关于main函数（可执行程序的执行起点），下面说法正确的是（ABCD）
A. main函数不能带参数
B. main函数不能定义返回值
C. main函数所在的包必须为main包
D. main函数中可以使用flag包来获取和解析命令行参数

在 Go 中，任何类型在未初始化时都对应一个零值：
	布尔类型是 false ，整型是 0 ，字符串是 "" ，
	而指针，函数，interface，
		slice，channel和map的零值都是 nil
	bool      -> false                              
	numbers -> 0                                 
	string    -> ""      

	pointers -> nil
	slices -> nil
	maps -> nil
	channels -> nil
	functions -> nil
	interfaces -> nil



b = a
值语义：b修改不会影响a的值 就是值类型
	内置类型-基本类型-原始类型-按值传递： 数值类型	  字符串类型	bool类型 字符类型
	数组类型: 按值传递 与C相比不同
	struct：按值传递
引用语义：b修改会影响a的值，这就是引用类型
引用类型-按指针传递：切片slice 映射map 通道channel 函数 指针pointer 接口interface

派生类型【指针pointer 接口interface struct】

36. 【中级】下面赋值正确的是（BDEFGHI）
	A. var x = nil
	B. var x interface{} = nil
	C. var x string = nil
	D. var x error = nil
	E. var x func() = nil
	F. var x *func() = nil
	G. var x chan int = nil
	H. var x map[int]string = nil
	I. var x []string = nil
1. Go的文档中说到，nil是预定义的标识符，代表指针、通道、函数、接口、映射或切片的零值,并不是GO 的关键字之一
2. nil只能赋值给 指针、channel、func、interface、map或slice类型的变量 (非基础类型) 否则会引发 panic
对于数组，声明即分配内存
	var arr [3]int



{}什么时候可以省略逗号什么时候不可以呢
	当最后一个元素与 } 在同一行 可以省略 也可以不省略
	反之必须添加逗号
逗号不能丢
x :=[]int{ 1, 2, 3,
	4, 5, 6, 
}
x := map[int]string{1: "aaa",
}

逗号可以丢
x :=[]int{ 1, 2, 3,
	4, 5, 6 }
x :=[]int{ 1, 2, 3, 4, 5, 6 }
x := map[int]string{1: "aaa"}
x := map[int]string{
	1: "aaa"}

指针
Go中不支持指针运算以及 -> 运算符，而直接采用 
. 选择符来操作指针目标对象成员
操作符 & 取变量地址，使用 * 通过指针间接访问目标对象
默认值为nil 而非 NULL

自增和自减
	在Go中 ++ 和 -- 只能作为语句而非表达式
	a++        //对的 只能作为一个语句使用
	++a 	   //错的 没有这种用法
	s := a++ //错的
	fmt.Printf("%v\n", x++)  //错的

函数声明
	返回值 要么都命名 要么都不命名
	func f(a, b int) (value int, error) //错误的

	以下几个是对的
	func f(a, b int) (value int, err error)
	func f(a int, b int) (value int, err error)
	func f(a int, b int) (int, int, error)


类型断言 type assertion
  类型断言就是将接口类型的值(x)，装换成类型(T)。格式为：
	x.(T)
	v:=x.(T)
	v,ok:=x.(T)
  类型断言的必要条件就是x是接口类型，非接口类型的x不能做类型断言：
  	var i int=10
	v:=i.(int) //错误 i不是接口类型
  T可以是非接口类型，如果想断言合法，则T应该实现x的接口
  T也可以是接口，则x的动态类型也应该实现接口T
  类型断言如果非法，运行时就会出现错误，为了避免这种错误，可以使用一下语法：
  	v,ok:=x.(T)
  	ok代表类型断言是否合法，如果非法,ok则为false,这样就不会出现panic了
  空接口赋值后方法
  	接口的结构
  		指向类型相关信息的指针
		指向数据相关信息的指针
		&(x.(T)) //无法获取断言的地址
  	var x interface{}
  	第一种：
	x = &I{ i: 1, s: "a"}
	x.(*I) 可以调用以 *I 和 I 类型值为接收者的值
	第二种：
	x = I{ i: 1, s: "a"}
	x.(I) 只可以可以调用以 I 类型值为接收者的值
	x.(I)是找不到地址的


？？？？
81. 【中级】关于GoMock，下面说法正确的是（）
A. GoMock可以对interface打桩
B. GoMock可以对类的成员函数打桩
C. GoMock可以对函数打桩
D. GoMock打桩后的依赖注入可以通过GoStub完成

参考答案：AD


82. 【中级】关于接口，下面说法正确的是（）
A. 只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等价的，可以相互赋值
B. 如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A
C. 接口查询是否成功，要在运行期才能够确定
D. 接口赋值是否可行，要在运行期才能够确定		//在编译时就可以确定

参考答案：ABC

83. 【初级】关于channel，下面语法正确的是（）
A. var ch chan int
B. ch := make(chan int)
C. <- ch
D. ch <-

D 缺少写入的值
参考答案：ABC


？？？？
84. 【初级】关于同步锁，下面说法正确的是（）
A. 当一个goroutine获得了Mutex后，其他goroutine就只能乖乖的等待，除非该goroutine释放这个Mutex
B. RWMutex在读锁占用的情况下，会阻止写，但不阻止读
C. RWMutex在写锁占用情况下，会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine独占
D. Lock()操作需要保证有Unlock()或RUnlock()调用与之对应
 正确的是：Lock()操作需要保证有Unlock()调用与之对应

参考答案：ABC

？？？？
85. 【中级】 golang中大多数数据类型都可以转化为有效的JSON文本，下面几种类型除外（）
A. 指针
B. channel
C. complex
D. 函数

参考答案：BCD


？？？？
86. 【中级】关于go vendor，下面说法正确的是（）
A. 基本思路是将引用的外部包的源代码放在当前工程的vendor目录下面
B. 编译go代码会优先从vendor目录先寻找依赖包
C. 可以指定引用某个特定版本的外部包
D. 有了vendor目录后，打包当前的工程代码到其他机器的$GOPATH/src下都可以通过编译

参考答案：ABD



==操作最重要的一个前提是：两个操作数类型必须相同！类型必须相同！类型必须相同！
87. 【初级】 flag是bool型变量，下面if表达式符合编码规范的是（）
A. if flag == 1
B. if flag
C. if flag == false
D. if !flag

参考答案：BD

88. 【初级】 value是整型变量，下面if表达式符合编码规范的是（）
A. if value == 0
B. if value
C. if value != 0
D. if !value

参考答案：AC


89. 【中级】关于函数返回值的错误设计，下面说法正确的是（）
A. 如果失败原因只有一个，则返回bool
B. 如果失败原因超过一个，则返回error
C. 如果没有失败原因，则不返回bool或error
D. 如果重试几次可以避免失败，则不要立即返回bool或error

参考答案：ABCD

90. 【中级】关于异常设计，下面说法正确的是（）
A. 在程序开发阶段，坚持速错，让程序异常崩溃
B. 在程序部署后，应恢复异常避免程序终止
C. 一切皆错误，不用进行异常设计
D. 对于不应该出现的分支，使用异常处理

参考答案：ABD

当未为channel分配内存时，channel就是nil channel，例如var ch1 chan int。
nil channel会阻塞对该channel的所有读、写。
所以，可以将某个channel设置为nil，进行强制阻塞，对于select分支来说，就是强制禁用此分支。
93. 【中级】关于channel的特性，下面说法正确的是（）
A. 给一个 nil channel 发送数据，造成永远阻塞
B. 从一个 nil channel 接收数据，造成永远阻塞
C. 给一个已经关闭的 channel 发送数据，引起 panic
D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值

参考答案：ABCD


94. 【中级】关于无缓冲和有缓冲的channel，下面说法正确的是（）
A. 无缓冲的channel是默认的缓冲为1的channel
B. 无缓冲的channel和有缓冲的channel都是同步的
C. 无缓冲的channel和有缓冲的channel都是非同步的
D. 无缓冲的channel是同步的，而有缓冲的channel是非同步的

参考答案：D


95. 【中级】关于异常的触发，下面说法正确的是（）
A. 空指针解析
B. 下标越界
C. 除数为0
D. 调用panic函数

参考答案：ABCD

重点啊
96. 【中级】关于cap函数的适用类型，下面说法正确的是（）
A. array
B. slice
C. map
D. channel

参考答案：ABD

？？？？
97. 【中级】关于beego框架，下面说法正确的是（）
A. beego是一个golang实现的轻量级HTTP框架
B. beego可以通过注释路由、正则路由等多种方式完成url路由注入
C. 可以使用bee new工具生成空工程，然后使用bee run命令自动热编译
D. beego框架只提供了对url路由的处理，而对于MVC架构中的数据库部分未提供框架支持

参考答案：ABC

？？？？
98. 【中级】关于goconvey，下面说法正确的是（）
A. goconvey是一个支持golang的单元测试框架
B. goconvey能够自动监控文件修改并启动测试，并可以将测试结果实时输出到web界面
C. goconvey提供了丰富的断言简化测试用例的编写
D. goconvey无法与go test集成

参考答案：ABC

go vet命令的参数既可以是代码包的导入路径，也可以是Go语言源码文件的绝对路径或相对路径。但是，这两种参数不能混用。也就是说，go vet命令的参数要么是一个或多个代码包导入路径，要么是一个或多个Go语言源码文件的路径。
go vet命令是go tool vet命令的简单封装。它会首先载入和分析指定的代码包，并把指定代码包中的所有Go语言源码文件和以“.s”结尾的文件的相对路径作为参数传递给go tool vet命令。其中，以“.s”结尾的文件是汇编语言的源码文件。如果go vet命令的参数是Go语言源码文件的路径，则会直接将这些参数传递给go tool vet命令。
如果我们直接使用go tool vet命令，则其参数可以传递任意目录的路径，或者任何Go语言源码文件和汇编语言源码文件的路径。路径可以是绝对的也可以是相对的。
go vet和go tool vet实际上是两个分开的命令。
go vet，只在一个单独的包内可用，不能使用flag 选项（来激活某些指定的检测）。
go tool vet更加完整，它可用用于文件和目录。目录被递归遍历来找到包。go tool vet也可以按照检测的分组激活选项。
go添加了限制： 不能直接调用 go tool vet ，推荐go vet
	vet: invoking "go tool vet" directly is unsupported; use "go vet"
99. 【中级】关于go vet，下面说法正确的是（）
A. go vet是golang自带工具go tool vet的封装
B. 当执行go vet database时，可以对database所在目录下的所有子文件夹进行递归检测
C. go vet可以使用绝对路径、相对路径或相对GOPATH的路径指定待检测的包
D. go vet可以检测出死代码

参考答案：ACD


vet是一个优雅的工具，每个Go开发者都要知道并会使用它。它会做代码静态检查发现可能的bug或者可疑的构造
99. 【中级】关于go vet，下面说法正确的是（）
A. go vet是golang自带工具go tool vet的封装. //简单封装
B. 当执行go vet database时，可以对database所在目录下的所有子文件夹进行递归检测
C. go vet可以使用绝对路径、相对路径或相对GOPATH的路径指定待检测的包
D. go vet可以检测出死代码

参考答案：ACD

？？？？
100.【中级】关于map，下面说法正确的是（）
A. map反序列化时json.unmarshal的入参必须为map的地址
B. 在函数调用中传递map，则子函数中对map元素的增加不会导致父函数中map的修改
C. 在函数调用中传递map，则子函数中对map元素的修改不会导致父函数中map的修改
D. 不能使用内置函数delete删除map的元素

参考答案：A


？？？？
101.【中级】关于GoStub，下面说法正确的是（）
A. GoStub可以对全局变量打桩
B. GoStub可以对函数打桩
C. GoStub可以对类的成员方法打桩
D. GoStub可以打动态桩，比如对一个函数打桩后，多次调用该函数会有不同的行为

参考答案：ABD


select就是用来监听和channel有关的IO操作，当 IO 操作发生时，触发相应的动作。
所有channel表达式都会被求值、所有被发送的表达式都会被求值。求值顺序：自上而下、从左到右.
如果有一个或多个IO操作可以完成，则Go运行时系统会随机的选择一个执行，
	否则的话，如果有default分支，则执行default分支语句，
	如果连default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行.
break 可以终止某个case，该case中break之后的代码不再执行
https://www.jianshu.com/p/2a1146dc42c3
102.【初级】关于select机制，下面说法正确的是（）
A. select机制用来处理异步IO问题
B. select机制最大的一条限制就是每个case语句里必须是一个IO操作
C. golang在语言级别支持select关键字
D. select关键字的用法与switch语句非常类似，后面要带判断条件

参考答案：ABC





103.【初级】关于内存泄露，下面说法正确的是（）
A. golang有自动垃圾回收，不存在内存泄露
B. golang中检测内存泄露主要依靠的是pprof包
C. 内存泄露可以在编译阶段发现
D. 应定期使用浏览器来查看系统的实时内存信息，及时发现内存泄露问题

参考答案：BD


7.   【初级】声明一个只用于读取int数据的单向channel变量ch__________

参考答案：var ch <-chan int



8.   【初级】假设源文件的命名为slice.go，则测试文件的命名为__________

参考答案：slice_test.go


9.   【初级】 go test要求测试函数的前缀必须命名为__________

参考答案：Test


defer 指定的函数会在函数退出前调用
多个defer 后进先出
10. 【中级】下面的程序的运行结果是__________

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

参考答案：4 3 2 1 0


13. 【中级】下面的程序的运行结果是__________
func main() {
 x := 1
 {
   	x := 2
    fmt.Print(x)
 }
 fmt.Println(x)
}

参考答案：21





21. 【中级】下面的程序的运行结果是__________
 func main() {
 strs := []string{"one","two", "three"}
  
 for _, s := range strs {
    go func() {
        time.Sleep(1 * time.Second)
        fmt.Printf("%s ", s)
    }()
 }
 time.Sleep(3 * time.Second)
}

参考答案：threethreethree
在运行过程中,发现callApiServer中参数apiServerAddr实际值一直是apiServerAddrList的最后一个元素值
for _, apiServerAddr := range apiServerAddrList {
	go func() {
		callApiServer(apiServerAddr)
	}()
}
每一个创建的协程中apiServerAddr则是循环的数值.
for _, apiServerAddr := range apiServerAddrList {
	go callApiServer(apiServerAddr)
}
原因：
	循环中,go callApiServer(apiServerAddr)创建的协程中,使用的apiServerAddr是其数值;
	循环中,go func(){callApiServer(apiServerAddr)}()创建的协程中,使用的apiServerAddr是其 引用 ,在循环执行至一定次数后,Golang调度器调度协程,协程获取当前时刻apiServerAddr的数值,此时就会产生错误.
闭包,能够从创建它的作用域中获取变量,且获取的是变量的 引用 而  非变量的副本.



32. 【中级】下面的程序的运行结果是__________
func main() {  
 x := []string{"a", "b","c"}
 for v := range x {
    fmt.Print(v)
 }
}

参考答案：012


defer只能执行一个函数。
defer后是一个链式函数而且defer是一个压栈的输入，
2被defer执行，1和3直接被执行，所以打印1234
func main() {  
 s := NewSlice()
 defer s.Add(1).Add(2)
 s.Add(3)
}

参考答案：132


基本类型：数值类型 字符串类型string 字符类型(byte rune) 布尔类型bool
复合类型：指针 数组 切片 map struct
值类型：基本类型 + 数组 + 结构体struct
	值类型变量声明后，不管是否已经赋值，编译器为其分配内存，此时该值存储于栈上。 默认相应类型零值
	值类型的变量的值存储在 栈 中
	栈：临时 局部 自动申请 自动释放(出栈时)
引用类型： 复合类型 - 数组 + 管道channel、接口interface、函数
	变量直接存放的就是一个内存地址值，这个地址值指向的空间存的才是值。所以修改其中一个，另外一个也会修改（同一个内存地址）。
	引用类型必须申请内存才可以使用，make()是给引用类型申请内存空间
	引用类型的实例分配在 堆 上
	堆：持久化 全局 手动申请 手动释放
4.   【初级】指针是基础类型（）

参考答案：F


5.【初级】 interface{}是可以指向任意对象的Any类型（）
参考答案：T


err就是错误   错误是异常的一种 ，
和painc异常 一样，painc异常是异常一种。
如果 open 出错，那么在 file.Close() 会抛出异常
6.【中级】下面关于文件操作的代码可能触发异常（）

 file, err := os.Open("test.go")
 defer file.Close()
 if err != nil {
  fmt.Println("open file failed:",err)
  return
 }
...

参考答案：T


？？？？
反射最常见的使用场景是做对象的序列化（serialization，有时候也叫Marshal & Unmarshal）。
例如，Go语言标准库的encoding/json、encoding/xml、encoding/gob、encoding/binary等包就 大量依赖于反射功能来实现。
序列化是指把对象（可以是stuct,string）按照协议编码填充头部和内容 变成二进制字节码 
反射是指通过对象本身获取类型、方法、字段元信息以及对象的value
深知事执行对象本身的调用 和序列化是两个事情，两个东西可以相互独立存在
14. 【初级】 Golang支持反射，反射最常见的使用场景是做对象的序列化（）

参考答案：T



CGO是C语言和Go语言之间的桥梁，原则上无法直接支持C++的类。
CGO不支持C++语法的根本原因是C++至今为止还没有一个二进制接口规范(ABI)。
CGO只支持C语言中值类型的数据类型，所以我们是无法直接使用C++的引用参数等特性的。
CGO不支持C++，但是可以通过C来封装C++的方法实现CGO调用
15. 【初级】 Golang可以复用C/C++的模块，这个功能叫Cgo（）

参考答案：F

 

16. 【初级】下面代码中两个斜点之间的代码，比如json:"x"，作用是X字段在从结构体实例编码到JSON数据格式的时候，使用x作为名字，这可以看作是一种重命名的方式（）
type Position struct {
 X int `json:"x"`
 Y int `json:"y"`
 Z int `json:"z"`
}

参考答案：T



21. 【初级】通过成员变量或函数首字母的大小写来决定其作用域（）

参考答案：T




Go语言的常量有个不同寻常之处。虽然一个常量可以有任意有一个确定的基础类型，例如int或float64，或者是类似time.Duration这样命名的基础类型，
但是许多常量并没有一个明确的基础类型。
编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算；
	你可以认为至少有256bit的运算精度。
这里有六种未明确类型的常量类型，分别是
	无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。
22. 【初级】对于常量定义zero(const zero = 0.0)，zero是浮点型常量（）

参考答案：F


GO 语言中, ^ 既是按位取反操作符，也是异或的操作符，没有 ~ 操作符
负数的存储形式是补码（反码 + 1）
23. 【初级】对变量x的取反操作是~x（）

参考答案：F

0编码问题
原则： 对于负数，首位不参与运算，负数的补码等于反码加一；那么，-0的补码等于：1000 0000 +0的补码等于：0000 0000
	+0 和 -0
	+0:
	原码：0000 0000 
	补码：0000 0000 

	-0：
	原码：1000 0000 
	反码：1111 1111
	补码：1000 0000 
什么？0 虽然在signed情况下，有两种表达方式（+0，和-0），但是毕竟都是一个数0啊，怎么0可以有两个补码？不允许啊，这台浪费了吧!!!
因此，在不能浪费的理念下，计算机又新增规定：对于0有+0， -0这个特殊情况下，这个时候，最高符号位要参与运算；此时，无论是数字0来说，无论是+0 还是 -0，其补码只有一个，那就是0000 0000；
这个时候，你又会发现，多余了一个1000 0000 没人用吗？怎么办呢？ 在计算器中，又出台规定将这个没有用的1000 0000表示-128;
所以，你会看到，对于有符号数，数据存储范围多了一个:-128
数据范围：
	有符号数：[-128, 127]
	无符号数：[0, 255]
在有符号数据类型中，使用补码的好处是：
	解决了0的编码问题
	可以多保存一个数：-128， 如signed char 可以多保存一个:-128
	多保存的一个数：-128，没有原码和反码
为什么使用补码来存储数据？
	解决0的编码问题
	减法运算可以转化为加法运算，省去硬件上的减法电路，
		CPU只需要有：全加器，求补电路


字符串 str[0] = 'z'  不能这样赋值
24. 【初级】下面的程序的运行结果是xello（）

func main() {
 str := "hello"
 str[0] = 'x'
 fmt.Println(str)
}

参考答案：F




四个问题: 
1.什么是野指针 
	野指针是指向不确定可造成危害的指针变量 
2.如何产生野指针 
	a.初始化时指向不明确 
	b.初始化指向明确，使用中销毁了其所指向变量，但指针指向未调整 
3.野指针的危害 
	a.指向不可访问地址 
		触发段错误 
	b.指向可用空间 但是这块空间不使用，程序正常run 
	c.指向可用空间，但是这块空间刚好在用，造成数据错误
		如果使用指针对其指向空间内容进行了修改，
		实际上这部分内容正在被使用，会对程序产生影响，数据损坏甚至程序崩溃
4.如何避免
	1、定义一个指针变量时一定记得初始化
	2、动态开辟的内存空间使用完free之后一定将对应的指针置为NULL/nil
	3、不要在函数中返回栈空间的指针和引用
	4、注意在使用时对指针的合法性的判断
30. 【初级】下面代码中的指针p为野指针，因为返回的栈内存在函数结束时会被释放（）
 type TimesMatcher struct {
 	base int
 }
 func NewTimesMatcher(base int) *TimesMatcher{
 	return &TimesMatcher{base:base}
 }
 func main() {
 	p := NewTimesMatcher(3)
  ...
 }

参考答案：F


40. 【初级】匿名函数可以直接赋值给一个变量或者直接执行（）

参考答案：T

41. 【初级】如果调用方调用了一个具有多返回值的方法，但是却不想关心其中的某个返回值，
可以简单地用一个下划线“_”来跳过这个返回值，该下划线对应的变量叫匿名变量（）

参考答案：T

42. 【初级】在函数的多返回值中，如果有error或bool类型，则一般放在最后一个（）

参考答案：T


43. 【初级】错误是业务过程的一部分，而异常不是（）

参考答案：T
错误指的是可能出现问题的地方出现了问题，比如打开一个文件时失败，这种情况在人们的意料之中 ；
而异常指的是不应该出现问题的地方出现了问题，比如引用了空指针，这种情况在人们的意料之外。
可见，错误是业务过程的一部分，而异常不是 。


panic 导致了异常，延迟函数仍然会执行
44. 【初级】函数执行时，如果由于panic导致了异常，则延迟函数不会执行（）

参考答案：F



45. 【中级】当程序运行时，如果遇到引用空指针、下标越界或显式调用panic函数等情况，则先触发panic函数的执行，
	然后调用延迟函数。调用者继续传递panic，
	因此该过程一直在调用栈中重复发生：函数停止执行，调用延迟执行函数。
	如果一路在延迟函数中没有recover函数的调用，则会到达该协程的起点，该协程结束，然后终止其他所有协程，
	其他协程的终止过程也是重复发生：函数停止执行，调用延迟执行函数（）

参考答案：F
先调用defer 然后触发panic




46. 【初级】同级文件的包名不允许有多个（）

参考答案：T


方法的接收者必须是 当前文件定义或者type的类型
47. 【中级】可以给任意类型添加相应的方法（）

参考答案：F


什么是匿名组合？
	struct 嵌套
	注意匿名冲突和隐式名字(如果嵌套的struct默认名字就是类型名也就是struct名字)
48. 【初级】 golang虽然没有显式的提供继承语法，但是通过匿名组合实现了继承（）

参考答案：T


49. 【初级】使用for range迭代map时每次迭代的顺序可能不一样，因为map的迭代是随机的（）

参考答案：T


50. 【初级】 switch后面可以不跟表达式（）

参考答案：T


51. 【中级】结构体在序列化时非导出变量（以小写字母开头的变量名）不会被encode，
	因此在decode时这些非导出变量的值为其类型的零值（）

参考答案：T


52. 【初级】 golang中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名（）

参考答案：T

多个defer时
	返回值与return 要严格匹配，顺序、类型、个数完全匹配
	defer函数参数是值而非指针如 []int int 等, 即使是引用类型传递的也是包括底层的数据的完整副本
		这个参数在defer定义时是什么值，传递的参数就是什么值
	defer函数参数明确指出是 指针 如 *[]int *map[int]string 等，defer函数间共享的数据会受影响
	func example() {
		res := []int{1,2,3}
		defer func(res []int){
			fmt.Println(res)
		}(res)
		defer func(res []int){
			res = []int{111111}
			fmt.Println(res)

		}(res)
		res = []int{5, 6, 7}
		defer func(res []int){
			fmt.Println(res)
		}(res)
	}
	输出 [5 6 7]
		&[111111]
		[1 2 3]
	func example() {
		res := []int{1,2,3}
		defer func(res *[]int){
			fmt.Println(res)
		}(&res)
		defer func(res *[]int){
			*res = []int{111111}
			fmt.Println(res)

		}(&res)
		res = []int{5, 6, 7}
		defer func(res []int){
			fmt.Println(res)
		}(res)
	}
	输出  [5 6 7]
		 &[111111]
		 &[111111]
	函数有返回值时
		不是命名返回值，defer中修改无效
		是命名返回值，
			defer中修改有效，return 必须放在代码最后一行 
				return 返回与命名返回值同类型的值，这个值默认赋给命名返回值
				之后 defer 获得命名返回值是已经被赋值后的。可以对其进行修改
			defer修改了命名返回值，即使return 其他的同类型值(不同类型报错)，也是返回命名返回值
			defer 如果没有给命名返回值赋值，return 其他的同类型值 会作为命名返回值返回
	函数无返回值时 
		return 可以放在任何一行
		return之后的代码不再执行，之后的defer也不执行
一个defer执行时 发生panic,当前defer中panic之后代码不再执行。不影响其他defer的执行
panic 发生在函数中，panic之后的代码包括defer不再执行。
可以用recover()捕捉panic
if r := recover(); r!= nil {
	fmt.Println("Recover from r : ",r)
}


80. 【中级】 channel本身必然是同时支持读写的，所以不存在单向channel（）

参考答案：F

close(ch) close关闭channel
给一个已经关闭的 channel 发送数据，引起 panic
从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
如果不关闭，所有的与之相关的协程都会被阻塞
单向channel无法转换为双向
channel定义方式。都是 nil
var ch1 chan int 　　　　　　//ch1是一个正常的channel，不是单向的
var ch2 chan <- float64 　　  //ch2是一个单向的channel，只用于写float64的数据
var ch3  <- chan  int　　　　 //ch3是一个单向的channel，只用于读取int数据
使用usage：
ch := make(chan int)
var writech chan <- int = ch
var readch <- chan int = ch
writech <- 555
x := <-readch
作为参数传递，双向channel可以作为单项channel参数进行传递
如下usage：
func producer(out chan <- int)  {
    for i:=0; i<=10; i++ {
        out<-i*i
    }
    close(out)
}
func consumer(in <-chan int)  {
    for num := range in{
        fmt.Println("num = ", num)
    }
}
func main() {
    //创建一个双向通道ch
    ch  := make(chan int)

    //生产者，生产数字，写入channel
    go producer(ch) //chennel传参，引用传递

    //消费者，从channel里读取数字、然后打印
    consumer(ch)
}



81. 【初级】 import后面的最后一个元素是包名（）

参考答案：F

是目录名不是包名



4 下面代码会输出什么？

type People struct{}
func (p *People)ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func(p*People)ShowB() {
    fmt.Println("showB")
}
typeTeacher struct {
    People
}
func(t*Teacher)ShowB() {
    fmt.Println("teachershowB")
}
funcmain() {
    t := Teacher{}
    t.ShowA()
}

考点：go的组合继承
解答：
这是Golang的组合模式，可以实现OOP的继承。 被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），
但它们的方法(ShowA())调用时接受者并没有发生变化。 此时People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。
showAshowB



5 下面代码会触发异常吗？请详细说明
func main() {
    runtime.GOMAXPROCS(1)
    int_chan := make(chan int, 1)
    string_chan := make(chans tring, 1)
    int_chan <- 1
    string_chan <- "hello"
    select {   
			case value := <-int_chan:
       fmt.Println(value)
			case value := <-string_chan:
				panic(value)
    }
}

考点：select随机性
解答：
select会随机选择一个可用通用做收发操作。 所以代码是有肯触发异常，也有可能不会。 
单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：

select 中只要有一个case能return，则立刻执行。
当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
如果没有一个case能return则可以执行”default”块。

 

6 下面代码输出什么？
func calc(index string, a,b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {   
	a := 1
  b := 2
  defer calc("1", a, calc("10", a, b))    
	a = 0
  defer calc("2", a, calc("20", a, b))    
	b = 1
}

考点：defer执行顺序
解答：
这道题类似第1题 需要注意到defer执行顺序和值传递 index:1肯定是最后执行的，
但是index:1的第三个参数是一个函数，所以最先被调用
calc("10",1,2)==>10,1,2,3 执行index:2时,与之前一样，
需要先调用calc("20",0,2)==>20,0,2,2
执行到b=1时候开始调用，index:2==>calc("2",0,2)==>2,0,2,2
	最后执行index:1==>calc("1",1,3)==>1,1,3,4
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4


7 请写出以下输入内容
func main() {    
	s := make([]int,5)
	s = append(s,1, 2, 3)
	fmt.Println(s)
}

考点：make默认值和append
解答：
make初始化是由默认值的哦，此处默认值为0

[00000123]

大家试试改为:
s := make([]int, 0)
s = append(s, 1, 2, 3)
fmt.Println(s)//[1 2 3]
 



TODO
8 下面的代码有什么问题?
type UserAges struct {
	ages map[string]int
	sync.Mutex
}
func(ua*UserAges) Add(name string, age int) {
	ua.Lock()  
	defer ua.Unlock()
	ua.ages[name] = age
}
func(ua*UserAges) Get(name string) int {    
	if age, ok := ua.ages[name]; ok {        
		return age
	}  
	return-1
}

考点：map线程安全
解答：
可能会出现
fatal error: concurrent map read and map write.

修改一下看看效果
func (ua *UserAges) Get(name string) int {
	ua.Lock()    
	defer ua.Unlock()    
	if age, ok := ua.ages[name]; ok {        
		return age
	}    
	return-1
}




TODO
9. 下面的迭代会有什么问题？
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})   
	go func() {
		set.RLock()   
		for elem := range set.s {
			ch <- elem
		}
		close(ch)
		set.RUnlock()
	}()
  return ch
}

考点：chan缓存池
解答：
看到这道题，我也在猜想出题者的意图在哪里。 chan?sync.RWMutex?go?chan缓存池?迭代? 
所以只能再读一次题目，就从迭代入手看看。 既然是迭代就会要求set.s全部可以遍历一次。
但是chan是为缓存的，那就代表这写入一次就会阻塞。 我们把代码恢复为可以运行的方式，看看效果

package main
import ("sync", "fmt")//下面的迭代会有什么问题？
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}
func (set *threadSafeSet) Iter() <-chan interface{} {    
	//ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{}, len(set.s))   
	go func() {
		set.RLock()       
		for elem,value := range set.s {
			ch <- elem            
			println("Iter:", elem, value)
    }       
		close(ch)
    set.RUnlock()
  }()    
	return ch
}
func main() {
	th := threadSafeSet{
		s: []interface{}{ "1", "2" },
	}
  v := <-th.Iter()
  fmt.Sprintf("%s%v", "ch", v)
}




10 以下代码能编译过去吗？为什么？
package main
import ("fmt")

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
    talk = "Youare a good boy"
  } else {
    talk = "hi"
  }
  return
}
func main() {
  var peo People = Stduent{}
  think := "bitch"
  fmt.Println(peo.Speak(think))
}

考点：golang的方法集
解答：
编译不通过！ 做错了！？说明你对golang的方法集还有一些疑问。 
一句话：golang的方法集仅仅影响接口实现和方法表达式转化，
与通过实例或者指针调用方法无关。
Speak方法接收者是*Stduent，应该这样 初始化
	var peo People = &Stduent{}





11 以下代码打印出来什么内容，说出为什么。
package main
import ("fmt")
type People interface {
	Show()
}
type Student struct{}
func (stu*Student) Show() {}
func live() People {
	var stu *Student
	return stu
}
func main() {   
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

考点：interface内部结构
解答：
很经典的题！ 这个考点是很多人忽略的interface内部结构。 go中的接口分为两种一种是空的接口类似这样：
var in interface{}
另一种如题目：
type People interface {
	Show()
}

他们的底层结构如下：

type eface struct {      //空接口
    _type *_type         //类型信息
    data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type iface struct {      //带有方法的接口
    tab  *itab           //存储type信息还有结构实现方法的集合
    data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type _type struct {
    size       uintptr //类型大小
    ptrdata    uintptr //前缀持有所有指针的内存大小
    hash       uint32  //数据hash值
    tflag     tflag
    align      uint8   //对齐
    fieldalign uint8   //嵌入结构体时的对齐
    kind       uint8   //kind 有些枚举值kind等于0是无效的
    alg       *typeAlg //函数指针数组，类型实现的所有方法
    gcdata    *byte   
		str       nameOff
    ptrToThis typeOff
}
type itab struct {
    inter  *interfacetype //接口类型
    _type  *_type         //结构类型
    link   *itab
    bad    int32
    inhash int32
    fun    [1]uintptr     //可变大小方法集合
}

可以看出iface比eface 中间多了一层itab结构。 itab 存储_type信息和[]fun方法集，
从上面的结构我们就可得出，因为data指向了nil 并不代表interface 是nil， 所以返回值并不为空，
这里的fun(方法集) 定义了接口的接收规则，在编译的过程中需要验证是否实现接口 
结果：
	BBBBBBB




12.是否可以编译通过？如果通过，输出什么？
package main
import ("fmt")

func main() {
	i := GetValue()
	switch i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case interface{}:
			fmt.Println("interface")
		default:
			fmt.Println("unknown")
	}
}
func GetValue() int {
	return 1
}

解析
考点：type

编译失败，因为type只能使用在interface
cannot type switch on non-interface value i (type int)

struct 类型变量 的 type 是 struct，case中没有struct类型，也能匹配interface{}
下面是有效的例子
type e interface{}
func mult2(f e) e { 
	switch f.(type) {
		case int:
			return f.(int) * 2
		case string:
			return f.(string) + f.(string) + f.(string) + f.(string)
	}
	return f
}
func Map(n []e, f func(e) e) []e { 
	m := make([]e, len(n))
	for k, v := range n { 
		m[k] = f(v)
	}
	return m
}
func main() {
	m := []e{1, 2, 3, 4}
	s := []e{"a", "b", "c", "d"} 
	mf := Map(m, mult2)
	sf := Map(s, mult2) 
	fmt.Printf("%v\n", mf) 
	fmt.Printf("%v\n", sf)
}


执行顺序：
func GetValue() int {
	// var it interface{}
	it := 222
	go func(){
		it += 444
		fmt.Println("hhhhhhhhhh", it)
	}()
	defer func(){
		it += 333
		fmt.Println(it)
		it += 333
		fmt.Println(it)
		it += 333
		fmt.Println(it)
		it += 333
		fmt.Println(it)
	}()
	return it
}
return
defer 如果defer修改了 命名返回值
go func 这个是异步的 不固定的


示例：iff 必须是接口
func f(i iff) {
	switch i.(type) {
		case interface{}:
			fmt.Println("interface")
		default:
			fmt.Println("unknown")
	}
	if t, ok := i.(interface{}) ; ok { 
		// 对于某些实现了接口 I 的 // t 是其所拥有的类型
		fmt.Println(t, "000000000000")
	}
}


示例：struct 嵌套
type _iint struct {
	i int
}
type iint struct {
	_iint
}
func (ii _iint) ff() {
	fmt.Println("womendoushirenmin")
}
func main() {
	ir := iint{
		_iint: _iint{
			i: 191919191919,
		},
	}
	// 两种都可以
	// ir := iint{
	// 	_iint{
	// 		i: 191919191919,
	// 	},
	// }
	ir.ff()
	fmt.Println(ir)
}




14.是否可以编译通过？如果通过，输出什么？
package main
func main() {    
	println(DeferFunc1(1)) 
	println(DeferFunc2(1)) 
	println(DeferFunc3(1))
}
func DeferFunc1(i int) (t int) {
	t = i   
	defer func() {
		t += 3
	}()
	return t
}
func DeferFunc2(i int) int {
	t := i  
	defer func() {
		t += 3
	}() 
	return t
}
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

解析
考点:defer和函数返回值
需要明确一点是defer需要在函数结束前执行。 
函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数 
DeferFunc1有函数返回值t作用域为整个函数，
	在return之前defer会被执行，所以t会被修改，返回 4; 
DeferFunc2函数中t的作用域为函数，返回 1;
DeferFunc3返回 3



15.是否可以编译通过？如果通过，输出什么？
func main() {
	list := new([]int)
	list = append(list,1)
	fmt.Println(list)
}
不能编译通过
append 不能传递指针
解析
考点：new
list:=make([]int,0)


16.是否可以编译通过？如果通过，输出什么？
package main
import "fmt"
funcmain() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1,s2)
	fmt.Println(s1)
}

解析
考点：append
append切片时候别漏了 "..."


17.是否可以编译通过？如果通过，输出什么？
func main() {
  sn1 := struct {
    age int
    name string
  }{age: 11,name: "qq"}
  sn2 := struct {
    age int
    name string
  }{age: 11,name: "qq"} 
  if sn1 == sn2 {
    fmt.Println("sn1== sn2")
  }
  sl1 := struct {
    age int
    name string
  }{age: 11,name: "qq"}
  sl2 := struct {
    name string
    age int
  }{age: 11,name: "qq"} 
  if sl1 == sl2 {
    fmt.Println("sl1== sl2")
  }
  if sl1 != sl2 {
    fmt.Println("sl1== sl2")
  }
  sm1 := struct {
    age int
    m  map[string]string
  }{age: 11}
  sm2 := struct {
    age int
    m  map[string]string
  }{age: 11} 
  if sm1 == sm2 {
    fmt.Println("sm1== sm2")
  }
}

结果：编译不通过
invalid operation: == !=
编译时会检查 结构体比较(==/!=) 
	两个具名结构体：invalid operation: == !=
	一个具名一个不具名：
		成员变量和成员类型必须完全相同 而且结构体成员不能包括 map、slice类型
		如果满足比较条件，在运行时 才会进行比较值是否相同
		具名结构体有结构体方法
			结构体比较时只比较成员，不比较结构体方法
			编译时会检查：即使两个结构体实例相等 ==，不是同一个结构体实例不能调用另一个结构体的方法
	两个不具名：
		成员变量和成员类型必须完全相同 而且结构体成员不能包括 map、slice类型
		如果满足比较条件，在运行时 才会进行比较值是否相同



18.是否可以编译通过？如果通过，输出什么？
func main() {
  var x *int = nil
  Foo(x)
}

func Foo(x interface{}) {
  if x == nil {
    fmt.Println("emptyinterface")      
    return
  }
  fmt.Println("non-emptyinterface")
}

解析
考点：interface内部结构
non-emptyinterface

19.是否可以编译通过？如果通过，输出什么？
func GetValue(m map[int]string, id int)(string, bool) {
  if _,exist := m[id]; exist {
    return"存在数据", true
  }
  return nil, false
}
func main() {
  intmap:=map[int]string{
    1:"a",
    2:"bb",
    3:"ccc",
  }
  v,err:=GetValue(intmap,3)
  fmt.Println(v,err)
}
结果：编译不通过
	cannot use nil as type string in return argument
nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。
但是如果不特别指定的话，Go 语言不能识别类型，所以会报错。



20.是否可以编译通过？如果通过，输出什么？
const (
  x = iota
  y
  z = "zz"
  k
  p = iota)
func ttt() {
  fmt.Println(x,y,z,k,p)
}

考点：iota
结果:
0 1 zz zz 4



21.编译执行下面代码会出现什么?
package main
var(
    size :=1024
    max_size = size*2)
func main() {    
	println(size,max_size)
}

解析
考点:变量简短模式
var 与 := 不能在同一个初始化中使用
var 可以声明全局也可以声明局部变量
变量简短模式限制：
	定义变量同时显式初始化
	不能提供数据类型
	只能在函数内部使用
结果：
syntax error: unexpected :=


22.下面函数有什么问题？
const cl = 100
var bl   = 123
func main() {
	println(&bl,bl)
  println(&cl,cl)
}

解析
考点:常量
常量不同于变量的在运行期分配内存，
常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，
cannot take the address of cl


23.编译执行下面代码会出现什么?
package main
func main() {    
	for i:=0;i<10;i++  {
		loop: println(i)
	}
	goto loop
}

解析
考点：goto
goto 的位置 必须能够 访问到 label loop
	goto 可以在一个内部作用域，只要能访问到 label loop就可以

goto loop jumps intoblock starting at




24.编译执行下面代码会出现什么?
package main
import "fmt"
func main() {
	type MyInt1 int
  type MyInt2 = int
  var i int =9
  var i1 MyInt1 = i
  var i2 MyInt2 = i
  fmt.Println(i1,i2)
}

解析
考点：**Go 1.9 新特性 Type Alias **
基于一个类型创建一个新类型，称之为defintion；
基于一个类型创建一个别名，称之为 alias。 
MyInt1为称之为defintion，
虽然底层类型为int类型，但是不能直接赋值，需要强转； 
MyInt2称之为 alias，可以直接赋值。

结果:
cannot use i (typeint) as type MyInt1 in assignment


25.编译执行下面代码会出现什么?
type User struct {
}
type MyUser1 User
type MyUser2 = User
func(i MyUser1)m1(){
  fmt.Println("MyUser1.m1")
}
func(i User)m2(){
  fmt.Println("User.m2")
}
func main() {
	var i11 MyUser1
  var i22 MyUser2
  i11.m1()
  i22.m2()
}

考点：**Go 1.9 新特性 Type Alias **
因为MyUser2完全等价于User，所以具有其所有的方法，并且其中一个新增了方法，另外一个也会有。 
但是
	i1.m2()
是不能执行的，因为MyUser1没有定义该方法。 
结果:
MyUser1.m1User.m2


26.编译执行下面代码会出现什么?
func main() {
  my := MyStruct{}
  my.m1()
}
type T1 struct {
}
func(t T1) m1(){
  fmt.Println("T1.m1")
}
type T2 = T1
type MyStruct struct {
  T1
  T2
}

解析
考点：**Go 1.9 新特性 Type Alias **
是不能正常编译的,异常：

ambiguousselectormy.m1

结果不限于方法，字段也也一样；也不限于type alias，type defintion也是一样的，
只要不同的名的嵌入struct 中 有重复的方法、字段，就会有这种提示，因为不知道该选择哪个。 改为:
my.T1.m1()
my.T2.m1()
type alias的定义，本质上是一样的类型，只是起了一个别名，源类型怎么用，
别名类型也怎么用，保留源类型的所有方法、字段等。



27.编译执行下面代码会出现什么?
var ErrDidNotWork = errors.New("did not work")
func DoTheThing(reallyDoIt bool) (err error) {
	// 改为
	// var result string
	// result, err = tryTheThing()
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}
func tryTheThing() (string,error) {
	return "", ErrDidNotWork
}
func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}

考点：变量作用域
因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量，
结果：
<nil>
<nil>

28.编译执行下面代码会出现什么?
package main
func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++  {
		funs = append(funs,func() {
			println(&i,i)
		})
	}
	return funs
}
func main(){
	funs := test()
	for _,f := range funs{
		f()
	}
}

解析
考点：闭包延迟求值
for循环复用局部变量i，每一次放入匿名函数的应用都是想一个变量。 结果：

0xc042046000 2
0xc042046000 2

如果想不一样可以改为：
func test() []func()  {
	var funs []func()
	for i:=0;i<2;i++ {
		x:=i
		funs = append(funs,func() {
			println(&x,x)
		})
	}
	return funs
}

29.编译执行下面代码会出现什么?
package main
func test(x int) (func(),func()) {
	return func() {
		println(x)
		x += 10
	}, func() {
		println(x)
	}
}
func main() {
	a,b :=test(100)
	a()
	b()
}

解析
考点：闭包引用相同变量*
结果：
100
110




30. 编译执行下面代码会出现什么?
package main
import ("fmt";"reflect")
func main1() {
	defer func() {
    if err:=recover(); err!=nil{
      fmt.Println(err)
    } else {
      fmt.Println("fatal")
    }
  }()
  defer func() {
    if err:=recover(); err!=nil{
      fmt.Println(err)
    } else {
      fmt.Println("fatal")
    }
  }()
  defer func() {
    panic("deferpanic")
  }()
  panic("panic")
}
func main() {
  main1()
  defer func() {
    if err:=recover();err!=nil {
      fmt.Println("++++")
      f := err.(func() string)
      fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
    } else {
      fmt.Println("fatal")
    }
  }()
  defer func() {
    panic(func() string {
      return "defer-panic"
    })
  }()
  panic("panic")
}

解析
考点：defer按从下往上(栈)执行 多个连续panic仅有最后一个可以被revover捕获
触发panic("panic")后顺序执行defer，但是defer中还有一个panic，所以覆盖了之前的panic("panic")
结果：
deferpanic
++++
0x109d480 defer-panic func




select作用
	select用来让我们的程序监视多个文件句柄(file descriptor)的状态变化的处理机制。
	当你发起一些阻塞的请求后，你可以使用select机制轮训扫描fd，
	直到被监视的文件句柄有某一个或多个发生了状态改变。

	我们可以把要监控读写的文件交给内核（epoll_add)，而内核通过中断的方式得知事件通知. 这要比select好不少.
	设置你关心的事件（epoll_ctl），比如读事件.
	然后等（epoll_wait），此时，如果没有哪个文件有你关心的事件，则休眠，直到有事件，被唤醒，然后返回那些事件.
	Epoll的优势在于，由接收数据的OS来负责通知你有数据可以操作，因为OS是知道什么时候有数据的。
	select 跟golang有啥关联? 首先我们通过上面的介绍得知，select是通过线性扫描的方式监视文件描述符是否有变动. 
		channnel在系统层面来说也是个文件描述符。 在golang里我们可以使用goroutine并发执行任务，
		接着使用select来监视每个任务的channel情况.  
		但如果这几个任务都长时间没有回复channel信息，如果我们又有超时timeout需求，
		那么我们可以使用起一个goroutine，这个goroutine任务逻辑启动sleep,等sleep之后回复channel信号。

	package main

	import (
	    "fmt"
	    "time"
	)

	func main() {
	    timeout := make(chan bool, 1)
	    go func() {
	        time.Sleep(3 * time.Second) // sleep 3 second
	        timeout <- true
	    }()
	    ch := make(chan int)
	    select {
	    case <-ch:
	    case <-timeout:
	        fmt.Println("task is timeout!")
	    }
	} 


for多个变量避开坑
// 错误写法
func main() {
    sum := 0
    for i, j := 1, 10; i < j; i++,j++ {
        sum += i + j
    }
    fmt.Println("sum is equal to ", sum)
}
// 正确写法
func main() {
    sum := 0
    for i, j := 1, 10; i < j; i, j = i+1, j-1 {
        sum += i + j
    }
    fmt.Println("sum is equal to ", sum)
}
