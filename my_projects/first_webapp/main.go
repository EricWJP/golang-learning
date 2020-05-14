package main
import ("fmt"; "net/http"; "./test"; _"./test1"; _"./test_test0")

func init() {
  fmt.Println(test.TestI)
  fmt.Printf("this is the first main init")
}
func init() {
  fmt.Printf("this is the second main init")
}
func handler(writer http.ResponseWriter, request *http.Request) {
  //fmt.Printf("%T\n", request.URL.Path[1:])
	fmt.Fprintf(writer, "hello World, %s!", request.URL.Path[1:])
}

func main() {
  fmt.Println("this is main program!!!!!!")
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}

