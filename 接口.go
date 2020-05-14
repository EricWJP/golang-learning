接口是一个或多个 方法签名名 的集合
只要某个类型拥有该接口的所有方法签名，就算实现该接口，无需显示声明实现了那个接口，这称为structural Typing
任何一个接口都实现了空接口
每个类型都有接口
type S struct { i int }
func (p *S) Get() int { return p.i }
func (p *S) Put(v int) { p.i = v }

// 注意 S 实现了 接口I
type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	fmt.Println(p.Get())
	p.Put(1)
}

var s S
// 之所以这里使用指针，是因为 结构类型S的定义的方法是 S指针的方法
// 当然也可以使用 传递 S 的类型值，这样 Put 方法就不会产生期望的结果
f(&s) 

Go编译器对 类型是否实现了接口 进行静态检查。
Go纯碎动态的方面，如可将一个接口类型转换到另一个。
转换的检查是在运行时进行的。
非法转换会抛出 runtime error
Go联合了 接口值、静态类型检查、运行时动态转换 以及 无须明确定义类型适配一个接口，这些使Go强大、灵活、高效和容易编写

switch 之外使用 type 是非法的
还可以使用 fmt.Printf("%T\n", p) 来查看类型 [*]包名.类型
switch p.(type) {
case *S:
case *R:
default:
}

判断一个接口类型是否实现了某个特定接口 ????
if t, ok := something.(I); ok {
	// t 是其所拥有的类型
}

空接口： 每个类型都能匹配到空接口 interface{}
	func g(something interface { }) int { 
		return something.(I).Get()
	}

方法 有接收者的函数
 	非本地类型(定义在其他包的类型)、内建类型-int类型 不能有方法， 其他任意类型都可以定义方法
 	但是这样的整数类型是可以的 新建一个拥有方法的整数类型
 	type Foo int
 	func (self Foo) Emit() {
 		fmt.Println(self)
 	}

接口类型方法
	接口定义为一个方法集合
	接收者不能定义为接口类型
	Go语言中创建指向接口的指针是没有意义的，也是非法的

接口名字
	当方法接口名 -er 后缀：  Reader Writer Buffer Formatter
	类型实现了众所周知的类型相同的方法，就使用相同的名字和声明 例如 字符串转换方法都命名为String

排序
	函数将接收一个空接口的 slice;
	使用 type switch 找到输入参数实际的类型; 然后排序;
	返回排序的 slice
	func sort(i []interface{}) { 
		switch i.(type) {
		case string:
			//...
		case int:
			//...
		}
		return 
	}
	注意 sort([]int{1,4,5}) 会报错 int不能作为接口参数

	转换接口容易，但是转换到slice开销就高了 Go不能（隐式）转换为 slice

接口嵌套
跟定义成一个接口 意义上是相同的
type Interface interface { 
	sort.Interface
	Push(x interface{})
	Pop() interface{} 
}

接口间的转换： 只有子类接口可以转换为父类接口，因为父类接口包含了子类接口
也就是说 参数是父接口类型，子接口也是可以作为参数传递的

将对象赋值给接口是，会发生拷贝，而接口内部存储的是只想这个复制品的指针，也就意味着接口无法修改状态，也无法获取指针。
func main() {
	pc := PhoneConnecter{"ipad book pro"} //实例化一个结构
	var a Connecter            //定义a为接口变量
	a = Connecter(pc)           //接口强制转换
	a.Connect()
	 
	pc.name = "Iphone 7"
	a.Connect()
	...
}
/*输出
pc.name= "ipad book pro"  Connected: ipad book pro
pc.name = "Iphone 7"    Connected: ipad book pro
*/


自省与反射
