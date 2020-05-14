go 特性
清晰并且简洁
	Go努力保持小并且优美，可以在短短几行代码里做许多事情
并行
	Go让函数很容易成为非常轻量的线程。goroutines
Channel
	goroutines 之间的通讯由 channel 完成
快速
	编译很快，执行很快。编译用秒计算
安全
	类型转换 显式 并遵循严格规则。垃圾收集 无须free()  语言自动处理
标准格式化
	Go程序可以被格式化为程序员希望的任何形式，但是官方格式是存在的。标准也非常简单：gofmt的输出就是官方认可的格式
类型后置
	类型在变量名的后面， var a int 代替 C中的int a；
UTF-8 
	任何地方都是UTF-8，包括字符串以及程序代码
开源
	Go的许可证是完全开源的
开心
	用Go写程序会非常开心
Erlang 与 Go 在部分功能上类似。其主要区别：Erlang 是函数式语言，运行在虚拟机上；Go是命令式语言，是编译的；

#所有的Go文件以package <something>开头，独立运行的执行文件必须以package main；
package main

#加入fmt，不是main的其他包都是库
import "fmt"

func main(){
	fmt.Printf("Hello, world!")
}

Go文件 package语句首先出现，其次是import语句，然后是其他所有内容
Go程序执行时 首先调用的函数是 main.main()
字符串 由 “""” 双引号包裹

程序构建： go build helloworld.go
执行 ./执行文件

变量定义默认赋值其类型的 null/nil 值
var (
	x int
	b bool
)
var 声明后必须进行初始化
x = 15 
b = false
等价

var x = 15
var b = false

等价

var x int = 15
var b bool = false

等价

a := 15
b := false
:= 不能用来声明 全局变量
或者
a, b := 20, 16

var声明全局变量，初始化必须在函数体内初始化

下划线 用来丢弃
_, b := 34, 35

对声明却未使用的变量报错

boolean 类型   true false
 
数值类型 
	数值类型 与 数值字面量比较，字符字面量会自动转换为对应类型比较，数字相等 == 就返回 true
	int8 int16 int32 int64
	byte(uint8) uint16 uint32 uint64
	float32 float64
 类型全部都是独立的，混合赋值引起编译器错误

类型转换	
合法的转换，float64 同 float32 类似
From 		b []byte 		i []int 		r []rune  	  s string 		f flt32   	i int
To
[]byte 		×																		[]byte(s)
[]int 								×												[]int(s)
[]rune 														×						[]rune(s)
string 	string(b) 	string(i) 	string(r) 	 		× 					
ftl32 																											×					flt32(i)
int 																											int(f) 				×



常量 constant
	只能是 数字 字符串 boolean值
	iota 生成枚举值
	在常量组中，如不提供类型和初始化值，那么视作与上一常量值相同。 
	const (
		a = iota
		b = iota
	)
	也可以
	const (
		a = iota
		b
	)

	const (
		a = 0
		b string = "0"
	)

	const a, b = 0, "0"


