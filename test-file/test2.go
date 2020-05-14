package main
import ("fmt")

type iff interface {
	ff()
}
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
	i := GetValue()
	// f(i)
	fmt.Println(i)

	// ir := iint{
	// 		_iint: _iint{
	// 			i: 191919191919,
	// 		},
	// 	}
	// f(ir)
	// ir.ff()
}

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

func f(i iff) {
	switch i.(type) {
		// case int:
		// 	fmt.Println("int")
		// case string:
		// 	fmt.Println("string")
	// case iint:
	// 	fmt.Println("=======iint")
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
