package main

import (
	"fmt"
)

func print(s []int) {
	fmt.Println("len=%d,cap=%d\n",len(s),cap(s))
}

func main() {
	var s []int // Zero value for slice is nil
	for i := 0; i < 100;i++{
		print(s)
		s  = append(s,2 * i + 1 )
	}
	fmt.Println(s)
	s1 := []int{2,4,6,8}
	print(s1)
	s2 := make([]int,16)
	s3 := make([]int,16,32)
	print(s2)
	print(s3)

	fmt.Println("copy slice ........")
	copy(s2,s1)
	fmt.Println(s2)

	fmt.Println("delete slice.......")
	s2 = append(s2[:3],s2[4:]...)
	fmt.Println(s2)

	fmt.Println("pop from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	fmt.Println(s2)

	fmt.Println("pop from back")
	tail := s2[len(s2) -1 ]
	s2 = s2[:len(s2) - 1 ]
	fmt.Println(tail)
	fmt.Println(s2)

}
