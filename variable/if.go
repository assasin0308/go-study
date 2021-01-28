package main

import (
	"fmt"
	"io/ioutil"
)

func eval(a,b int,op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupport operator:" + op)
	}
	return result
}
func grade(score int) string {
	g := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score: %d",score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return  g
}

func forever (){
	for {
		fmt.Println("abc")
	}
}

func main() {
	const filename  = "abc.txt"
	if contents,err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)
	}
	//contents,err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n",contents)
	//}

	fmt.Println(eval(4,5,"*"))
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(90),
		grade(100),
		//grade(101),
		)
	forever()

}
