package main

import "fmt"

//var aa = 3
//var ss = "assasin"
//var bb = true
var (
	 aa = 3
	 ss = "assasin"
	 bb = true
)

func variable(){
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}
func variableInt(){
	var a,b int = 3,4
	var s string = "assasin"
	fmt.Println(a,b,s)
}

func variableProduction(){
	var a,b,c,s = 3,4,true,"assasin"
	fmt.Println(a,b,c,s)
}

func variableShort(){
	a,b,c,s := 3,4,true,"assasin"
	b = 5
	fmt.Println(a,b,c,s)
}

func main(){
	fmt.Print("hello assasin")
	variable()
	variableInt()
	variableProduction()
	variableShort()
	fmt.Println(aa,bb,ss)
}