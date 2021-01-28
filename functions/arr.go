package main

import "fmt"

func printArr(arr *[5]int) {
	arr[0] = 100
	for _,v := range arr {
		fmt.Println(v)
	}
}
func main() {
	var arr [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}
	var grid [4][5]int
	fmt.Println(arr,arr2,arr3,grid)
	for _,v := range arr3 {
		fmt.Println(v)
	}
	//for i := 0;i < len(arr3);i++ {
	//	fmt.Println(arr3[i])
	//}
	printArr(&arr3)
	fmt.Println(arr3)

}
