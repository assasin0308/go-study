package main

import "fmt"

func main()  {
	var a int
	var b float64
	var c string
	fmt.Println("请输入一个整数,一个浮点数,一个字符串:")
	//fmt.Printf("%p\n",&a)
	fmt.Scanln(&a,&b,&c)
	fmt.Println("a: ",a)
	fmt.Println("b: ",b)
	fmt.Println("c: ",c)
}