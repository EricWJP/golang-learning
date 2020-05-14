Go有指针，但没有指针运算
var p *int	 p是nil 也就是空(nil)指针
var i int
p = &i
*p = 8
*P++  表示 (*p)++: 首先获取指针指向的值，然后对这个值加一

Go有垃圾收集，无须担心内存分配和回收
内存分配 new/make
	new(T)	返回 *T 指向一个零值 T
	make(T)	返回初始化后的 T 	 仅适用于slice, map 和 channel
	new(T): 分配了零值填充的T类型的内存空间，并且返回其地址  *T类型 
		也就是指针（指向新分配的类型T的零值）
	make(T, args): 只能创建slice,map,channel;返回一个有初始值(非零)的T类型（注意不是*T）
		原因：这三种类型指向数据结构的引用在使用前必须被初始化
		对于这三种类型：make初始化了内部的数据结构，默认填充相应类型的零值
	示例：
		var p *[]int = new([]int)	//很少使用
		var v []int = make([]int, 100)	

		//不必要这样使用，但是可以这样用
		var p *[]int = new([]int)
		*p = make([]int, 100, 100)

		v := make([]int, 100)		//更常见
复合声明
	f := new(File) 
	f.fd = fd 
	f.name = name 
	f.dirinfo = nil 
	f.nepipe = 0 
	return f
	等价于
	f := File{fd, name, nil, 0}
	return &f
	等价于
	return &File(fd, name, nil, 0)
	等价于
	return &File{fd: fd, name: name}  // 键值对 任意顺序 并且可以省略初始化为零值的字段

	复合声明不包含任何字段 等价于 new(File) 和 &File{}
	符合声明同样可以用户创建array, slice 和 map

定义自己的类型
	类型首字母大写 是可导出的
	其成员变量/函数 首字母大写 是可导出的
	type foo int		// 不可导出
	type Str string		// 可导出
	type NameAge struct {
		Name string		// 可导出
		age int			// 不可导出
	}

	var n NameAge
	n.doSomething(2) 
	 	Go 首先会查找变量 n 的方法列表，没有找到就会
	 	再查找 *NameAge类型的方法列表，如果在这里找到 就会转化为 (&n).doSomething(2)
	// Mutex 数据类型有两个方法，Lock 和 Unlock。
	type Mutex struct { /* Mutex 字段 */ }
	func (m *Mutex) Lock() { /* Lock 实现 */ }
	func (m *Mutex) Unlock() { /* Unlock 实现 */ }

	type NewMutex Mutex 	// NewMutux 等同于 Mutex，但是它没有任何 Mutex 的方法，它的方法 是空的。

	type PrintableMutex struct { Mutex }  	// PrintableMutex 已经从 Mutex 继承了方法集合

	struct初始化
	var P Student
	var P Student = Student{"huang", 15}
	ptr := new(person)

	struct 嵌套 
		用于扩展或者修改已有类型的行为

		标识符包括变量和方法
		外部类型添加新的标识符和
		外部类型可以覆盖内部类型标识符

		标识符提升：不同包的struct，必须是首字母大写的标识符
		标识符可以被提升的标识符就像直接声明在外部类型里的标识符一样
转换
	[]byte(string)
	[]int(string)
	[]rune(string)
	string([]byte) string([]int) string([]rune)
	float32(int)
	int(float)

byteslice := []byte(mystring)
runeslice := []rune(mystring)
b := []byte{'h', 'e', 'l', 'l', 'o'}
s := string(b)
i := []rune{257, 1024, 65}
r := string(i)

unit8(int) ???
int(float32)  截断浮点数的小数部分 ???
float32(int)

type foo struct { int } 	// 匿名字段
type bar foo
var b bar = bar{1}
var f foo = b 		// error: cannot assignment 
var f foo = foo(b)	// 修改后

组合 嵌入类型
  Go不是面向对象语言，因此并没有继承
  弥补：嵌入一个类型的方式来实现

  TODO

指针运算
	这仅能工作于指向数字 int，uint等等 的指针值
	++ 仅仅定义在数字类型上，同时由于在 Go 中没有运算符重载，所以会在其他类型上失败--编译错误

示例：
type e interface {}

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

双链表 "container/list"
	l := list.New()
	l.PushBack(1)
	e := l.Front()
	e.next()
	e.Value

	type Node struct {
		Value
		prev, next *Node
	}




















