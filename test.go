package main

import (
	"fmt"
	"time"
)

func main() {
	go hello() //启动子 Goroutine
	fmt.Println("main......")
	time.Sleep(1*time.Second)
}

func hello() {
	fmt.Println("这是一个函数,在另外一个Goroutine中执行")
}





