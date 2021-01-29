package main

import "fmt"

func lengthOfNoRepeatStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI  >= start {
			start = lastI + 1
		}

		if lastOccurred[ch] >= start {
			start = lastOccurred[ch] + 1
		}

		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return  maxLength
}



func main() {

	fmt.Println(lengthOfNoRepeatStr("abcabcbb"))
	fmt.Println(lengthOfNoRepeatStr("bbbbbb"))
	fmt.Println(lengthOfNoRepeatStr("pwwkew"))
	fmt.Println(lengthOfNoRepeatStr("abcdef"))
	/*m := map[string]string {
		"name": "assasin",
		"site": "github.com",
		"age":"25",
		"quality": "good",
	}
	fmt.Println(m)
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m2)
	fmt.Println(m3)*/

	for k,v := range m {
		fmt.Println(k,v)
	}


	if name,ok := m["nawme"]; ok {
		fmt.Println(name)
	}else{
		fmt.Println("key is not exits")
	}

	fmt.Println("delete values")
	delete(m,"name")
	name,ok := m["name"]
	fmt.Println(name,ok)

}
