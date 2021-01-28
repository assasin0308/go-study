package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (q, r int){
	return a / b ,a % b
}

func apply(op func(int,int) int ,a,b int) int{
	p := reflect.ValueOf(op).Pointer()
	op_name := runtime.FuncForPC(p).Name()
	fmt.Println("calling %s with args %d,%d\n",
		op_name,
		a,b)
	return op(a,b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	q, _ := div(13,3)
	fmt.Println(q)
	apply(func(a int,b int) int{
		return int(math.Pow(float64(a),float64(b)))
	},3,4)
	fmt.Println(sum(1,2,3,4,5))
}
