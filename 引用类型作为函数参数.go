/*
  验证引用类型
*/
package main
import "fmt"

func main() {
	// var sl3 []int

	// var ar [10]int
	// ar[0] = 42
	// arr := [3][2]int{ {1,2}, {3,4}, {5,6} }
	// arr1 := [...]int{1, 2, 3}
	// arr2 := [...][3]int{ {1,2,3}, {3,4,5}, {5,6} }
	// arr3 := [1][3]int{}
	// sl := make([]int, 10)
	// sl1 := make([]int, 10, 100)
	// s := []int{1, 2, 3, 4}
	// s0 := []int{11, 22, 33, 44}
	// // s.update()
	// ss := append(s, s0...)
	// sss := append(ss, 333, 444)
	// fmt.Println(sss)
	// fmt.Println(sl)
	// fmt.Println(arr)
	// fmt.Println(arr1)
	// fmt.Println(arr2)
	// fmt.Println(arr3)
	// fmt.Println(ar)
	// fmt.Println(cap(sl1))
	// fmt.Println(len(sl1))
	// fmt.Println(append(sl3, 1,3,4))
	// var sl2 = [...]int{ 1,3,4,4 }
	sl2 := []int{ 1,3,4,4 }
	// fmt.Println(sl2)
	// update(sl2)
	change(&sl2)
	// sl2.change() // 未定义
	fmt.Println(sl2)
	var c complex64 = 5 + 5i
	// var c = 5 + 5i
	fmt.Println(c)
	var b byte = 234
	fmt.Printf("%T", b)
}

// 无法更改参数
// func update(s []int) {
// 	// s = append(s, 33333333)
// 	// s[2] = 333333333
// 	copy(s, s[0:2]) 
// }

func change(s *[]int) {
	// (*s)[3] = 55555		// 更改了 s
	*s = append(*s, 999) // 更改了s
	// copy((*s), (*s)[0:1])	// 无法更改 s  out of range
	// copy((*s)[4:5], (*s)[0:1])	// 无法更改 s
}

//未定义 
// func (s [4]int) change() {
// 	s[3] = 55555
// }



