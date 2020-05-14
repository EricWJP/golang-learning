package main

import "fmt"

func read(ch1 chan int,ch2 chan bool){
	for{
		v ,ok:= <- ch1
		if ok{
			fmt.Printf("read a int is %d\n",v)
		}else{
			ch2 <- true
		}
		// fmt.Printf("read a int is %d\n",<-ch1)
	}
	ch2 <- true
}

func write(ch chan int){
	for i:=0;i<10;i++{
		 ch <- i
	}
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	// go func(){
	// 	for {
	// 		fmt.Println(11111111)	
	// 	}
	// }()
	go write(ch1)
	go read(ch1,ch2)

	<- ch2
	go func(){
		fmt.Println(11111111)
	}()
}
