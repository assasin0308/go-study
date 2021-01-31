# golang

### 1. Installation

```go
// https://studygolang.com/dl
// 国内镜像: go env -w GOPROXY=https://goproxy.cn,direct
// go env -w GO111MODULE=on
// go get -v golang.org/x/tools/cmd/goimports
```

```json
# 数据类型:
# 基本数据类型:bool,string,int
# 符合数据类型:array,slice,map,struct,指针,interface,func......
```

### 2. 变量与常量

```go
// 声明: 
var name type   
name = value
// 变量: 本质是一小块内存,用于存储某个数值,该数值可以在程序的运行过程中可以被改变
// 常量: 通变量类似,程序执行过程中,数值不可以改变
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
if 条件1 {
    // TODO 1
} else if 条件2 {
    //TODO 2
}
// 或
if a := 5;a > 2 {
    //TODO
    // 这里的a只是局部变量,外部不可使用
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
// case与switch的数值类型必须一致

func grade(score int) string {
	g := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score: %d",score))
	case score < 60:
		g = "F"
        fallthrough
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}


// break: 强制结束case执行
// fallthrough: 穿透case执行

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
var s []int // Zero value for slice is nil
s2 := make([]int,16)
s3 := make([]int,16,32)

arr := [...]int{0,1,2,3,4,5,6,7}
s := arr[2:6] //[2 3 4 5]  包前不包后
// slice本身没有数据,是对数组array的一种视图view

arr1 := [...]int{0,1,2,3,4,5,6,7}
s1 := arr1[2:6] //[2 3 4 5]
s2 := s1[3:5] // [5 6]
// slice可以向后扩展,不可以向前扩展
// s[i]不可以超越len(s),先后扩展不可以超越底层数据cap(s)
s3 := append(s2,10)
s4 := append(s3,11)
s5 := append(s4,12)
// 添加元素时如果超越cap,系统会重新分配更大的底层数组
// 由于值传递的关系,必须接受append的返回值
s = append(s,val)


```

### 11.  Map 

```go
m := map[string]string {
    "name": "assasin",
    "site": "github.com",
    "age":"25",
    "quality": "good",
}
// 遍历map: k在map中时无序的
for k,v := range m {
		fmt.Println(k,v)
	}

// map[k]V,map[K1]map[K2]V

m2 := make(map[string]int)
var m3 map[string]int


name := m["name"]
// 当获取不存在的key时返回空字符串
// 用value,ok := m[key] 判断是否存在key
if name,ok := m["nawme"]; ok {
		fmt.Println(name)
}else{
    fmt.Println("key is not exits")
}

delete(m,"name")
name,ok := m["name"]

// map的key
// map使用哈希表,必须可以比较相等
// 除了slice,map,function的内建类型都可以作为key
// struct类型不包含上述字段,也可以作为key
```

```go
// map示例: 寻找最长不含有重复字符的字串 
abcabcbb -> abc
bbbbbb -> b
pwwkew -> wke

// 对于每一个字母x
// (1) lastOccurred[x]不存在,或者 < start  -> 无需操作
// (2) lastOccurred[x] >= start -> 更新start
// (3) 更新lastOccurred[x],更新maxLength

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
```

### 12. Rune

```go
//rune相当于go的char
// 使用range遍历pos,rune对
//使用utf8.RuneCountInString获得字符数量
//使用len获取字节长度
//使用[]byte获取字节

func lengthOfNoRepeatStr(s string) int {
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


```

### 13. 结构体与方法

```go
// go语言仅支持封装,不支持继承与多态,没有class,只有struct
type treeNode struct {
    Left,Right,*TreeNode
    Value int
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}
root.left.right = createNode(10)
// 使用自定义工厂函数,返回了局部变量的地址

func (node treeNode) print() {
	fmt.Println(node.value)
}
// 显示定义和命名方法接受者

func (node *treeNode) setvalue(value int) {
	node.value = value
}
// 只有使用指针草可以改变结构体内容
// nil指针也可以调用方法

// 值接收者 指针接收者
// 要改变内容必须使用指针接收者
// 结构过大也可以考虑使用指针接收者
// 一致性:如有指针接收者,最好都是指针接收者
// 值接收者是go的特性
// 值/指针接收者均可接收值/指针
```

### 14. 封装与包

```go
// 名字一般使用CamelCase 大驼峰
// 首字母大写: public
// 首字母小写: private

// 每个目录就是一个包,main包包含可执行入口
// 为结构定的方法必须放在同一个包内,可以是不同的文件

```

### 15. 依赖管理

```shell
go mod init
go build ./...
go install ./...


go get -u go.uber.org/zap
```

### 16. 接口

```go
// 
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