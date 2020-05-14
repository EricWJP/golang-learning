package main
// import "os"
import (
	"os"
	"bufio"
	"fmt"
)

// // 无缓冲
// func main() {
// 	buf := make([]byte, 1024)
// 	f, _ := os.Open("/etc/passwd")
// 	defer f.Close()
// 	for {
// 		n, _ := f.Read(buf) 	// f读取存进 buf 一次 1024字节
// 		if n == 0 { break }
// 		os.Stdout.Write(buf[:n])
// 	}
// }

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, _ := r.Read(buf)
		if n == 0 { break }
		w.Write(buf[0:n])
	}
}

type II struct {
	int
}
type I struct {
	int
}
func puts() {
	fmt.Println(1111111)
}


func (str *II) puts(x int) {
	fmt.Println(x)
}
func (str *I) puts(x int) {
	fmt.Println(x)
}