字符串 必须使用双引号 ""
	字符串一旦给变量赋值，字符串就不能修改了： 字符串不可变

	如果需要更改 需要：
	s := "Hello World!"
	c := []rune(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	注意多行字符串  
		+ 必须在第一行
		s := "Starting part" +
			"Ending part"
	  用反引号 `   同时包含换行
	  s := `starting part
	  		Ending part`

	rune 是 int32 别名，UTF-8编码 遍历字符串中的字符

	for k, v := range "abc" {
		fmt.Println(k)
		fmt.Println(v)
	}
	k序号 从0开始  v是字符 rune码

复数
	Go原生支持复数。 
	complex128---64位虚数
	complex64---32位虚数
	var c complex64 = 5 + 5i

错误 error
	var e error     定义了一个error类型变量 值为nil

运算符 TODO
	数字元算 > 移位 > 按位操作 > 

	Go不支持运算符重载（方法重载）
	一些内建运算符却支持重载

控制结构
	for
	switch
	if
	select

必须有 {} , 第一个 { 必须与 if 在同一行
if x > 0 {
	return y
} else {
	return x
}
if 和 switch接受初始化语句
if err := Chomd(0664); err != nil {
	fmt.Printf(err)
	return err
}
if true && true {
	fmt.Println("true")
}

goto语句  定义的标签是大小写敏感的
func myfunc(){
	i := 0
Here:
	println(i)
	i++
	goto Here
}

for 循环有三种形式
for init; condition; post {...}	和 C 中的 for 一样
for condition {...}		和 while 一样
for {...}.   			死循环

支持i++    sum += i

for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	a[i], a[j] = a[j], a[i]
}

break 提前退出循环  中止当前循环
break 后可以指定标签，退出那个循环
J: for j := 0; j < 5; j++ {
	for i := 0; i < 10; i++ {
		if i > 5 {
			break J
		}
		println(i)
	}
}

continue 跳过本次循环

range： 可用于 slice、array、string、map 和 channel
for pos, char := range "aΦx" {
	fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
}

switch 没有表达式 会匹配 true
switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		renturn c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
}

switch i {
	case 0, 1, 2: // 空的 case 体。匹配的话 不会再向下执行
	case 3:
		f() // 当 i == 0 时，f 不会被调用
}

switch i {
	case 0: fallthrough //当匹配 0 时  仍会向下执行
	case 1:
		f()  // 当 i == 0 时， f 会被调用！
}

switch i {
	case 0:
	case 1:
		f()
	default: 
		g()		// 没有匹配时调用
}

分支可以使用逗号分隔的列表
func shouldEscape(c byte) bool {
	switch c {
		case ' ', '?', '&', '=', '#', '+':   //类似 or
			return true
	}
	return false
}

字节数组进行比较
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
			case a[i] > b[i]:
				return 1
			case a[i] < b[i]:
				return -1
		}
	}
	switch {
		case len(a) < len(b):
			return -1
		case len(a) > len(b):
			return 1
	}
	return 0
}


内建函数
	close new  panic complex
	delete make recover real
	len append print imag
	cap copy println
close  用于 channel通讯 关闭channel
delete 在map中删除实例
len 和 cap
	可用于不同的类型， len 返回字符串、slice 和 数组的长度
new 各种类型内存分配
make
	内建类型（map、slice 和 channel）的内存分配
	sl := make([]int, 6)
	sl := make([]int, 6, 10)
	mp := make(map[int]string)
	ci := make(chan int)		  // 用于发送和接收整数
	cs := make(chan string)		  // 用于字符串
	cf := make(chan interface{})。// 使用空接口满足各种类型
copy	复制 slice
append 用于追加 slice
panic 和 recover
	用于异常处理机制
	一个defer执行时 发上panic,当前defer中panic之后代码不再执行。不影响其他defer的执行
print 和 println
	底层打印函数， 在不引入 fmt 情况下使用
complex、real、imag
	全部用于处理 复数

array、slice 和 map
  array
  	数组是一个有大小的数据结构
  	相同类型 内存连续分配 声明长度就固定了
  	空数组：
  		var arr [0]int
  		arr := [0]int{}
	var arr [10]int
	arr[0] = 42

	array := [5]*int{0: new(int), 1: new(int)}
	arr   := [100]int{99: 0} 	 初始化了一个100个元素的int数组

	a := [3][2]int{ [2]int{1,2}, [2]int{3,4}, [2]int{5,6} }
	a := [3][2]int{ [...]int{1,2}, [...]int{3,4}, [...]int{5,6} }
	简化后
	a := [3][2]int{ {1,2}, {3,4}, {5,6} }

	只有相同类型 相同大小 的数组才能相互赋值（复制）
	var array [3]int
	array1 := [3]int{ 1,2,3 }
	array1 = array 	  //两个不是同一个数组
  	函数间传递数组
  		var array [6]int
  		func foo(array [6]int) { ... }
  		foo(array)	// 传递的是数组的副本，需要为副本分配相同的内存空间
  		func foo1(array *[6]int) { ... }
  		foo1(&array) 	// 传递的是数组的指针, 内存空间小
  	var arrptr *[1]int 	// 没有为数组分配内存空间， arr == nil 	true
  	var arr [1]int 		// 元素默认 int的零值 0
  	array := [3]int{1,2,3}
  	array1 := [1]int{1}
  	*arrptr[2]	// 无法访问
  	*arrptr = array1	// 报错 不能这样赋值
  	*arrptr = &array    // 报错 因为 arrptr 定义为 1个元素int数组的指针
  	*arrptr = &array1    // 是可以的 指向array的内存空间
  	arr = array1 		// 复制array1 给 arr
  	多维数组
  		var array [4][2]int
  		array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
  slice
  	切片是一种占用 很小 内存空间的数据结构，对底层数组进行了抽象
  		3个字段：24字节
  			底层数组指针	长度(可访问元素的个数)	容量(允许增长(append)的最大元素个数)
  	空slice
  		var sl []int
  		sl := []int{}
  		sl := make([]int, 0)
  	创建方式
	  	1、直接声明：var slice []int 
		2、new: slice := *new([]int)
		4、make: slice :=  make([]int, 5, 10)
		3、字面量：slice := []int{1,2,3,4,5}
		5、从切片或数组“截取”:slice := array[1:5] 或 slice := sourceSlice[1:5]
	注意：
		sliceptr := new([]int)
		sliceptr := new([]int{1,2,3}) // error  new 只接受类型参数
		sliceptr := new([]int，5) // error  new 只接受类型参数
			
		*sliceptr = make([]int, 5) //能够成功
		*sliceptr = make([]int, 5);*sliceptr = make([]int, 10) //能够成功 fuzhishi
		sliceptr = make([]int, 5) //error assignment 类型不同
		*sliceptr = make([]string, 10) //error assignment 只能是 []int类型 赋值


  	sl := make([]int, 10)
  	sl := make([]int, 10, 100)  分配了100个整数的数组，slice结构指向前10个零值元素
  	slice 总是与 一个固定长度的array成对出现，其影响slice的容量和长度
  	等价于
  	var array [10]int
  	sl := array[0:n]

  	len(slice)== n; cap(slice)== m; len(array)== cap(array)== m

  	a := [...]int{1, 2, 3, 4, 5}
  	s1 := a[2:4]
  	s1[4]		//报错	访问不能超过长度
  	s2 := a[1:5]
  	s3 := a[:]
  	s4 := a[:4]
  	s5 := s2[:]  	从slice s2 创建 slice， s5 s2 同时指向 array a


  	slice := []int{10, 20, 30, 40, 50}
  	newSlice := slice[1:3]		// 长度和容量相等：3-1		都指向同一个数组
  	newSlice1 := slice[1:2:4]	// 长度：2-1， 容量：4-1	指向同一个数组

  	slice 超出 最大index 会报错 throw: index out of range

  	append 追加 新值， 返回追加后的 新的 相同类型的 slice；如果超出容量，会创建新的array 和 slice 
  	s0 := []int{0, 0}
  	s1 := append(s0, 2)
  	s2 := append(s1, 3, 5, 7)
  	s3 := append(s2, s0...)

  	copy 返回 复制元素的个数，从源slice src 复制元素 到 目标 dst 
  		复制的数量 是 len(src) 和 len(dst) 中的最小值
  	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  	var s = make([]int, 6)
  	n1 := copy(s, a[0:])	←n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
  	n2 := copy(s, s[2:])	←n2 == 4, s == []int{2, 3, 4, 5, 4, 5}

	a := []rune(s)   等价于 var a = []rune(s)
	a[i], a[j] = a[j], a[i]

	多维切片
		slice := [][]int{{10}, {100, 200}}
	函数间传递切片
		切片尺寸很小
  map
  	表现为键值对 无序	没有容量限制，可以无限增长
  	最强大的功能：能够基于键快速检索数据
  	映射的键可以是任何值。这个值的类型可以是内置的类型，也可以是结构类型，
  		只要这个值 可以使用==运算符做比较
  		bool, 数字，string, 指针, channel , 还有 interface types, structs, arrays 
  	不能作为映射的键： map、切片、函数以及 包含切片的结构类型 这些类型由于具有 引用 语义，使用这些类型会造成编译错误

  	键通过散列函数生成唯一散列值(16位)，低八位(桶-bucket的个数求余)用来选择桶，高八位存储在桶里面
  		桶结构： 
  			数组： tophash 数组-存储散列值的高八位	
  			字节数组： 存储键值对，先一次存储这个桶里面所有的键，然后在依次存储所有的值
  	map创建
  	  1、直接声明：var slice map[string]int
	  2、new: slice := *new(map[string]int)
	  3、make: slice :=  make(map[string]int)
	  	slice :=  make(map[string]int, 5) // 可以成功 参数5无效
	  4、字面量：
	  	monthdays := map[string]int{ "Jan": 31, } // 最后的 逗号 , 不能丢
	  	var map map[string]int
	  	monthdays := make(map[string]int)
	  	monthdays := map[string]int{
	  		"Jan": 31, "Feb": 28, "Mar": 31, 
	  	}
	  	mpptr := *new(map[int]string)
	  	*mpptr = make(map[string]string) // error assignment 会检查类型 map[int]string
	 

  	month, exists := monthdays["Dec"] // err 如果不存在 exists 
  	monthdays["Dec"] = 20
  	monthdays["Undecim"] = 30	//添加新的值
  	for _, days := range monthdays {
  		year += days
  	}
  	fmt.Printf("Numbers of days in a year: %d\n", year)

  	var value int
  	var present bool
  	value, present = monthdays["Jan"]	
  	存在 present 为 true；不存在 present 为 false

  	delete(monthdays, "Mar")	删除“Mar” 存不存在没有影响

  	最后的 ',' 逗号不能丢	func()的括号 () 也不能丢
  	var xs = map[int]func() int {
  		1: func() int { return 10 },
  		2: func() int { return 20 },
  		3: func() int { return 30 },
  	}

  	函数间传递映射map

b = a
值语义：b修改不会影响a的值 就是值类型
	内置类型-原始类型-按值传递： 数值类型	  字符串类型	bool类型
	数组类型: 按值传递 与C相比不同
引用语义：b修改会影响a的值，这就是引用类型
引用类型-按指针传递：切片slice 映射map 通道channel 函数 
	派生类型【指针pointer 接口interface struct】

[]byte(str) 按字节分隔成数组	uint8的别名
[]rune(str) 按字符分隔成数组	int32的别名
string(a)   byte数组/rune数组 变成 字符串
sum / float64(len(xs))	除法强制float64

单引号/双引号/反引号
	单引号用于 rune类型	rune是int32的别名 对应于 其他语言的 char类型
		var c = 's'
	  等价于
		c := 's'
	双引号用于 string类型  只能单行
		var s = "str"
		s := "str"
	反引号类似双引号，创建原生字符串字面量，可以多行

string
	string里存储的是字符按照utf8编码后的“8-bit bytes”二进制数据，再说得明确点，就是我们熟悉的byte类型
	type byte = uint8 	 golang的type alias语法，相当于给某个类型起了个别名，而不是创建了新类型，所以byte就是uint8
	type rune = int32

	for _, v := range string 
	等价于
	for _, v := range []rune(string)

	支持 + +=
	text1 := "Hello" + " GoLang"
	text1 += " Java"
	text1 := "abcdefg"

	不支持
	text1[0] = 'x'

	fmt.Println(text1[n]) 	//获取字符串索引位置为n的原始字节，比如a为97
	fmt.Println(text1[n:m]) //截取得字符串按字节索引位置为 n 到 m-1 的字符串
	fmt.Println(text1[n:])  //截取得字符串按字节索引位置为 n 到 len(s)-1 的字符串
	fmt.Println(text1[:m])  //截取得字符串按字节索引位置为 0 到 m-1 的字符串
	fmt.Println(len(text1)) //获取字符串的字节数
	fmt.Println(utf8.RuneCountInString(text1)) //获取字符串字符的个数
	fmt.Println([]rune(text1)) // 将字符串的每一个字节转换为码点值，比如这里会输出[97 98 99 100 101 102 103]
	fmt.Println(string(text1[n])) // 获取字符串索引位置为n的字符值
练习：
	avg = sum / float64(len(xs))
	进行除法，必须将值转换为 float64

	utf8.RuneCount([]byte(str))
	len([]byte(str))

	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	// fmt.Println("输出：", string(r)) #链接时 添加一个空格
	// fmt.Printf("输出：%s\n", string(r))
	fmt.Print("输出：", string(r), "\n")




































