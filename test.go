package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 200
	}()
	select {
	case data,ok := <- ch1:
		fmt.Println("ch1中读取数据:",data,ok)
	case data := <- ch2:
		fmt.Println("ch2中读取数据:",data)
	default:
		fmt.Println("执行了default...")
	}
}




