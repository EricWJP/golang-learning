https://i6448038.github.io/2018/10/01/Golang-interface/
Go的interface是由两种类型来实现的：eface 和 
eface: 对于没有方法的接口实现
type eface struct {
	_type  *_type
	data unsafe.Pointer
}
_type可以认为是Go语言中所有类型的公共描述，Go语言中几乎所有的数据结构都可以抽象成_type，是所有类型的表现，可以说是万能类型，
data是指向具体数据的指针

type _type struct {
	size       uintptr 
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        *typeAlg
	// gcdata存储垃圾收集器的GC类型数据。
	// 如果KindGCProg位设置为kind，则gcdata是GC程序。
	// 否则，它是一个ptrmask位图。有关详细信息，请参见mbitmap.go
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}
例子：
b := Binary(200)
any := (interface{})(b)
fmt.Println(any)

_type ===> type(Binary)
data  ===> b的地址





所有包含方法的接口，都会使用iface结构
type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type itab struct {
	inter *interfacetype //此属性用于定位到具体interface
	_type *_type //此属性用于定位到具体interface
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

属性interfacetype类似于_type，其作用就是interface的公共描述，
类似的还有maptype、arraytype、chantype…其都是各个结构的公共描述，
可以理解为一种外在的表现信息。interfacetype源码如下：
type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}
type imethod struct {
	name nameOff
	ityp typeOff
}


