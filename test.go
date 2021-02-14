package main

import "fmt"

func main() {
	var ch1 chan int
	fmt.Println(ch1) //<nil>
	fmt.Printf("%T\n",ch1) //chan int
	// nil 的channel同map一样,不能使用
	ch1 = make(chan int)
	fmt.Println(ch1)
	go func() {
		fmt.Println("子Goroutine启动......")
		data := <- ch1 // 阻塞式,从channel中读取数据
		fmt.Println("子Goroutine从channel读取的数据:",data)
	}()
	ch1 <- 100 // 阻塞式,主Goroutine向channel中写入数据
	fmt.Println("main over ....")

}








