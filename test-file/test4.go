package main
import ("fmt")

type SN2 struct{
    age int
    name string
}

type SN1 struct{
  age int
  name string
}

func main() {
  // sn1 := SN1{age: 11,name: "qq"}
  // sn2 := SN2{age: 11,name: "qq"} 
  // if sn1 == sn2 {
  //   fmt.Println("sn1== sn2")
  //   sn2.f()
  // }
  // sl1 := struct {
  //   age int
  //   name string
  // }{age: 11,name: "qq"}
  // sl2 := struct {
  //   name string
  //   age int
  // }{age: 11,name: "qq"} 
  // if sl1 == sl2 {
  //   fmt.Println("sl1== sl2")
  // }
  // if sl1 != sl2 {
  //   fmt.Println("sl1== sl2")
  // }
  // sm1 := struct {
  //   age int
  //   m  map[string]string
  // }{age: 11}
  // sm2 := struct {
  //   age int
  //   m  map[string]string
  // }{age: 11} 
  // if sm1 == sm2 {
  //   fmt.Println("sm1== sm2")
  // }
  // var x *int = nil
  // Foo(x)
  ttt()
  tt()
  println(&bl,bl)
  // println(&cl,cl)
  // loop: fmt.Println(11111)
  // loop: for i:=0;i<10;i++  {
  //   println(i)
  //   goto loop  
  // } 
  type MyInt1 int
  type MyInt2 = int
  var i int =9
  var i1 MyInt1 = MyInt1(i)
  var i2 MyInt2 = i
  fmt.Println(i1,i2) 


  var i11 MyUser1
  var i22 MyUser2
  i11.m1()
  i22.m2()

  my := MyStruct{}
  my.T1.m1()
  fmt.Println('x', "" != nil)
}

type T1 struct {
}
func(t T1) m1(){
  fmt.Println("T1.m1")
}

type T3 struct {
}
func(t T3) m1(){
  fmt.Println("T1.m1")
}
type T2 = T1
type MyStruct struct {
  T1
  T3
}



type User struct {
}
type MyUser1 User
type MyUser2 = User
func(i MyUser1)m1(){
  fmt.Println("MyUser1.m1")
}
func(i User)m2(){
  fmt.Println("User.m2")
}


const cl = 100
var bl   = 123
func Foo(x interface{}) {
  if x == nil {
    fmt.Println("emptyinterface")
    return
  }
  fmt.Println("non-emptyinterface")
}
func GetValue(m map[int]string, id int)(string, bool) {
  if _,exist := m[id]; exist {
    return"存在数据", true
  }
  return "nil", false
}
func tt() {
  intmap:=map[int]string{
    1:"a",
    2:"bb",
    3:"ccc",
  }
  v,err:=GetValue(intmap,3)
  fmt.Println(v,err, iii)
}

const (
  x = iota
  y
  z = "zz"
  k
  p = iota)
var iii = 2222
func ttt() {
  var iii = 11111
  fmt.Println(x,y,z,k,p, iii)
}