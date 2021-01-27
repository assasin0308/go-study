# golang

### 1. Installation

```go
// https://studygolang.com/dl
// 国内镜像: go env -w GOPROXY=https://goproxy.cn,direct
// go env -w GO111MODULE=on
// go get -v golang.org/x/tools/cmd/goimports
```

### 2. 变量

```go
var a,b,c bool
var s1,s2 string = "hello","assasin"
var a,b,i,s1,s2 = true,false,3,"hello","assasin"
a,b,i,s1,s2 := true,false,3,"hello","assasin"
// 可妨碍函数内,或直接放在包内
// 使用 var()集中定义变量:
var (
	 aa = 3
	 ss = "assasin"
	 bb = true
)
// 使用 := 定义变量(只能在函数内部使用)
a,b,c,s := 3,4,true,"assasin"
```

### 3. 内建变量类型

```go
bool,string
(u)int,(u)int8,(u)int16,(u)int32,(u)int64,unitptr
byte,rune
float32,float64,complex64,complex128
// 欧拉公式:
func main() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c)) // 5
	fmt.Println(cmplx.Pow(math.E,1i * math.Pi  ) + 1)
}

// 强制类型转换
func triangle(){
	var a,b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

```

### 4. 常量与枚举

```go
const filename  = "abc.txt"
const a,b  = 3,4
// const数值可以作为各种类型使用
const (
    filename  = "abc.txt"
    a,b  = 3,4
)
// 枚举
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

// 变量类型写在变量名之后
// 编译器可以推测变量类型
// 没有char,只有rune
// 原生支持复数类型
```

### 5. 流程控制

```go
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
}

// if 的条件里是可以赋值的
// if 的条件里赋值的变量作用域就在这个if语句中

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

// switch 会自动break,除非使用fallthrough 
// switch 可以没有表达式

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
	return g
}



```

### 6.

```go

```

### 7.

```go

```

### 8.

```go

```

### 9.

```go

```

### 10.

```go

```

