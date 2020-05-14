/*
	验证 接收者 指针和非指针的区别
*/
package main
import (
	"fmt"
	"strconv"
)
func main() {
	var s stack 
	var ss = new(stack)
	fmt.Println(ss)
	s.push(3)
	fmt.Println(s)
}

type stack struct {
	i int
	data [10]int
}

// 调用这个 s.push(5), s不会改变，改变的是 s 的副本
// func (s stack) push(k int) {
// 	if s.i+1 > 10 {
// 		return
// 	}
// 	s.data[s.i] = k
// 	s.i++
// }
// func (s stack) pop() int {
// 	if s.i-1 < 0 {
// 		return -1
// 	}
// 	s.i--
// 	return s.data[s.i]
// }

// 创建指针s  s := new(stack)
func (s *stack)push(k int) {
	if s.i+1 > 10 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	if s.i-1 < 0 {
		return -1
	}
	s.i--
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i <= s.i; i++ {
		str = str + "[" +
				strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}














