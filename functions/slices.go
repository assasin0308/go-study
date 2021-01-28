package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s := arr[2:6]
	fmt.Println(s)
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])

	s1 := arr[2:]
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	s2 := arr[:5]
	s2 = s2[2:]
	fmt.Println(s2)

	arr1 := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println(arr1)
	s4 := arr1[2:6]
	s5 := s1[3:5]
	fmt.Println(s4)
	fmt.Println(s5)


}
