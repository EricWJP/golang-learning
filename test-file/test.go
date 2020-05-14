package main
import ("fmt"; _"os"; "crypto/sha1")

func main() {
	// fmt.Println(os.Hostname())
	// fmt.Println(os.Getpagesize())
	// fmt.Println(os.Environ())
	// fmt.Println(os.Getuid())
	// fmt.Println(os.Geteuid())
	// i := new(I)
	// fmt.Printf("%+v\n", i)
	// fmt.Printf("%#v\n", i)
	// fmt.Printf("%T", " ")
	// var p *[]int = new([]int)
	// *p = make([]int, 100, 100)
	// pp := *p
	// *p = make([]int, 10, 100)
	// pp := *p
	// p = make([]string, 10)
	// fmt.Println(p)
	// fmt.Println(pp)
	// arr := [...]int{1,2,3}
	// barr := arr
	// barr[1] = 11
	// fmt.Println(arr)
	// fmt.Println(barr)

	// mpptr := new(map[string]int)
	// mp := map[string]int{}
	// *mpptr = map[string]int{
	//   		"Jan": 31, "Feb": 28, "Mar": 31,
	//   	}
	// fmt.Printf("%+v\n", mpptr)
	// fmt.Printf("%v\n\n\n", mpptr)
	// fmt.Printf("%+v\n", mp)

	// fmt.Printf("%v\n", mp)
	// fmt.Printf("%t\n", true)

	// // var iii interface{} = 23
	// fmt.Printf("%v\n", len("abcde我们"))
	// fmt.Printf("%v\n", len([]rune("abcde我们")))

	// _i := I{
	// 	5,
	// 	"555",
	// }
	// fmt.Println((&_i).i)
	// sss := `women` + "111"
	// fmt.Printf("%s\n%T\n", sss, sss)
	// // fmt.Printf("%v\n%T\n", bool(1), bool(1))

	// sliceptr := new([]int)
	// *sliceptr = make([]int, 5)
	// *sliceptr = make([]int, 6)
	// fmt.Printf("%v\n%T\n", &sliceptr, &sliceptr)
	// fmt.Printf("%v\n%T\n", string(1), sliceptr)

	// var x []string = nil
	// x :=[]int{ 1, 2, 3,
	// 4, 5, 6,
// 	x := map[int]string {
// 		1: "aaa",
// } 
	// x := 1
	// ++x
	// x++
	// fmt.Printf("%v\n%T\n", x, x)
	var x f_i
	x = I{ i: 1, s: "a"}
	// x = 11
	// fmt.Println(x.(I).puts(), "======")
	flag := true
	fmt.Printf("%T\n%T\n", (x.(I)), flag)

	// var m map[string]int
	// var m map[string]int = map[string]int{}
	// m["one"]= 1
	// fmt.Printf("%v\n%T\n", (&m), m)



	// var ar [3]int
	// ar[2] = 11
	// fmt.Println(ar)

	// const zero = 0.0
	// var zz float64 = zero
	// var strr = "abcdefgh"
	// strr[0] = 'x'
	// p := &strr

	// fmt.Printf("%T\n%T\n%v\n%v\n%v\n", zero + 1, zz, 1 << 5, ^1, strr)
	// pptr := NewTimesMatcher(3)
	// fmt.Printf("%v\n%T\n", pptr, pptr)
	example()

	mpp := map[int]string{}
	mpp[1] = "str"
	mpp[1] = "sss"
	fmt.Println(mpp)
}

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

	data := []byte("This page intentionally left blank.")
	fmt.Printf("sha1----%x", sha1.Sum(data))
}

type I struct {
	i int
	s string
}

func (i I) puts() (ii int)  {
	// defer func() {
	// 	if r := recover(); r!= nil {
 //            fmt.Println("Recover from r : ",r)
 //        }
	// }()
	fmt.Println("iiiiiiiiiiiiii")
	defer func(){
		fmt.Println("111111111111")
	}()
	iii := 12121221212211212

	defer func(){
		// panic("Panic 1!")
		fmt.Println("222222222222", iii)
		ii = 999999
	}()
		
	 panic("Panic 1!")
	 panic("Panic 2!")
	defer func(){
		fmt.Println("333333333")
	}()
	return  11
}

type f_i interface {
	puts() int
}


type TimesMatcher struct {
	base int
}
func NewTimesMatcher(base int) *TimesMatcher{
	return &TimesMatcher{base:base}
}