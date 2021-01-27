package main

import (
	"fmt"
	"math"
	"math/cmplx"
)
func triangle(){
	var a,b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts(){
	const filename  = "abc.txt"
	const a,b  = 3,4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename,c)
}

func enums(){
	const (
		cpp = iota
		_
		java
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,javascript,golang)
	fmt.Println(kb,mb,gb,tb,pb)
}

func main() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E,1i * math.Pi) + 1)
	triangle()
	consts()
	enums()
}

