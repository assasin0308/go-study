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

### 6. for 循环

```go
sum := 0
	for i := 1;i <= 100;i++ {
		sum += i
	}
	return sum

// for的条件不需要括号
// for的条件可以沈略出初始条件,结束条件,递增表达式

func convertToBin(n int) string{
	result := ""
	for ; n > 0;n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
```

### 7. 函数

```go
func div(a, b int) (q, r int){
    q = a / b
    r = a % b
    return
	//return a / b ,a % b
}
//函数返回多个值可以起别名

func apply(op func(int,int) int ,a,b int) int{
	fmt.Println("calling %s with args %d,%d\n",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(),
		a,b)
	return op(a,b)
}

// 可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
```

### 8. 指针

```go
var a int = 2
var pa *int = &a
*pa = 5
fmt.Println(a)
// 指针不能参与运算
// 参数传递是值传递的方式

```

### 9. 数组

```go
// 数组是值类型数据
var arr [5]int
arr2 := [3]int{1,3,5}
arr3 := [...]int{2,4,6,8,10}
var grid [4][5]int

//遍历数组
for i := range arr3 {
		fmt.Println(arr3[i])
	}

for i,v := range arr3 {
		fmt.Println(i,v)
	}

for _,v := range arr3 {
		fmt.Println(v)
	}
//for i := 0;i < len(arr3);i++ {
//	fmt.Println(arr3[i])
//}

// [10]int 与 [20]int 是不同类型
// 调用func f(arr [10]int) 会拷贝数组
```

### 10. 切片 Slice

```go
arr := [...]int{0,1,2,3,4,5,6,7}
s := arr[2:6] //[2 3 4 5]  包前不包后
// slice本身没有数据,是对数组array的一种视图view

arr1 := [...]int{0,1,2,3,4,5,6,7}
s4 := arr1[2:6] //[2 3 4 5]
s5 := s1[3:5] // [5 6]
// slice可以向后扩展,不可以向前扩展
// s[i]不可以超越len(s),先后扩展不可以超越底层数据cap(s)
```

### 11.

```go

```

### 12.

```go

```

### 13.

```go

```

### 14.

```go

```

### 15.

```go

```

### 16.

```go

```

### 17.

```go

```

### 18.

```go

```

### 19.

```go

```

### 20.

```go

```

### 21.

```go

```

### 22.

```go

```

### 23.

```go

```

### 24.

```go

```

### 25.

```go

```

### 26.

```go

```

### 27.

```go

```

### 28.

```go

```

### 19.

```go

```

### 20.