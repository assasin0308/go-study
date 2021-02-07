package main

import "fmt"

func main() {
	// 1. 定义一个int类型的变量
	a := 10
	fmt.Println("a:",a)
	fmt.Println("a的内存地址是:",&a) // a的内存地址是: 0xc00000a0a8
    // 2. 创建一个指针变量,用于存储变量a的地址
	var p1 *int
	fmt.Println(p1) // <nil> 空指针
	p1 = &a
	fmt.Println(`p1的数值是:`,p1) // p1的数值是: 0xc00000a0a8
	fmt.Println(`p1的数值是变量的地址,该地址存储的数据是:`,*p1)
	// p1的数值是变量的地址,改地址存储的数据是: 10
	fmt.Println(`p1自己的地址是:`,&p1) // p1自己的地址是: 0xc000006030

	// 3. 操作变量,更改数值,并不会更改地址
	a = 100
	fmt.Println("此时改变a的值,a的内存地址是:",&a)
	// 4. 操作指针改变变量的值
	*p1 = 200
	fmt.Println(a)
	// 5. 指针的指针
	var p2  **int
	fmt.Println(p2) // <nil>
	p2 = &p1
	fmt.Println("p2存储p1的地址,p2的数值是:",p2) // p2存储p1的地址,p2的数值是: 0xc000006030
    fmt.Println("p2的数值,p1的地址,对应的数据",*p2) // a的地址  p2的数值,p1的地址,对应的数据 0xc00000a0a8
    fmt.Println("p2中存储的地址对应的数值,对应的数值:",**p2) // a的数值 p2中存储的地址对应的数值,对应的数值: 200


}








