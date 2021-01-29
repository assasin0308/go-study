package main

import "fmt"

//rune相当于go的char
// 使用range遍历pos,rune对
//使用utf8.RuneCountInString获得字符数量
//使用len获取字节长度
//使用[]byte获取字节

func lengthOfNoRepeatStr1(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
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
	/*s := "yes我爱奥琦玮"  // UTF-8
	fmt.Println(len(s))
	for _,b := range []byte(s) {
		fmt.Printf("%X ",b)
	}
	fmt.Println(utf8.RuneCountInString(s))
	for i,ch := range []rune(s) {
		fmt.Printf("(%d %c)",i,ch)
	}*/
	fmt.Println(lengthOfNoRepeatStr1("一二三二一"))
	fmt.Println(lengthOfNoRepeatStr1("这里是天通苑"))
	fmt.Println(lengthOfNoRepeatStr1("abcabcbb"))
	fmt.Println(lengthOfNoRepeatStr1("bbbbbb"))
	fmt.Println(lengthOfNoRepeatStr1("pwwkew"))
	fmt.Println(lengthOfNoRepeatStr1("abcdef"))


}
