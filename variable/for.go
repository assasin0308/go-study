package main

import (
	"fmt"
	"strconv"
)

func sum() int {
	sum := 0
	for i := 1;i <= 100;i++ {
		sum += i
	}
	return sum
}


func convertToBin(n int) string{
	result := ""
	for ; n > 0;n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	fmt.Println(sum())
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1101
		)
}
