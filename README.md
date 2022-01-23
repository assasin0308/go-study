# golang

### 1. Installation

```go
// https://studygolang.com/dl
// 国内镜像: go env -w GOPROXY=https://goproxy.cn,direct
// go env -w GO111MODULE=on
// go get -v golang.org/x/tools/cmd/goimports

// 常用插件

git clone https://github.com/golang/tools.git tools
git clone https://github.com/golang/lint lint
git clone https://github.com/golang/image image
git clone https://github.com/golang/protobuf protobuf
git clone https://github.com/golang/sys sys


go get github.com/ramya-rao-a/go-outline
go get github.com/acroca/go-symbols
go get golang.org/x/tools/cmd/guru
go get golang.org/x/tools/cmd/gorename
go get golang.org/x/image/font
go get github.com/golang/freetype
go get github.com/josharian/impl
go get github.com/rogpeppe/godef
go get github.com/sqs/goreturns
go get github.com/cweill/gotests/gotests
go get github.com/lukehoban/go-find-references
go get github.com/disintegration/imaging
go get github.com/mojocn/base64Captcha

go get github.com/skip2/go-qrcode
go get github.com/bitly/go-simplejson
go get gopkg.in/yaml.v2
go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm
go get github.com/aliyun/aliyun-oss-go-sdk/oss
go get github.com/go-redis/redis
go get github.com/gin-gonic/gin
go get github.com/fvbock/endless
go get github.com/mozillazg/go-slugify
go get github.com/astaxie/beego
go get github.com/beego/bee
go get github.com/gorilla/websocket


go mod init gin
go mod edit -require github.com/gin-gonic/gin@latest
```

```json
# 数据类型:

# 基本数据类型: bool,string,int,float
# 复合数据类型: array,slice,map,struct,pointer,interface,channel,function......
	* 基本数据类型: 数值类型 布尔类型 字符串类型
			** 数值类型: 整数 int ,浮点 float ,复数 complex
	                     *** 整数类型:
				   **** 有符号: 正数与负数
	                                   int8,8个01码表示的整数,最高位符号位0表示正数,1表示负数,剩余7位表示数值 [-128,127]
	                                   int16 [-32768,32767]
	                                   int32 [-2147483648,2147483647]
	                                   int64 [-9223372036854775808,9223372036854775807]
	                           ****无符号: 正数(无负数)
	                                   uint8:8个01码都表示数值 [0,255]
	                                   uint16: [0,65535]
	                                   uint32: [0,4294967295]
	                                   uint64: [0,18446744073709551615]

	                           byte 可以理解为 uint8 的别名
	                           rune 可以理解为 int32 的别名

	                     *** 浮点数类型:
	                           **** float32: 32位浮点型数(单精度)
	                           **** float64: 64位浮点型数(双精度)
	                     *** 复数类型:
	                           **** complex32:
	                           **** complex64:
	            ** 布尔类型: true , false
				** 字符串类型: 使用双引号 或者 ``;
	* 复合数据类型:

# 值类型数据: bool,string,int,float,array,struct
# 引用类型数据: slice,map,function,pointer,channel

# 浅拷贝: 拷贝的是数据的内存地址,多个变量指向同一内存地址,引用类型的数据默认都是浅拷贝: slice,map,function,channel
# 深拷贝: 拷贝的是数据本身(副本),对原始数据没有影响,值类型的数据都是深拷贝: array,int,float,string,bool,struct

# Printf
%T : 打印变量的数据类型
%d : 打印正数,10进制
%f : 打印浮点数
%t : 打印bool数据
%s : 打印字符串
%v : 原样输出
%c : 打印对应的字符
%p : 与%c类似 数值对
%p : 地址
%b : 2进制表示
%o : 8进制表示
%x : 16进制表示，字母形式为小写 a-f
%X : 16进制表示，字母形式为大写 A-F



```

```go
// 数组的拷贝:
arr1 := [4]int{1,2,3,4}
arr2 := arr1 // 深拷贝
arr2[0] = 100
fmt.Println(arr2,arr1) // [100 2 3 4] [1 2 3 4]
arr3 := &arr1 // 浅拷贝
(*arr3)[0] = 200
fmt.Println(arr3,arr1) // 理论写法: &[200 2 3 4] [200 2 3 4]
arr3[0] = 200 // go语言优化: 语法糖
fmt.Println(arr3,arr1) // &[200 2 3 4] [200 2 3 4]
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


// 获取时间戳
timestamp := time.Now().Unix() // 秒
timestamp_na := time.Now().UnixNano() // 纳秒

// 获取[m,n]随机数: rand.Intn(n-m+1)+m 
// break & continue 
count := 0
	for a := 1;a <= 100;a++ {
		if a % 3 == 0 && a % 5 != 0 {
			fmt.Println(a)
			count ++
			if count == 5 {
				break
			}
		}
	}
```

### 7. 函数 function

```go
// 函数: 就是执行特定任务的代码块,具有独立功能的代码块,可以多次调用
func funcName(parameter type1,parameter type2) (outpot1 type1,output2 type2) {
    // 这里是逻辑代码
    // 返回多个值
    return value1,value2
}

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

// 可变参数函数: 可变参数相当于切片;调用时可以传入0-多个该种类型的数值,或者是直接传入该种类型的切片...
func getAdd(nums... int) int {
    sum := 0
    for _,v := range nums {
        sum += v
    }
    return sum
}
fmt.Println(getAdd(1,2,3,4)) // 10 
fmt.Println(getAdd([]int{1,2,3,4}...)) // 10 

// 注意事项1: 如果一个函数的参数是可变参数同时还有其他参数,可变参数一般放在参数列表末尾
func func2(name string,nums... int) {
    //TODO
}
// 注意事项2: 一个函数的参数列表中最多只能有一个可变参数

// 函数参数的传递:
// 1. 值类型(基本数据类型,array)的数据,默认都是值传递
// 2. 引用传递,传递的是数据的内存地址,导致多个变量指向同一块内存; slice,map

// 递归函数:
func getSum(n int) int {
	if n == 1 {
		return  1
	}
	return getSum(n - 1) + n
}


// 匿名函数(通常只能使用一次):
func (){
    fmt.Println("这是一个匿名函数")
}()

// 带参数的匿名函数:
func (a,b int) {
    fmt.Println("a: ",a,"b: ",b)
}(1,2)

// 带返回值的匿名函数:
result := func (a,b int) int {
    return  a + b
}(1,2)
fmt.Println(result)

// 用变量存储匿名函数: 
res2 := func (){
    fmt.Println("这也是一个匿名函数")
}
res2() // 调用

```

```go
// 回调函数:
// 函数func1作为另一个函数func2的参数,函数func1就是回调函数,func2就是高阶函数

func add(a,b int) int {
	return a + b
}


func sub(a,b int) int {
	if a < b {
		return 0
	}
	return a - b
}

func div(a,b int) int {
	if b == 0 || a < b {
		return 0
	}
	return a / b
}


func oper(m,n int,func1 func(int,int) int) int {
	res := func1(m,n)
	return res
}

fmt.Printf("%T\n",add) // func(int, int) int
fmt.Printf("%T\n",oper) // func(int, int, func(int, int) int) int
res1 := add(1,2)
fmt.Println(res1)
res2 := oper(10,20,add)
fmt.Println(res2)
res3 := oper(15,3,div)
fmt.Println(res3)
res4 := oper(20,10,sub)
fmt.Println(res4)
// 匿名函数做回调函数
res5 := oper(10,10,func(a,b int) int{
    //fmt.Println("a:",a,"b:",b)
    return a * b
})
fmt.Println(res5)
```

```go
// 闭包函数(将函数作为另一个函数的返回值):
func increment() func()int {
	i := 0
	fun := func() int{
		// 内层函数,一般会操作外层函数的局部变量,并且外层函数的返回值就是该内层函数.
		// 该内层函数和内层函数操作的局部变量,统称为闭包结构,closure
        // 可以延长该外层函数的布局变量生命周期
		i++
		return i
	}
	return fun // 将内层函数返回,返回的是该内层函数的内存地址
}

res1 := increment()
fmt.Printf("%T\n",res1)
res2 := res1()
fmt.Println(res2) // 1
res3 := res1()
fmt.Println(res3) // 2
```

```go
// defer函数: 一个函数的执行被延迟
// 多个defer遵循LIFO（后进先出）的顺序执行 哪怕函数或某个延迟调用发生错误，这些调用依旧会被执。
// defer 调用函数是就已经传递参数了,只是暂不执行

// 当外围函数中的语句正常执行完毕时,只有其中所有的延迟函数都执行完毕,外围函数才追真正的执行结束;
// 当执行外围函数中的return语句是,只有其中所有的延迟函数都执行完毕后,外围函数才会真正返回;
// 当外围函数中的代码引发恐慌时,只有其中所有的延迟函数都执行完毕后,该运行是恐慌才会真正被扩展至调用函数.



func f1() (result int) {
	defer func(){
		result++
	}()
	return 0
}

func f2()(r int) {
	t := 5
	defer func(){
		t = t + 5
	}()
	return t
}

func f3()(r int){
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

res1 := f1()
fmt.Println(res1) // 1
res2 := f2()
fmt.Println(res2) // 5
res3 := f3()
fmt.Println(res3) // 1
```

### 8. 指针 pointer

```go
// 指针: 是存储另一个变量的内存地址的变量
// 指针不能参与运算
// 参数传递是值传递的方式
// 指针的类型: *int,*float,*string,*array,*struct
// 指针存储的数据类型,int,float,string,array,struct
// 指针中存储的数据的地址: 指针中存储的数值
// 指针自己的地址

// 1. 定义一个int类型的变量
a := 10
fmt.Println("a:",a)
fmt.Println("a的内存地址是:",&a) // a的内存地址是: 0xc00000a0a8

// 2. 创建一个指针变量,用于存储变量a的地址
var p1 *int
fmt.Println(p1) // <nil> 空指针
p1 = &a
fmt.Println("p1的数值是:",p1) // p1的数值是: 0xc00000a0a8
fmt.Println("p1的数值是变量的地址,该地址存储的数据是:",*p1) // p1的数值是变量的地址,改地址存储的数据是: 10
fmt.Println(`p1自己的地址是:`,&p1) // p1自己的地址是: 0xc000006030

// 3. 操作变量,更改数值,并不会更改地址
a = 100
fmt.Println("此时改变a的值,a的内存地址是:",&a)

// 4. 操作指针改变变量的值
*p1 = 200
fmt.Println(a)

// 5. 指针的指针
var p2  **int
fmt.Println(p2) // <nil>
p2 = &p1
fmt.Println("p2存储p1的地址,p2的数值是:",p2) // p2存储p1的地址,p2的数值是: 0xc000006030
fmt.Println("p2的数值,p1的地址,对应的数据",*p2) // a的地址  p2的数值,p1的地址,对应的数据 0xc00000a0a8
fmt.Println("p2中存储的地址对应的数值,对应的数值:",**p2) // a的数值 p2中存储的地址对应的数值,对应的数值: 200

```

```go
// 指针作为函数参数: 引用传递
// 数组作为函数参数: 值传递
```

```go
// 数组指针: 是指针,一个数组的指针,存储的数组的地址 *[4]int
arr1 := [4]int{1,2,3,4}
fmt.Println(arr1)
// 创建一个指针,存储数组的地址
var p1 *[4]int
p1 = &arr1
fmt.Printf("%p\t",p1) // 0xc00009e120
// 根据数组指针,操作数组
(*p1)[0] = 100
p1[0] = 200
fmt.Println(arr1) // [200 2 3 4]


// 指针数组: 是数组,存储的是某种数据类型的指针 [4]*type
a := 1
b := 2
c := 3
d := 4
arr2 := [4]int{a,b,c,d}
arr3 := [4]*int{&a,&b,&c,&d}
fmt.Println(arr2) // [1 2 3 4]
fmt.Println(arr3) // [0xc00000a0a8 0xc00000a0c0 0xc00000a0c8 0xc00000a0d0]

*(arr3[0]) = 200
for i := 0;i < len(arr3);i++ {
    fmt.Println(*arr3[i])
}
fmt.Println(a) // 200 

```

```go
// 指针函数: 一个函数,该函数的返回值是一个指针
func func1() [4]int {
	arr := [4]int{1,2,3,4}
	return arr
}
func func2() *[4]int { // 指针函数
	arr := [4]int{1,2,3,4}
	return &arr
}

arr1 := func1()
arr2 := func1()
fmt.Printf("arr1的类型:%T,地址:%p,数值:%v\n",arr1,&arr1,arr1) 
// arr1的类型:[4]int,地址:0xc00011c120,数值:[1 2 3 4]
fmt.Printf("arr2的类型:%T,地址:%p,数值:%v\n",arr2,&arr2,arr2) 
// arr2的类型:[4]int,地址:0xc00011c140,数值:[1 2 3 4]
p3 := func2()
p4 := func2()
fmt.Printf("p3的类型:%T,地址:%p,数值:%v\n",p3,p3,p3) 
// p3的类型:*[4]int,地址:0xc00011c200,数值:&[1 2 3 4]
fmt.Printf("p4的类型:%T,地址:%p,数值:%v\n",p4,p4,p4) 
// p4的类型:*[4]int,地址:0xc00011c220,数值:&[1 2 3 4]




// 函数指针: 指针,指向一个函数的指针,因为func 默认可以看做一个指针
func func1(){
	fmt.Print("func1中的内容....")
}

var a func()
a  = func1
fmt.Println(a) // 0x49e210
a() // func1中的内容
```

### 9. 数组 array

```go
// 数组是值类型数据,一旦创建,大小不能改变
// var array_name [size] data_type
// var balance [10] float32
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
// 由于数组是定长数据,len(arr) 长度 === cap(arr) 容量
```

```go
// 冒泡排序: 比较相邻的两个数,小的靠前,大的靠后
func paoSort(arr []int) []int {
	nums := len(arr)
	for i := 1;i < nums;i++ {
		for j := 0;j < nums -i;j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
		//fmt.Println("第",i,"轮排序",arr)
	}
	return arr
}
```

### 10. 切片 Slice

```go
// 指向一个底层数组,有序,变长,元素可以重复
var s []int // Zero value for slice is nil
s2 := make([]int,16) // 默认容量与长度一致
s3 := make([]int,16,32)

arr := [...]int{0,1,2,3,4,5,6,7}
s := arr[2:6] //[2 3 4 5]  包前不包后
// slice本身没有数据,是对数组array的一种视图view

arr1 := [...]int{0,1,2,3,4,5,6,7}
s1 := arr1[2:6] //[2 3 4 5] [start,end)
s2 := s1[3:5] // [5 6]
// slice可以向后扩展,不可以向前扩展
// s[i]不可以超越len(s),先后扩展不可以超越底层数据cap(s)
s3 := append(s2,10)
s4 := append(s3,11)
s5 := append(s4,12)
// 添加元素时如果超越cap,系统会重新分配更大的底层数组
// 由于值传递的关系,必须接受append的返回值
s = append(s,val)

// 从已有数组上创建切片,该切片的底层数组就是当前数组
// 长度是从start到end切割的数据量,但是长度是从start到数组的末尾.
// 当向切片中添加数据时,如果没有超过容量,直接添加,如果超过容量,自动扩容(成倍增长)

```

### 11.  映射 Map 

```go
// 存储的是无序的键值对;键与值一一对应(映射项);键不能重复,如果重复,新的value或覆盖原来的value值;
// var map3 map[int]string 如果不初始化map,那么就会创建一个nil map,nil map不能用来存放键值对 map3 == nil
// 当key不存在时,会得到该value值类型的默认值,如string类型会得到空字符串,int类型会得到0,程序不会报错.
// 可以使用 value,ok := map[key] 知道key/value是否存在, ok是bool值

m := map[string]string {
    "name": "assasin",
    "site": "github.com",
    "age": "25",
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

```go
// 切片中存放map
map1 := make(map[string]string)
    map1["name"] = "assasin"
    map1["age"] = "29"
    map1["sex"] = "male"
    map1["address"] = "bj"
    map2 := make(map[string]string)
    map2["name"] = "王大狗"
    map2["age"] = "30"
    map2["sex"] = "female"
    map2["address"] = "hz"
    map3 := map[string]string{
        "name" : "李大狗",
        "age" : "45",
        "sex":"male",
        "address": "sz",
    }
    s1 := make([] map[string]string,0,3)
    s1 = append(s1,map1)
    s1 = append(s1,map2)
    s1 = append(s1,map3)
    for _,val := range s1 {
        fmt.Println(val)
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

### 13. 结构体与方法 struct

```go
// 结构体: 值类型数据,默认深拷贝
// 1. 定义一个结构体
type Person struct {
	name string
	age int
	sex string
	address string

}

// 2. 创建一个结构体类型的变量对象
var p1 Person
p1.name = `assasin`
p1.age = 29
p1.sex = `男`
p1.address = `天通苑东一区`
fmt.Println(p1) // {assasin 29 男 天通苑东一区}
// 方法二:
p2 := Person{}
fmt.Println(p2)
p2.name = "tom"
p2.age = 35
p2.sex = "女"
p2.address = "上海"
fmt.Println(p2) // {tom 35 女 上海}
// 方法三:
p3 := Person{
    name:"李晓华",
    age:48,
    sex:"男",
    address:"南京",
}
fmt.Println(p3) //{李晓华 48 男 南京}
// 方法四:
p4 := Person{"李二牛",35,"女","贵州"} // 按照定义的顺序
fmt.Println(p4) // {李二牛 35 女 贵州}
// 方法五: 内置函数 new(),go语言中专门用于创建某种类型指针的函数
p5 := new(Person)
fmt.Printf("%T,%T,\n",p1,p5) // main.Person,*main.Person,
//(*p5).namme = "jerry"
p5.name = "jacson"
p5.age = 58
p5.sex = "男"
p5.address = "江苏"
fmt.Println(p5) // &{jacson 58 男 江苏}
fmt.Printf("%p\n",p5) // 0xc00003a180
```

```go
// 结构体作为参数:

type Cat struct {
	name string
	age int
}
func changeName1(c Cat) { // c = c1 值传递,传递的是数据的副本
	c.name = "花花"
}
func changeName2(c *Cat) { // c = &c1 引用传递,传递的是地址
	c.name = "默默"
}

func  getCat1() Cat {
	c := Cat{"露娜",2}
	return c
}

func  getCat2() *Cat {
	//c := new(Cat)
	c := Cat{"加菲猫",3}
	return &c
}


c1 := Cat{"tom",1}
fmt.Println(c1) // {tom 1}
changeName1(c1)
fmt.Println(c1) // {tom 1}

changeName2(&c1)
fmt.Println(c1) // {默默 1}

c2 := getCat1()
c3 := getCat1()
fmt.Println(c2,c3) // {露娜 2} {露娜 2}
c2.name = "jerry"
fmt.Println(c2,c3) // {jerry 2} {露娜 2}

c4 := getCat2()
c5 := getCat2()
fmt.Println(*c4,*c5) // &{加菲猫 3} &{加菲猫 3}
c4.name = "小花猫"
fmt.Println(*c4,*c5) // &{小花猫 3} &{加菲猫 3}


```

```go
// 匿名结构体:没有名字的结构体,往往同时创建该匿名结构体的对象
p1 := struct {
		name string
		age int
	}{
		name: "张三",
		age: 16,
	}
fmt.Println(p1.name,p1.age)

// 匿名字段:一个结构体的字段没有名字,默认使用类型作为字段名.同一类型只能定义一次!
type student struct {
	string // 匿名字段
	int // 匿名字段
}

s1 := student{"张三",18}
fmt.Println(s1) // {张三 18}
fmt.Println(s1.string,s1.int) // 张三 18

```

```go
// 结构体的嵌套: 

// 聚合关系: 一个类作为另一个类的属性 has a 
type A struct{}
type B struct{
    a  A // 聚合关系
}
// 1. 定义一个书的结构体
type Book struct {
	bookname string
	price float64
	author string
}
// 2. 定义一个人的结构体
type Person struct {
	name string
	age int
	book Book // 聚合关系 -> has a 关系
}

b1 := Book{}
b1.bookname = "金瓶梅里没有梅"
b1.price = 45.8
b1.author = "田中天"

p1 := Person{}
p1.name = "王二狗"
p1.age = 29
p1.book = b1 // 将结构体作为属性

fmt.Printf("姓名:%s,年龄是:%s,书名:%s,价格:%.2f,作者:%s\n",
           p1.name,p1.age,
           p1.book.bookname,
           p1.book.price,
           p1.book.author,
          )
p1.book.bookname = "金瓶梅里没有梅2"
fmt.Println(p1)
fmt.Println(b1)

p2 := Person{
    name: "老王",
    age:35,
    book: Book{bookname:"金瓶梅里没有梅3", price:48.2, author:"assasin"},
}
fmt.Println(p2.name,p2.age)
fmt.Println("\t",p2.book.bookname,p2.book.price,p2.book.author)

p3 := Person{
    name: "jeryy",
    age: 45,
    book: Book{
        bookname: "go语言怎么炼成的",
        price: 55.9,
        author: "王建",
    },
}
fmt.Println(p3.name,p3.age)
fmt.Println("\t",p3.book.bookname,p3.book.price,p3.book.author)

p4 := Person{"李小花",68,Book{"十万个为啥 ? ",45.2,"张晓春"}}
fmt.Println(p4.name,p4.age)
fmt.Println("\t",p4.book.bookname,p4.book.price,p4.book.author)


// 实现一个学生对象有多本书(创建一个装书的容器):
type Book struct {
 	bookname string
 	price float64
 }

type Student struct {
	name string
	age int
	books  []*Book
}

// 1.创建多个书对象
b1 := Book{"孙子兵法",120.8}
b2 := Book{"天局",28.30}
b3 := Book{"易经经",88.9}

// 2.创建一个装书的容器 切片
bookcase := make([]*Book,0,10)
// 3.将书放入"书架"
bookcase = append(bookcase,&b1,&b2,&b3)
// 4.创建学生对象
s1 := Student{"王大狗",19,bookcase}
// 5.获取对象属性
fmt.Printf("姓名:%s,年龄:%d\n",s1.name,s1.age)
//fmt.Println(s1.books) //[0xc0000044c0 0xc0000044e0 0xc000004500]
for i := 0;i < len(s1.books);i++ {
    p := s1.books[i] // 书对象的地址
    fmt.Printf("\t第 %d 本书名:%s,价格:%.2f\n",
               i+1,
               //(*p).bookname,
               p.bookname,
               //(*p).price,
               p.price,
              )
}

// s1 输出:
姓名:王大狗,年龄:19
第 1 本书名:孙子兵法,价格:120.80
第 2 本书名:天局,价格:28.30
第 3 本书名:易经经,价格:88.90

s2 := Student{"李小花",18,make([]*Book,0,10)}
s2.books = append(s2.books,
                  &Book{"红露梦",77.8},
                  &Book{"西游记",98.2},
                 )
fmt.Printf("姓名:%s,年龄:%d\n",s2.name,s2.age)
//fmt.Println(s2.books)
for i := 0;i < len(s2.books);i++ {
    p := s2.books[i] // 书对象的地址
    fmt.Printf("\t第 %d 本书名:%s,价格:%.2f\n",
               i+1,
               p.bookname,
               p.price,
              )
}

// s2输出:
姓名:李小花,年龄:18
第 1 本书名:红露梦,价格:77.80
第 2 本书名:西游记,价格:98.20

// 升级版: 创建一个容器,存储学生对象

ss1 := make([]*Student,0,10)
ss1 = append(ss1,&s1,&s2)
for i := 0;i < len(ss1);i++{
    fmt.Printf("第%d个学生的信息:\n",i+1)
    fmt.Printf("学生姓名:%s,年龄:%d\n",ss1[i].name,ss1[i].age)
    bookshelf := ss1[i].books //切片
    if len(bookshelf) == 0 {
        fmt.Println("\t\t该生不看书")
    }else{
        for j := 0;j < len(bookshelf);j++ {
            p := bookshelf[j]
            fmt.Printf("\t\t第 %d 本书名:%s,价格:%.2f\n",
                       j+1,
                       p.bookname,
                       //ss1[i].books[j].bookname,
                       p.price,
                       //ss1[i].books[j].price,
                      )
        }
    }
}

// 输出:
/*
第1个学生的信息:
学生姓名:王大狗,年龄:19
		第 1 本书名:孙子兵法,价格:120.80
		第 2 本书名:天局,价格:28.30
		第 3 本书名:易经经,价格:88.90
第2个学生的信息:
学生姓名:李小花,年龄:18
		第 1 本书名:红露梦,价格:77.80
		第 2 本书名:西游记,价格:98.20

*/

// 继承关系: 一个类作为另一个类的子类:子类,父类 -> is a
type A struct{}
type B struct{
     A // 继承
}
// 一个类(子类,派生类,subClass)继承另一个类(父类,超类,基类,superClass)
// 子类可以直接访问父类已有的属性和方法
// 子类也可以直接创建自己的属性和方法
// 子类也可以重写父类已有的方法

type Person struct {
	name string
	age int
}

type Student struct {
	Person // 继承 Person父类
	school string // 子类的新增属性
}

// 1. 创建父类对象
p1 := Person{name:"张三",age: 30}
fmt.Println(p1.name,p1.age) // 张三 30
// 2. 创建子类对象
var s1 Student
s1.name = "李四" // 子类对象访问父类属性
s1.age = 18
s1.school = "Peking-University"
fmt.Println(s1.name,s1.age,s1.school) // 李四 18 Peking-University

s2 := Student{Person{"王五",17},"清华大学"}
fmt.Println(s2.name,s2.age,s2.school) // 王五 17 清华大学

s3 := Student{Person:Person{"Rose",25},school:"新东方"}
fmt.Println(s3.Person.name,s3.Person.age,s3.school) // Rose 25 新东方
fmt.Println(s3.name,s3.age,s3.school) // [提升字段]  // Rose 25 新东方

```



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
// 要改变内容必须使用指针+ 接收者
// 结构过大也可以考虑使用指针接收者
// 一致性:如有指针接收者,最好都是指针接收者
// 值接收者是go的特性
// 值/指针接收者均可接收值/指针
```

### 14. 方法 method

```go
// 定义:某个类别的行为功能,需要有接受者调用
func (t Type) methodName(parameter list){
    //TODO
}

// 1. 定义一个工人类
type Worker struct {
    name string
    age int
    sex string
}
type Cat struct {
    color string
    age int
}
// 2. 定义类的行为方法
func (w Worker) work() {
   fmt.Println(w.name,"在干活")
}

func (w *Worker) rest(){
   fmt.Println(w.name,"在休息")
}

func (c Cat) eat(){
   fmt.Println(c.color,"的猫在吃草")
}

func (w *Worker) printInfo() {
   fmt.Printf("工人姓名:%s,年龄:%d,性别:%s\n",w.name,w.age,w.sex)
}

func (c Cat) printInfo(){
   fmt.Printf("猫咪的颜色:%s,猫咪的年龄:%d",c.color,c.age)
}


// 3. 创建对象
w1 := Worker{name:"王二狗",age: 30,sex:"男"}
w1.work()  // 对象调用方法 // 王二狗 在干活
w1.printInfo()

w2 := Worker{"李小花",45,"女"}
w2.work()
w2.rest()
w2.printInfo()

w3 := &w2
w3.work()
w3.rest()
w3.printInfo()

c1 := Cat{color:"白色",age: 3}
c1.eat()  // 白色 的猫在吃草
c1.printInfo()


// 继承中方法的使用:
// 1. 结构体
type Person struct {
    name string
    age int
}

type Student struct {
    Person
    school string
}
// 2. 方法
func (p Person) eat() {
   fmt.Println("父类的方法,吃窝窝头")
}

func (s Student) study() {
   fmt.Printf("学生在学习...")
}
// 子类重写父类方法
func (s Student) eat() {
 fmt.Println("父类的方法,吃炸鸡,喝啤酒...")
}


p1 := Person{name:"王二狗",age: 25}
fmt.Println(p1.name,p1.age) // 父类对象访问父类属性
s1 := Student{Person{"rose",18},"清华大学"}
fmt.Println(s1.name,s1.age) // 子类对象访问父类属性
fmt.Println(s1.school) // 子类访问自己新增属性
p1.eat() // 父类对象访问父类方法
s1.eat() // 子类对象访问父类方法,存在重写,会访问子类重写的方法
s1.study() // 子类对象访问子类新增方法

```

### 15. 接口 Interface

```go
// 定义:"接口定义对象的行为",它只指定对象应该什么,实现这种行为的方法(实现细节)是针对对象的.
// Go中,接口是一组方法签名.当类型为接口中的所有方法提供定义时,他被称为实现接口.它与OOP非常相似.接口指定了类应该具有的方法.类型决定了如何实现这些方法.
// 它把所有的具有共性的方法定义在一起,任何其他类型只要实现了这些方法就是实现了这个接口.
// 接口定义了一组方法,如果某个对象实现了某个接口的所有方法,则此对象就实现了接口.

type interface_name interface {
	method_name1 [return_type]
	method_name2 [return_type]
	method_name3 [return_type]
	...
	method_namen [return_type]
}


// 1. 定义接口
type USB interface {
	start()
	end()
}
// 2. 定义实现结构体
type Mouse struct {
	name string
}
// 3. 定义实现方法
func (m Mouse) start() {
	fmt.Println(m.name,"鼠标准备就绪,可以开始工作了,点点点....")
}
func (m Mouse) end() {
	fmt.Println(m.name,"鼠标结束工作,可以安全退出了....")
}
// 4. 定义第二个实现结构体
type FlashDisk struct {
	name string
}
// 5. 定义实现方法
func (f FlashDisk) start() {
	fmt.Println(f.name,"U盘准备就绪,可以开始工作了....")
}
func (f FlashDisk) end() {
	fmt.Println(f.name,"U盘结束工作,可以安全退出了....")
}

func testInterface(usb USB) {
	usb.start()
	usb.end()
}

m := Mouse{name: "罗技"}
f := FlashDisk{name:"台积电"}
var usb USB
//var usb USB = Mouse{name:"罗技"}
//var usb USB = FlashDisk{name:"台积电"}

//usb = m // 创建该接口的任意实现类对象
usb = f
//fmt.Println(usb.name) //接口对象不能访问实现类的属性
usb.start()
usb.end()
fmt.Println("-------------------------------")
testInterface(m)
//testInterface(f)

// 注意点: 1.当需要接口类型的对象时,那么可以使用任意实现类对象代替
//		  2.接口对象不能访问实现类的属性
//        3.一个函数如果接受接口类型作为参数,那么实际上可以传入该接口的任意实现类对象作为参数
//        4.定义一个类型为接口,那么实际上可以赋值任意实现类对象 
```

```go
// 定义一个接口: Shape形状,方法: 面积,周长
// 实现类: 矩形(长宽),圆(半径),三角形(三边长,a,b,c)
// 测试接口函数: testShape()

// 1. 定义接口
type Shape interface {
	peri() float64 // 周长
	area() float64
}
// 2. 定义实现类
type Triangle struct {
	a,b,c float64
}
// 定义实现方法
func (t Triangle) peri() float64{
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	// 海伦公式
	peri := t.peri() / 2  // 半周长
	s := math.Sqrt(peri * (peri - t.a) *(peri - t.b) * (peri - t.c))
	return s
}
// 定义圆的类
type Circle struct {
	radius float64
}
// 定义实现方法
func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius,2)*math.Pi
}

func testShape(s Shape) {
	fmt.Println("周长:",s.peri(),"面积:",s.area())
}

// 对象转型方法一
func getType(s Shape) {
	// 获取类型: instance,ok := 接口对象.(实际类型)
	// 如果该接口对象是对应的实际类型,那么instance就是转型之后的对象,OK是true
	if ins,ok := s.(Triangle);ok {
		fmt.Println("是三角形,三边长是:",ins.a,ins.b,ins.c)
	}else if ins,ok := s.(Circle);ok{
		fmt.Println("是圆形,半径是:",ins.radius)
	}else {
		fmt.Println("不知道是啥形状...")
	}
}

// 对象转型方法二
func getType2(s Shape) {
	// 结合switch语句  接口对象.(type)
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("是三角形,三边长是:",ins.a,ins.b,ins.c)
	case Circle:
		fmt.Println("是圆形,半径是:",ins.radius)
	
	}
}


var t1 Triangle
t1 = Triangle{3,4,5}
fmt.Println(t1.peri())
fmt.Println(t1.area())

var s1 Shape
s1 = t1
fmt.Println(s1.peri())
fmt.Println(s1.area())

var c1 Circle
c1 = Circle{4}
fmt.Println(c1.peri())
fmt.Println(c1.area())
fmt.Println(c1.radius)

var s2 Shape = Circle{4}
fmt.Println(s2.peri())
fmt.Println(s2.area())

testShape(t1)

// 定义一个接口类型的数组
arr := [4]Shape{t1,s1,s2,c1}
fmt.Println(arr)

fmt.Println(strings.Repeat(`-`,50))

getType(c1)
getType2(s2)
```

```go
// 空接口:没有任何方法,可以将任意类型作为该接口的实现.
type A interface {

}

type Cat struct {
	name string
	age int
}
type Person struct {
	name string
	sex string
}
// 将匿名空接口作为参数:表示该函数可以接受任意类型的数据
func test(a interface{}) {
	fmt.Println(a)
}

func test2(slice []interface{}) {
	for i := 0;i < len(slice); i++ {
		fmt.Println("第",i+1,"个数据:")
		switch ins := slice[i].(type) {
		case Cat:
			fmt.Println("\t cat对象:",ins.name,ins.age)
		case Person:
			fmt.Println("\t person对象:",ins.name,ins.sex)
		case int:
			fmt.Println("\t int对象:",ins)
		case string:
			fmt.Println("\t string对象:",ins)
		}
	}
}




var a1 A = Cat{name:"花猫",age: 5}
var a2 A = Person{name:"王二狗",sex: "男性"}
var a3 A = "hello"
var a4 A = 100
test(a1)
test(a2)
test(a3)
test(a4)
// 定义一个map string作key,任意类型作为值
map1 := make(map[string]interface{})
map1["name"] = "assasin"
map1["age"] = 25
fmt.Println(map1) // map[age:25 name:assasin]
// 定义一个切片,能够存储任意类型的数据
slice1 := make([]interface{},0,10)
slice1 = append(slice1,a1,a2,a3,a4)
fmt.Println(slice1) // [{花猫 5} {王二狗 男性} hello 100]
test2(slice1)
//第 1 个数据:
//cat对象: 花猫 5
//	第 2 个数据:
//person对象: 王二狗 男性
//	第 3 个数据:
//string对象: hello
//	第 4 个数据:
//int对象: 100

```

### 16. error处理

```go
// 如下: 
f,err := os.Open("./template1.md")
if err != nil {
    fmt.Println("err:",err)
    return
}
fmt.Println(f.Name())


// 设计一个函数,验证一个人的年龄是否合法
func checkAge(age int) error {
	if age < 0 {
		return errors.New("年龄不合法") // string error
	}
	fmt.Println("年龄是: ",age)
	return nil
}


// 1. 创建一个 error 对象
err1 := errors.New("这就是自定义错误的信息")
fmt.Println(err1.Error())
fmt.Printf("%T\n",err1)

err2 := checkAge(30)
if err2 != nil {
    fmt.Println(err2.Error())
    return
}
// 2. 另一种创建 error 的方法
err3 := fmt.Errorf("这是错误信息码:%d",100)
fmt.Println(err3)

s2 := "100"
ll,err := strconv.ParseInt(s2,10,64)
if err != nil {
    fmt.Println(err.Error())
}
fmt.Println(ll)

var l2 int64 = 100
// int64 转字符串
s4 := strconv.FormatInt(l2,10)
s3 := "hello" + s4
fmt.Println(s3) // hello100

// 构建字符串
s1 := fmt.Sprint("helloworld",25,3.145,"你好")
fmt.Println(s1)
s5 := fmt.Sprintf("name:%s,age:%d","王二狗",30)
fmt.Println(s5)
// 读取键盘输入
//fmt.Scanln()
//fmt.Scanf()
// 拼接字符串
//fmt.Sprint()
//fmt.Sprintf()
//fmt.Sprintln()
```

### 17. 自定义error

```go
// 求矩形面积,若长宽为负数,返回自定义错误信息
type errorRectangle struct {
	msg string
	length,width float64
}
func (e *errorRectangle) Error() string {
	return fmt.Sprintf("宽是:%.2f,长是:%.2f,错误信息是:%s", e.width,e.length,e.msg)
}

func getArea(width,length float64) (float64,error) {
	errorMsg := ""
	if width < 0 {
		errorMsg = "宽度是负数"
	}
	if length < 0 {
		if errorMsg == "" {
			errorMsg = "长度是负数"
		}else{
			errorMsg += ",长度也是负数"
		}
	}
	if errorMsg != "" {
		return 0,&errorRectangle{errorMsg,length,width}
	}
	area :=  length * width
	return area,nil
}

res,err := getArea(-6,-4)
if err != nil {
    fmt.Println(err.Error())
}else{
    fmt.Println("面积是:",res)
}
```

```go
// 练习:

// 1. 求(三维)两点之间的距离
type Point struct {
	x,y,z float64
}
func (p Point) printInfo() {
	fmt.Printf("x轴:%.2f,y轴:%.2f,z轴:%.2f\n",p.x,p.y,p.z)
}
func (p Point) getDistance(x1,x2,y1,y2,z1,z2 float64) float64{
	dis := math.Sqrt(math.Pow(x1-x2,2) + math.Pow(y1-y2,2) + math.Pow(z1-z2,2))
	return dis
}

func (p Point) getDistance2(p1,p2 Point) float64{
	dis := math.Sqrt(math.Pow(p1.x - p2.x,2) + math.Pow(p1.y - p2.y,2) + math.Pow(p1.z - p2.z,2))
	return dis
}

func (p Point) getDistance3(p2 Point) float64 {
	dis := math.Sqrt(math.Pow(p.x - p2.x,2) + math.Pow(p.y - p2.y,2) + math.Pow(p.z - p2.z,2))
	return dis
}

p1 := Point{2,4,3}
p2 := Point{0,0,0}
p1.printInfo()
p2.printInfo()
res1 := p1.getDistance(p1.x,p2.x,p1.y,p2.y,p1.z,p2.z)
fmt.Println(res1) // 5.385164807134504
res2 := p1.getDistance2(p1,p2)
fmt.Println(res2) // 5.385164807134504
res3 := p1.getDistance3(p2)
fmt.Println(res3) // 5.385164807134504
res4 := p2.getDistance3(p1)
fmt.Println(res4) // 5.385164807134504

// 2. 定义一个学生类,具有6个属性:姓名,年龄,性别,英语成绩,语文成绩,数学成绩,提供3个方法,求总分,求平均分,打印信息 
type Student struct {
	name string
	age int
	sex string
	englishScore,chineseScore,mathScore float64
}

func (s Student) getSum() float64 {
	sum := s.englishScore + s.chineseScore + s.mathScore
	return sum
}

func (s Student) getAvg() float64 {
	return s.getSum() / 3
}

func (s Student) printInfo() {
	fmt.Printf("姓名:%s,年龄:%d,性别:%s\n",s.name,s.age,s.sex)
	fmt.Printf("英语成绩:%.2f,语文成绩:%.2f,数学成绩:%.2f\n",s.englishScore,s.chineseScore,s.mathScore)
	fmt.Printf("总分:%.2f,平均分:%.2f",s.getSum(),s.mathScore)
}

s1 := Student{"王二狗",25,"男",69.7,88.2,85.2}
fmt.Println(s1.getSum()) // 243.10000000000002
fmt.Println(s1.getAvg()) // 81.03333333333335
s1.printInfo()
// 姓名:王二狗,年龄:25,性别:男
// 英语成绩:69.70,语文成绩:88.20,数学成绩:85.20
// 总分:243.10,平均分:85.20

// 3. 定义一个IEngine接口,3个方法,start(),speedup(),stop(),表示启动,加速,停止.定义2个结构体实现该接口:YAMAHA和HONDA,实现方式分别是 YAMAHA: 启动:60,加速:80,停止:0;HONDA 启动:40,加速:120,停止:0.定义一个Car结构体,将IEngine作为字段,再提供一个car的方法,testIEngine(),用于测试该小汽车的发动机.

// 1. 定义发动机接口
type IEngine interface {
	start()
	speedup()
	stop()
}
// 2.定义实现类
type YAMAHA struct {
	name string
}
func (y YAMAHA) start(){
	fmt.Println(y.name,"启动,速度60")
}

func (y YAMAHA) speedup(){
	fmt.Println(y.name,"加速,速度80")
}

func (y YAMAHA) stop(){
	fmt.Println(y.name,"停止,速度0")
}

type HONDA struct {
	name string
}

func (h HONDA) start(){
	fmt.Println(h.name,"启动,速度40")
}
func (h HONDA) speedup(){
	fmt.Println(h.name,"加速,速度120")
}
func (h HONDA) stop(){
	fmt.Println(h.name,"停止,速度0")
}

type Car struct {
	iEngine IEngine // 接口类型作为字段
}

func (c Car) testIEngine() {
	c.iEngine.start()
	c.iEngine.speedup()
	c.iEngine.stop()
}

c1 := Car{}
c1.iEngine = YAMAHA{"雅马哈"}
c1.testIEngine()

c2 := Car{}
c2.iEngine = HONDA{"本田"}
c2.testIEngine()
```

### 18. Panic & Recover

```go
func funcA () {
	fmt.Println("这是一个函数funcA")
}
func funcB(){
	defer func() {
		// msg 就是panic传入的数据
		if msg := recover();msg != nil {
			fmt.Println(msg,"恢复啦......")
		}
	}()
	for i := 0;i <=10;i++ {
		fmt.Println("i:",i)
		if i == 5 {
			// 让程序中断
			panic("funcB panic......")
		}
	}
	// 当外围函数中的代码引发运行恐慌时,只有其中所有的延迟函数都执行完毕后,该运行恐慌时才会真正被扩展至调用函数
}

func funcC() {
	defer func() {
		fmt.Println("funcC 的延迟函数...")
		recover()
	}()
	fmt.Println("这是funcC 函数")
	panic("funcC 玩完了...")
}

funcA()
funcB()
funcC()
fmt.Println("func over......")
```

### 19. Time

```go
// 1. 获取当前时间
t1 := time.Now()
fmt.Printf("%T\n",t1) // time.Time
fmt.Println(t1) // 2021-02-12 23:30:59.2155127 +0800 CST m=+0.001994801

// 2. 获取指定时间
t2 := time.Date(2008,7,15,16,45,28,0,time.Local)
fmt.Println(t2) // 2008-07-15 16:45:28 +0800 CST

// 3. time -- string之间转换
// 模板的日期必须是固定的: 6-1-2-3-4-5
s1 := t1.Format("2006年1月2日 15:04:05")
fmt.Println(s1) // 2021年2月12日 23:53:42
s2 := t1.Format("2006年1月2日")
fmt.Println(s2) // 2021年2月12日
s3 := "1999年10月10日"
t3,err := time.Parse("2006年1月2日",s3)
if err != nil {
    fmt.Println("err: ",err.Error())
}
fmt.Println(t3) // 1999-10-10 00:00:00 +0000 UTC

// 4. 时间戳
t4 := time.Date(1970,1,1,1,0,0,0,time.UTC)
timestamp1 := t4.Unix() // 秒的差值
fmt.Println(timestamp1)
timestamp2 := t1.Unix()
fmt.Println(timestamp2) // 当前时间戳
timestamp3 := t4.UnixNano() // 纳秒的差值
fmt.Println(timestamp3)
timestamp4 := t1.UnixNano()
fmt.Println(timestamp4)

// 5. 根据当前时间获取指定的日期
year,month,day := t1.Date()
fmt.Println(year,month,day) // 2021 February 13
hour,minite,second := t1.Clock()
fmt.Println(hour,minite,second) // 0 3 29

year2 := t1.Year()
fmt.Println("年:",year2) // 年: 2021
fmt.Println(t1.YearDay()) // 一年中第几天  44
month2 := t1.Month()
fmt.Println("月: ",month2) // 月:  February
fmt.Println("日: ",t1.Day()) // 日:  13
fmt.Println("小时: ",t1.Hour()) // 小时:  0
fmt.Println("分钟: ",t1.Minute()) // 分钟:  7
fmt.Println("秒: ",t1.Second()) // 秒:  0
fmt.Println(t1.Weekday()) // 星期几 Saturday
fmt.Println(t1.ISOWeek()) // 返回 年,第几周  2021 6

// 6. 时间间隔
t5 := t1.Add(time.Nanosecond) // duration时间间隔之后
//t6 := t1.Add(time.Minute) // duration时间间隔之后
fmt.Println(t1)
fmt.Println(t5)
fmt.Println(t1.Add(24*time.Hour))

t6 := t1.AddDate(1,0,0)
fmt.Println(t6)

t7 := t5.Sub(t1) // 两个time的时间差值
fmt.Println(t7)

// 7. 睡秒
time.Sleep(3 * time.Second) // 当前程序进入睡眠状态
fmt.Println("main...over....")
```

### 20.  文件操作 & Seek

```go
// 1. 获取文件信息
fileinfo,err := os.Stat(`E:\golang\README.md`)
if err != nil {
    fmt.Println(err.Error())
    return
}
fmt.Printf("%T\n",fileinfo) // *os.fileStat
fmt.Println(fileinfo) // &{README.md 32 {2947189995 30867544} {3558147031 30867886} {376089464 30867886} 0 44404 0 0 {0 0} E:\golang\README.md 0 0 0 false}
fmt.Println(fileinfo.Name()) // 文件名: README.md
fmt.Println(fileinfo.IsDir()) // 是否目录 bool: false
fmt.Println(fileinfo.Size()) // 文件大小: 44404
fmt.Println(fileinfo.Mode()) // 文件权限: -rw-rw-rw-
fmt.Println(fileinfo.ModTime()) // 文件修改时间: 2021-02-13 10:15:24.274572 +0800 CST


// 2. 文件操作
filename1 := `E:\golang\cases.md`
filename2 := `template.md`
// 是否绝对路径
fmt.Println(filepath.IsAbs(filename1)) // true
fmt.Println(filepath.IsAbs(filename2)) // false
// 获取绝对路径,返回路径,错误
fmt.Println(filepath.Abs(filename1)) //E:\golang\cases.md <nil>
fmt.Println(filepath.Abs(filename2)) //E:\golang\template.md <nil>
// 获取父级(上一级)目录
fmt.Println(path.Join(filename2))


// 3. 创建文件夹,若存在,则报错
// 仅创建一层目录
err := os.Mkdir(`E:\golang\assasin`,os.ModePerm)
if err != nil {
    fmt.Println(err.Error())
}else {
    fmt.Println("文件夹创建成功...")
}
// 可创建错层目录
err1 := os.MkdirAll(`E:\golang\assasin\assasin01`,os.ModePerm)
if err1 != nil {
    fmt.Println(err1.Error())
}else {
    fmt.Println("文件夹创建成功...")
}


// 4. 创建文件,若已经存在,则覆盖
file,err := os.Create(`E:\golang\assasin\assasin01\abc.txt`)
if err != nil {
    fmt.Println(err.Error())
}else {
    fmt.Println(file)
}

// 5. 打开文件
   // os.Open 只读模式
file,err := os.Open(`E:\golang\assasin\assasin01\abc.txt`)
if err != nil {
    fmt.Println("\t",err.Error())
} else {
    fmt.Println(file)
}
// 参数3个: 文件名称,文件打开方式,文件权限
file1,err := os.OpenFile(`E:\golang\assasin\assasin01\abc.txt`,os.O_WRONLY|os.O_CREATE,os.ModePerm)
// 关闭文件
file1.Close()


// 6. 删除文件夹与文件
// 删除文件, os.Remove,若是文件夹,必须空
err := os.Remove(`E:\golang\assasin\assasin01`)
fmt.Println(err)
// 删除所有
err1 := os.RemoveAll(`E:\golang\assasin`)
fmt.Println(err1)


// 7. 文件的读取
file,err := os.Open(`E:\golang\assasin\assasin01\123.txt`)
if err != nil {
    fmt.Println("打开文件有错误",err.Error())
    return
}

// 优化 循环读取
bs := make([] byte,4,4)
n := -1
for {
    n,err = file.Read(bs) // --> n,err
    // 从文件开始读取数据,存入byte切片中,返回值n是本次实际读取的数据量,如果读取到文件末尾,n =0,err = EOF
    if n == 0 || err == io.EOF {
        fmt.Println("读取到文件末尾了,结束文件操作...")
        break
    }
    fmt.Println(string(bs[:n]))
}
/*
	fmt.Println("第一次读取....")
	n,err := file.Read(bs) // 从file中对应的文件中读取最多len(bs)个数据,存入bs数组中,n是实际读取的数量
	fmt.Println(n) // 4
	fmt.Println(err) // nil
	fmt.Println(bs) //  [97 98 99 100]
	fmt.Println(string(bs)) // abcd
	// 二次读取
	fmt.Println("第二次读取....")
	n,err = file.Read(bs)
	fmt.Println(n) // 4
	fmt.Println(err) // <nil>
	fmt.Println(bs) //[101 102 103 104]
	fmt.Println(string(bs[:n])) // efgh
	// 三次读取
	fmt.Println("第三次读取....")
	n,err = file.Read(bs)
	fmt.Println(n) // 4
	fmt.Println(err) // <nil>
	fmt.Println(bs) //[105 106 13 10]
	fmt.Println(string(bs)) // ij
	// 四次读取
	fmt.Println("第四次读取....")
	n,err = file.Read(bs)
	fmt.Println(n) // 2
	fmt.Println(err) // EOF: End Of File
	fmt.Println(bs) //[105 106 13 10]
	fmt.Println(string(bs)) // ij
	*/
defer file.Close() // 关闭文件


// 8. 文件的写入,若文件不存在,创建
//file,err := os.Open(`E:\Egolang\assasin\assasin01\write.txt`)
file,err := os.OpenFile(`E:\golang\assasin\assasin01\write.txt`,os.O_WRONLY|os.O_CREATE,os.ModePerm)
if err != nil {
    fmt.Println("文件不存在,",err.Error())
}


bs := []byte{65,66,67,68,69,70}
n,err := file.Write(bs)
fmt.Println(n) // 6
fmt.Println(err) // <nil>

n,err = file.WriteString("hello面朝大海")
fmt.Println(err) // <nil>
fmt.Println(n) // 17

file.Close()

file,_ = os.OpenFile(`E:\golang\assasin\assasin01\write.txt`,os.O_APPEND,os.ModePerm)
file.WriteString("再次写入......")
file.Close()


// 9. 文件复制

// copyfile,返回值: 拷贝数量,错误
func copyFile(srcfile,destfile string) (int,error){
	src,err := os.Open(srcfile)
	if err != nil {
		return 0,err
	}
	dest,err := os.OpenFile(destfile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil {
		return 0,err
	}
	defer src.Close()
	defer dest.Close()

	bs := make([]byte,1024,1024)
	n := -1 //读取的数据量
	total := 0
	for {
		n,err = src.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("复制结束...")
			break
		} else if err != nil {
			fmt.Println("复制出错...")
			return total,err
		}
		total += n
		dest.Write(bs[:n])
	}

	return total,nil
}

srcfile := `E:\golang\assasin\assasin01\img.jpg`
destfile := `E:\golang\assasin\img_copy.jpg`
total,err := copyFile(srcfile,destfile)
fmt.Println(total) // 226871
fmt.Println(err) // nil

// 优化方法 io.Copy
func copyFile1(srcfile,destfile string) (int64,error) {
	src,err := os.Open(srcfile)
	if err != nil {
		return 0,err
	}
	dest,err := os.OpenFile(destfile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil {
		return 0,err
	}
	defer src.Close()
	defer dest.Close()

	return  io.Copy(dest,src)
}

srcfile := `E:\golang\assasin\assasin01\img.jpg`
destfile := `E:\golang\assasin\img_copy123.jpg`
total,err := copyFile1(srcfile,destfile)
fmt.Println(total) // 226871
fmt.Println(err) // nil


// 10. seek设置光标位置
// seek(偏移量,从哪开始)
// 0: 开头, 1: 距离光标当前位置,2:距离文件末尾
file,_ := os.OpenFile(`E:\golang\assasin\assasin01\123.txt`,os.O_RDWR,0)
bs := []byte{0}

file.Read(bs)
fmt.Println(string(bs)) // a

file.Seek(4,0)
file.Read(bs)
fmt.Println(string(bs)) // e

file.Seek(2,0)
file.Read(bs)
fmt.Println(string(bs)) // c

file.Seek(3,1)
file.Read(bs)
fmt.Println(string(bs)) // g

file.Seek(0,2)
file.WriteString("ABC")


// 11. ioutil包读文件
fileName1 := `E:\golang\assasin\assasin01\123.txt`
// 读取文件中的所有数据
data,err := ioutil.ReadFile(fileName1)
fmt.Println(err)
fmt.Println(string(data))


// 12 写文件
fileName2 := `E:\golang\assasin\assasin01\bbb.txt`
s1 := "helloworld我的哥"
err = ioutil.WriteFile(fileName2,[]byte(s1),os.ModePerm)
fmt.Println(err)


// 11. ioutil包读文件
fileName1 := `E:\golang\assasin\assasin01\123.txt`
// 读取文件中的所有数据
data,err := ioutil.ReadFile(fileName1)
fmt.Println(err)
fmt.Println(string(data))

// 12 写文件
fileName2 := `E:\golang\assasin\assasin01\bbb.txt`
s1 := "helloworld我的哥"
err = ioutil.WriteFile(fileName2,[]byte(s1),os.ModePerm)
fmt.Println(err)

// 13. 读取文档所有数据,返回数组
s2 := `sdfghjkxcvbnmertyuio`
r1 := strings.NewReader(s2)
data,err = ioutil.ReadAll(r1)
fmt.Println(data)

// 14. 读取目录,返回切片,只读一层
dirName := `E:\golang`
fileInfos,err := ioutil.ReadDir(dirName)
fmt.Println(len(fileInfos))
for i := 0;i < len(fileInfos);i++ {
    fmt.Printf("%T\n",fileInfos)
    fmt.Println(i,fileInfos[i].Name(),fileInfos[i].IsDir())
}


// 15. bufio包,高效读写,带缓存
filename := `E:\golang\assasin\assasin01\123.txt`
file,_ := os.Open(filename) // 看作是io包的Reader,Writer对象
r1 := bufio.NewReader(file) // 构建带缓存的Reader对象:bufio.Reader
fmt.Printf("%T\n",r1) // *bufio.Reader
//data,flag,err := r1.ReadLine()
//fmt.Println(err)
//fmt.Println(flag)
//fmt.Println(string(data))

//s1,err := r1.ReadString('\n')
//fmt.Println(s1)
//fmt.Println(err)
for {
    s1,err := r1.ReadString('\n')
    if err == io.EOF {
        fmt.Println("读取完毕...")
        break
    }
    fmt.Println(s1)
}
defer  file.Close()

data,_ := r1.ReadBytes('\n')
fmt.Println(string(data))

// 获取键盘输入 带空格
r2 := bufio.NewReader(os.Stdin)
s2,_ := r2.ReadString('\n')
fmt.Println(s2)
```

```go
// case1: 遍历文件夹
func travelDir(dirname string) {
	fileInfos,_ := ioutil.ReadDir(dirname)
	for _,fi := range fileInfos {
		fileName := dirname + `\` + fi.Name()
		fmt.Println(fileName)
		if fi.IsDir() {
			travelDir(fileName)
		}
	}
}
// case2: 模拟聊天记录,写入本地文件
fileName := `E:\golang\chat.txt`
file,_ := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
//fmt.Println(filepath.Abs(fileName)) // E:\golang\chat.txt, nil
r1 := bufio.NewReader(os.Stdin) // 创建bufio.NewReader对象读取键盘输入
w1 := bufio.NewWriter(file) // 创建bufio.NewWriter对象跌数据到chat.txt

defer file.Close()
str := "" //表示读取到的数据
w1.WriteString(time.Now().Format("2006-01-02"))
w1.WriteString(time.Now().Format("\n"))
w1.Flush() // 刷新缓存区:将缓存区的数据,写入目标文件中

name := ""
content := ""
flag := true
for {
    // 1. 设置name
    if flag {
        name = "小明"
    }else{
        name = "小红"
    }
    flag = !flag
    // 2. 读取键盘输入
    str,_ = r1.ReadString('\n') // 读取的内容包含 \n
    if str == "over\n" {
        fmt.Println("程序即将退出...")
        w1.WriteString(str)
        w1.Flush()
        break
    }
    // 3. 拼接字符串
    //content = name + ":" + str
    content = fmt.Sprint(name,":",str)
    fmt.Print(content)
    // 4. 将数据写入文件中
    w1.WriteString(time.Now().Format("15:04:05"))
    w1.WriteString("\n")
    w1.WriteString(content)
    w1.Flush()
}

// case3: 断点续传,以文件复制为例

srcfile := `E:\golang\assasin\123.jpg`
destfile := `E:\golang\assasin\assasin01\123_copy.jpg`
tmpfile := `E:\golang\assasin\assasin01\123_copy.jpg.txt`

file1,_ := os.Open(srcfile)
defer  file1.Close()
file2,_ := os.OpenFile(destfile,os.O_CREATE|os.O_WRONLY,os.ModePerm)
defer  file2.Close()
file3,_ := os.OpenFile(tmpfile,os.O_CREATE|os.O_RDWR,os.ModePerm)

// 1. 读物临时文件的数据
file3.Seek(0,0)
bs := make([]byte,100,100)
n1,err:= file3.Read(bs)
countStr := string(bs[:n1]) // 去掉 \n
fmt.Println(countStr)
//count,_ := strconv.Atoi(countStr)
count,_ := strconv.ParseInt(countStr,10,64)
fmt.Println(count)
// 2. 设置读写的偏移量
file1.Seek(count,0)
file2.Seek(count,0)

data := make([]byte,1024,1024)
n2 := -1  //读入的数据量
n3 := -1  // 写出的数据量
total := int(count) // 读取的数据量
for {
    // 读取数据
    n2,err = file1.Read(data)
    if err == io.EOF {
        fmt.Println("文件复制完毕")
        file3.Close()
        os.Remove(tmpfile)
        break
    }
    //将数据写入目标文件
    n3,_ = file2.Write(data[:n2])
    total += n3
    // 将复制总量写入临时文件
    file3.Seek(0,0)
    file3.WriteString(strconv.Itoa(total))


    // 模拟断点
    //if total > 10000 {
    //	panic("断电了....")
    //}
}

```

### 21. Process & Thread & Coroutine & Goroutines

```go
// 进程Process: 一个独立正在执行的程序;进程是一个程序在一个数据集中的一次动态执行过程,可以简单地理解为"正在执行的程序",它是CPU资源分配和调度的独立单位.进程一般由程序,数据及,进程控制块三部分构成.我们编写的程序用来描述进程要完成哪些功能以及如何完成;数据集则是程序在执行过程中所需要使用的资源;进程控制块用来记录进程的外部特征,描述进程的执行变化过程,系统可以利用它来控制和管理进程,它是系统感知进程存在的唯一标志.进程的缺陷是创建,撤销和切换的开销比较大.

// 线程Thread: 一个进程中的一条执行路径(单元),线程又叫轻量级进程,它是一个基本的CPU执行单元,也是程序执行过程中的最小单元.由线程ID,程序计数器,寄存器集合和堆栈共同组成.一个进程可以包好多个线程.线程的优点是减少了程序并发执行时的开销,提高了操作系统的并发性能,缺点是线程没有自己的系统资源,只拥有在运行时必不可少的资源,但同一进程的各线程可以共享进程所拥有的的系统资源,如果把进程比作一个车间,那么线程就好比是车间里的工人.不过对于某些独占性资源存在锁机制,处理不当可能会出现"死锁".

// 协程Coroutine: 轻量级的用户线程,微线程,协程的调用完全由用户控制.子程序调用总是一个入口,一次返回,一旦退出即完成了子程序的执行.协程的最大优势在于其"轻量级",可以轻松创建上百万个而不会导致系统资源衰竭,而线程和进程通常最多也不能超过1万的,这也是协程叫轻量级线程的原因.

// Goroutines: Go使用Goroutines来实现并发,Goroutines是与其他函数或方法同时运行的函数或方法.Goroutines可以被认为是轻量级的线程.与线程相比,创建Goroutine的成本很小,它就是一段代码,一个函数入口.以及在堆上为其分配的一个堆栈(初始大小为4K,会随着程序的执行自动增长删除).因此它非常廉价,go应用程序可以并发运行数千个Goroutines.
// Goroutine在线程上的优势:
// 1. 与线程相比, 它们只是堆栈大小的几个KB,堆栈可以根据应用程序的需要增长和收缩,而在线程的情况下,堆栈大小必须指定并且是固定的.
// 2. Goroutines被多路复用到较少的OS线程.在一个程序中可能只有一个线程与数千个Goroutines.如果线程中的任何Goroutines都表示等待用户输入,则会创建另一个OS线程,剩下的Goroutines被转移到新的OS线程.所有这些都由运行时进行处理,我们从这些复杂的细节中抽象出来,并得到了一个与并发工作相关的干净的API.
// 3. 当时用Goroutines访问共享内存时,通过设计的通道可以防止竞态条件发生.通道可以被认为是Goroutines通信的管道.

// 1. 启动主Goroutine,执行main函数,main函数结束,主Goroutine随之结束,程序结束.封装main函数的Goroutine成为主Goroutine.主Goroutine首先要做的是: 设定每一个Goroutine所能申请的栈空间的最大尺寸.在32位的计算机系统中此最大尺寸是250MB,在64位的计算机系统中此尺寸是1GB.如果有某个Goroutine的栈空间尺寸大于这个限制,那么运行时系统就会引发一个栈溢出(stack overflow)的运行时恐慌.随即这个go程序的运行也会终止.
// 此后,主Goroutine会进行一系列的舒适化工作,涉及的工作内容大致如下:
// 1. 创建一个特殊的defer语句,用于在主Goroutine退出时作必要的善后处理.因为主Goroutine也可能非正常退出;
// 2. 启动专用于在后台清扫内存垃圾的Goroutine,并设置GC可用的标识;
// 3. 执行main包中的init函数;
// 4. 执行main函数,执行完main函数后,它还会检查主Goroutine是否引发了运行时恐慌,并进行必要的处理.最后主Goroutine会结束自己以及当前进程的运行.


// 2. 子Goroutine,用户启动的Goroutine

func main() {
	go hello() //启动子 Goroutine
	fmt.Println("main......")
}

func hello() {
	fmt.Println("这是一个函数,在另外一个Goroutine中执行")
}


//输出:
这是一个函数,在另外一个Goroutine中执行
main......
//或
main......
这是一个函数,在另外一个Goroutine中执行
//或
main......


```

```go
// case 1
func func1 (){
	for i := 0;i <= 1000;i++ {
		fmt.Println("子Goroutine中i的值:",i)
		time.Sleep(1*100)

	}
}

func func2() {
	for i := 1;i <= 1000;i++ {
		fmt.Printf("\t子Goroutine中字母:%d,%c\n",i,i)
		time.Sleep(1*10)
	}
}

go func1()
go func2()
fmt.Println("main......")
time.Sleep(1*time.Second)

// case 2
func func1(){
	for i := 1;i <= 5;i++{
		time.Sleep(250 * time.Millisecond)
		fmt.Print(i,"\t")
	}
}

func func2(){
	for i := 65;i <= 70;i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\t",i)
	}
}

go func1()
go func2()
time.Sleep(3000*time.Millisecond)

// case 4 

// 获取当前计算机逻辑CPU核数
	fmt.Println("逻辑CPU的核数:",runtime.NumCPU())
	// 设置当前go程序执行的最大逻辑CPU核数
	runtime.GOMAXPROCS(runtime.NumCPU()) //默认是 1
```

### 22. Channel 

```go
// go语言通过数据传递实现共享内存,而不是通过共享内存来实现消息传递.
// channel,多个Goroutine之间传递消息的介质.
// 操作符: <-
// 从channel中读数据: data,ok := <- channel,ok为true,通道正常,data有效可用;ok为false,通道关闭,data是类型的零值
// 向channel中写数据: channel <- data
// 本身是安全的,同一时间只能允许一个Goroutine来读写.

var ch1 chan int
fmt.Println(ch1) //<nil>
fmt.Printf("%T\n",ch1) //chan int
// nil 的channel同map一样,不能使用
ch1 = make(chan int)
fmt.Println(ch1)
go func() {
    fmt.Println("子Goroutine启动......")
    data := <- ch1 // 阻塞式,从channel中读取数据
    fmt.Println("子Goroutine从channel读取的数据:",data)
}()
ch1 <- 100 // 阻塞式,主Goroutine向channel中写入数据
fmt.Println("main over ....")

// channel阻塞:
ch1 := make(chan string)
go func() {
    fmt.Println("子Goroutine...")
    time.Sleep(5 * time.Second)
    ch1 <- "hello"
}()

time.Sleep(2 * time.Second)
data := <- ch1
fmt.Println("读取到的数据是:",data)



func func1(ch chan bool){
	fmt.Println("func1...ch",ch)
	for i := 1; i <= 100;i++ {
		fmt.Println(i)
		time.Sleep(1000)
	}
	ch <- true
}

func func2(ch chan bool){
	for i := 1; i <= 100;i++ {
		fmt.Printf("\t%d,%c\n",i,i)
		time.Sleep(1)
	}
	ch <- true
}

ch1 := make(chan bool)
fmt.Println("main... ->ch1")
go func1(ch1)
go func2(ch1)

<- ch1
<- ch1


// channel关闭:发送方如果数据写入完毕,关闭通道,通知接收方数据传递完毕.
// close(chan)

func sendData(ch chan int) {
	for i := 1; i <= 10;i++ {
		ch <- i
	}
	fmt.Println("发送方写入数据完毕...")
	// 关闭通道
	close(ch)
}

ch1 := make(chan int)
go sendData(ch1)
for {
    time.Sleep(1 * time.Second)
    data,ok := <- ch1
    if !ok {
        fmt.Println("读取数据完毕....",ok)
        break
    }
    fmt.Println("main中读取到数据:",data,ok)
}

//输出:
/*  main中读取到数据: 1 true
    main中读取到数据: 2 true
    main中读取到数据: 3 true
    main中读取到数据: 4 true
    main中读取到数据: 5 true
    main中读取到数据: 6 true
    main中读取到数据: 7 true
    main中读取到数据: 8 true
    main中读取到数据: 9 true
    main中读取到数据: 10 true
    发送方写入数据完毕...
    读取数据完毕.... false
 */



// for...range 读取通道数据,返回通道数据:
//  停止条件: 通道关闭,显示的调用close函数
func sendData(ch chan string) {
	for i := 1;i <= 10;i++ {
		ch <- fmt.Sprint("数据:",i)
	}
	fmt.Println("写入数据完毕")
	close(ch)
}

ch1 := make(chan string)
go sendData(ch1)
for value := range ch1 {
    // 停止条件: 通道关闭,显式的调用close函数
    fmt.Println("从channel中读取到的数据:",value)
}

//输出:
/*
从channel中读取到的数据: 数据:1
从channel中读取到的数据: 数据:2
从channel中读取到的数据: 数据:3
从channel中读取到的数据: 数据:4
从channel中读取到的数据: 数据:5
从channel中读取到的数据: 数据:6
从channel中读取到的数据: 数据:7
从channel中读取到的数据: 数据:8
从channel中读取到的数据: 数据:9
从channel中读取到的数据: 数据:10
写入数据完毕
*/

// 缓存通道(自带一块缓存区,可以暂存数据,若缓存区满了,才会阻塞,默认创建的都是非缓存通道):

// 1. 非缓存通道
ch1 := make(chan int)
go func() {
    time.Sleep(3 * time.Second)
    <- ch1 // 即时阻塞
}()
ch1 <- 100 // 阻塞
fmt.Println("写入数据完毕...")

// 2. 缓存通道
// ch2 := make(chan string,5)

func sendData(ch chan string) {
	for i := 1;i <= 100;i++ {
		ch <- fmt.Sprint("数据:",i)
		fmt.Println("-->写入数据:",i)
	}
	close(ch)
}

ch2 := make(chan string,5)
//fmt.Printf("%T\n",ch2)
go sendData(ch2)

for {
    time.Sleep(100 * time.Millisecond)
    data,ok := <- ch2
    fmt.Println("\t<--读取数据: ",data,",ok:",ok)
    if !ok {
        fmt.Println("<--读取完毕...")
        break
    }
}


// case: 一个goroutine写数据,另外两个goroutine读取数据
ch1 := make(chan int,5)
ch2 := make(chan bool) // 判断结束
// 写入数据: 生产者
go func() {
    rand.Seed(time.Now().UnixNano())
    for i := 1;i <= 100;i++ {
        ch1 <- i
        fmt.Println("-->写入数据:",i)
        //time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    }
    close(ch1)
}()
// 读取数据: 消费者
go func() {
    for data := range ch1 {
        fmt.Println("\t<--消费者1:",data)
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    }
    ch2 <- true
}()
go func() {
    for data := range ch1 {
        fmt.Println("\t<--消费者2:",data)
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    }
    ch2 <- true
}()

<-ch2
<-ch2
fmt.Println("main over...")


// 定向通道(单向通道)只读的或只写的通道:
// 只读通道: make(<- chan Type),只能读取数据,不能写入
// 只写通道: make(chan <- Type),只能写入数据,不能读取
// 创建通道时,采用单向通道,是没有意义的.
// 使用场景: 传递参数的时候使用,限定某个函数只有写数据或读取数据,在语法级别保证通道的操作安全.


func main() {
	// 双向通道
	ch1 := make(chan string)
	go func1(ch1)
	time.Sleep(3 * time.Second)
	data := <- ch1
	fmt.Println("接收到数据:",data)
	ch1 <- "你要上学吗?"


	func2(ch1)
	func3(ch1)

}
func func1(ch chan string) {
	ch <- "我是小明"
	time.Sleep(2 * time.Second)
	data := <- ch
	fmt.Println("回应:",data)
}

// 只能写入数据,不能读取
func func2(ch chan <- string) {
	//TODO
}
// 只能读取数据,不能写入
func func3(ch <- chan string) {
	//TODO

}

// 注意点:
for i := 1;i <= 3;i++ {
    go func(i int) {
        fmt.Println("第",i,"个goroutine")
    }(i)
    //time.Sleep(10 * time.Millisecond)
}
//或
for i := 1;i <= 3;i++ {
    i := i
    go func() {
        fmt.Println("第",i,"个goroutine")
    }()
}
time.Sleep(1 * time.Second)


```

### 23. Timer计时器

```go
//  Timer 计时器
timer1 := time.NewTimer(3 * time.Second)
fmt.Printf("%T\n",timer1) // *time.Timer
fmt.Println(time.Now()) // 2021-02-14 23:41:03.0132543 +0800 CST m=+0.002032701
time1 := <- timer1.C // 读取数据
fmt.Println(time1) // 2021-02-14 23:41:06.0160538 +0800 CST m=+3.004832201

fmt.Println("-----------------")
// 使用after,返回值 <- chan Time
ch1 := time.After(5 * time.Second)
fmt.Println(time.Now()) // 2021-02-14 23:45:13.3266743 +0800 CST m=+3.007278401
time2 := <- ch1
fmt.Println(time2) // 2021-02-14 23:45:18.3318006 +0800 CST m=+8.012404701
```

### 24.  Select

```go
// 类似于switch语句,但select会随机执行一个可运行的case.如果没有case可以运行,若有default语句,则执行default中,否则将会阻塞,直到有case可以运行.
// 每个case后的操作都会运算,若多个都可执行,select会随机执行一个可运行的case.

var ch1,ch2,ch3 chan int
var n1,n2 int
select {
    case n1 = <- ch1:
    fmt.Println("received:",n1,"from ch1\n")
    case ch2 <- n2:
    fmt.Println("send",n2,"to ch2 \n")
    case n3,ok := (<-ch3): //same as n3,ok := <- ch3
    if ok {
        fmt.Println("received",n3,"from ch3\n")
    } else {
        fmt.Println("ch3 is closed...\n")
    }
    default:
    fmt.Println("no communication\n")
}
 
// 输出: no communication


ch1 := make(chan int)
ch2 := make(chan int)
go func() {
    time.Sleep(3 * time.Second)
    ch1 <- 100
}()
go func() {
    time.Sleep(3 * time.Second)
    ch2 <- 200
}()
select {
    case data,ok := <- ch1:
    fmt.Println("ch1中读取数据:",data,ok)
    case data := <- ch2:
    fmt.Println("ch2中读取数据:",data)
    default:
    fmt.Println("执行了default...")
}


// 注意循环的结束, 贴标签

ch1 := make(chan int)
ch2 := make(chan int)

go func() {
    ch1 <- 100
    close(ch1)
}()
go func() {
    ch2 <- 200
}()

// 通道关闭或超时5次结束循环
count := 0
out : for {
    time.Sleep(1 * time.Second)
    select {
        case data,ok := <- ch1 :
        if !ok {
            fmt.Println("ch1关闭了...")
            break out
        }
        fmt.Println("ch1读取数据:",data)
        case data := <- ch2:
        fmt.Println("ch2读取数据:",data)
        case <- time.After( 2 * time.Second):
        fmt.Println("超时 2 秒钟")
        count++
        if count == 5 {
            break out
        }
        //default:
        //	fmt.Println("default....")
    }
}

```

### 25. WaitGroup 同步等待组

```go
// 同步等待组: 
// 内置计数器:要执行的goroutine的数量
// Add(),设置计数器的数量
// Done(),将内置计数器的数值减1,同Add(-1)
// Wait(),等待,导致执行wait的goroutine进入阻塞.同步等待组的计数器为0时,解除阻塞.


func func1(wg *sync.WaitGroup){
	rand.Seed(time.Now().UnixNano())
	for i := 1;i <= 100;i++ {
		fmt.Println("子goroutine1,i:",i)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}
	wg.Done()

}
func func2(wg *sync.WaitGroup){
	for j := 1;j <= 100;j++ {
		fmt.Println("子goroutine2,j:",j)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}
	wg.Done()
}

/*
同步等待组:WaitGroup,执行Wait()的goroutine,要等待同步等待组中的其他goroutine执行完毕.
内置的计数器 counter: 0
Add():设置counter的值
Done():将counter的值 -1 ,同Add(-1)
Wait():哪个goroutine执行了,name就会阻塞,知道counter = 0 解除阻塞
*/
var wg sync.WaitGroup
//fmt.Printf("%T\n",wg) //sync.WaitGroup
//fmt.Println(wg) // {{} [0 0 0]}
wg.Add(2)

go func1(&wg)
go func2(&wg)
wg.Wait() // main,进入阻塞状态,底层计数器为0,解除阻塞
fmt.Println("main解除阻塞,结束...")



// 共享数据的安全性问题:

// case : 模拟车站卖票,车票100张,4个售票口(4个goroutine)
var tickets = 100 // 全局变量
var wg sync.WaitGroup
func main() {
	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	wg.Wait()
	fmt.Println("本次列车所有车票已售完,程序结束...")

}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)))
			fmt.Println(name,":",tickets)
			tickets --
		} else {
			fmt.Println(name,"停止售票...")
			break
		}
	}
	wg.Done()
}
/*
......
售票口2 : 3
售票口3 : 2
售票口3 停止售票...
售票口1 : 0
售票口1 停止售票...
售票口2 : -1
售票口4 : -1
售票口4 停止售票...
售票口2 停止售票...
本次列车所有车票已售完,程序结束...
*/

// 优化方法: 互斥锁

/*
互斥锁:
Lock(),上锁事件是阻塞的,只能有一个goroutine上锁,其他goroutine处于阻塞状态
Unlock(),
*/
var mutex sync.Mutex
fmt.Println("main即将锁定...")
mutex.Lock()
fmt.Println("main已经锁定...")

for i := 1;i <= 3;i++ {
    go func(i int) {
        fmt.Println("子goroutine,i:",i,"即将锁定....")
        mutex.Lock() //导致阻塞
        fmt.Println("子goroutine,i:",i,"已经锁定")
    }(i)
}

time.Sleep(5 * time.Second)
fmt.Println("main即将解锁...")
mutex.Unlock()
fmt.Println("main已经解锁...")

time.Sleep(3 * time.Second)

/*
输出:
main即将锁定...
main已经锁定...
子goroutine,i: 1 即将锁定....
子goroutine,i: 3 即将锁定....
子goroutine,i: 2 即将锁定....
main即将解锁...
main已经解锁...
子goroutine,i: 3 已经锁定
*/


// 卖票案例优化:
var tickets = 100 // 全局变量
var wg sync.WaitGroup
var mutex sync.Mutex // 互斥锁
func main() {
	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	wg.Wait()
	fmt.Println("本次列车所有车票已售完,程序结束...")

}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		mutex.Lock()
		if tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)))
			fmt.Println(name,":",tickets)
			tickets --
		} else {
			fmt.Println(name,"停止售票...")
			mutex.Unlock() // 结束之前开锁
			break
		}
		mutex.Unlock()
	}
	wg.Done()
}

/*
输出:
......
售票口4 : 5
售票口2 : 4
售票口3 : 3
售票口1 : 2
售票口4 : 1
售票口2 停止售票...
售票口3 停止售票...
售票口1 停止售票...
售票口4 停止售票...
本次列车所有车票已售完,程序结束...
*/

```

### 26. 读写锁

```go
/*
读写锁: sync.RWMutex()
Lock(),写锁定
Unlock(),写解锁
Rlock(), 读锁定
RUnlock(),读解锁
锁定规则:
	1. 读写锁使用中,所有写错做都是互斥的
	2. 读写锁使用中,读和写是互斥的
	2. 读写锁使用中,读和读不互斥
	可以允许多个goroutine同时读数据,但是写操作只允许一个goroutine执行
*/


func main() {
	var rwm sync.RWMutex

	for i := 1;i <= 3;i++ {
		go func(i int) {
			fmt.Printf("goroutine %d 尝试读锁定...\n",i)
			rwm.RLock()
			fmt.Printf("goroutine %d 已经读锁定...\n",i)
			time.Sleep(5 * time.Second)
			fmt.Printf("goroutine %d 已经读解锁...\n",i)
			rwm.RUnlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main...尝试写锁定")
	rwm.Lock()
	fmt.Println("main已经写锁定...")
	rwm.Unlock()
	fmt.Println("main写解锁...")
	
}
/*
输出:
goroutine 2 尝试读锁定...
goroutine 2 已经读锁定...
goroutine 1 尝试读锁定...
goroutine 1 已经读锁定...
goroutine 3 尝试读锁定...
goroutine 3 已经读锁定...
main...尝试写锁定
goroutine 1 已经读解锁...
goroutine 3 已经读解锁...
goroutine 2 已经读解锁...
main已经写锁定...
main写解锁...
*/


var n int
var wg sync.WaitGroup
var rwm sync.RWMutex
func main() {
	wg.Add(10)
	for i := 1;i <= 5;i++ {
		go Read(i)
	}

	for i := 1;i <= 5;i++ {
		go Write(i)
	}

	wg.Wait()

}
func Write(i int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	rwm.Lock()
	fmt.Println("写操作:", i, "即将写入数据...")
	randNum := rand.Intn(100) + 1
	n = randNum
	fmt.Println("写操作:", i, "已经写入数据:", n)
	rwm.Unlock()
}

func Read(i int) {
	defer wg.Done()
	rwm.RLock()
	fmt.Println("读操作:",i,"即将读取数据...")
	v := n
	fmt.Println("读操作:",i,"读取了数据:",v)
	rwm.RUnlock()
}



```

### 27. 并发操作map

```go
var rwm sync.RWMutex
map1 := make(map[string] string)
for i := 1;i <= 100;i ++{
    go func(i int) {
        rwm.Lock()
        map1[fmt.Sprintf("key,%d",i)] = fmt.Sprint("数据:",i)
        rwm.Unlock()
    }(i)
}
time.Sleep(2 * time.Second)
fmt.Println(map1)

// 输出: map[key,1:数据:1 key,10:数据:10 key,2:数据:2 key,3:数据:3 key,4:数据:4 key,5:数据:5 key,6:数据:6 key,7:数据:7 key,8:数据:8 key,9:数据:9]

// 封装map支持并发操作:
type MyMap struct {
	map1 map[string] string
	rwm sync.RWMutex
}

func (m *MyMap) Put(key,value string) {
	m.rwm.Lock()
	m.map1[key] = value
	m.rwm.Unlock()
}

func (m *MyMap) Get(key string) string {
	m.rwm.RLock()
	defer m.rwm.RUnlock()
	return m.map1[key]
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	m1 := make(map[string] string)
	var rwm sync.RWMutex
	m2 := MyMap{m1,rwm}
	for i := 1;i <= 10;i ++ {
		go func(i int) {
			m2.Put(fmt.Sprint("key",i),fmt.Sprint("data",i))
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("main读取数据...")
	for i := 1;i <= 10;i ++ {
		fmt.Println(m2.Get(fmt.Sprint("key",i)))
	}
	fmt.Println("main over...")
}

/*
输出:
main读取数据...
data1
data2
data3
data4
data5
data6
data7
data8
data9
data10
main over...
*/

```

### 28. 条件变量Cond

```go
/*
	条件变量: sync.Cond,读个goroutine等待或接受通知的集合地
	L : Locker 接口
	3个指针方法:
		Wait(),goroutine等待接受通知,Single(),Broadcast()
		Single(),发送通知,一个
		Broadcast(),广播,发送给所有
*/
func main() {
	var mutex  sync.Mutex
	cond :=  sync.Cond{L:&mutex}
	condition := false

	go func() {
		time.Sleep(1 * time.Second)
		cond.L.Lock()
		fmt.Println("子goroutine已经锁定")
		fmt.Println("子goroutine更改条件数值,并发送通知...")
		condition = true
		cond.Signal() // 发送通知:一个goroutine,
		time.Sleep(5 * time.Second)
		fmt.Println("子goroutine继续...")
		fmt.Println("子goroutine已经解锁")
		cond.L.Unlock()
	}()

	cond.L.Lock()
	fmt.Println("main已经锁定...")
	if !condition {
		fmt.Println("main 即将等待")
		// wait尝试先解锁,等待 --->当前的goroutine进入阻塞状态,等待被唤醒,Single或Broadcast
		// 被唤醒后,又会去锁定
		cond.Wait()
		fmt.Println("main被唤醒...")
	}
	fmt.Println("main继续其他工作...")
	fmt.Println("main解锁...")
	cond.L.Unlock()
}

/*
输出:
main已经锁定...
main 即将等待
子goroutine已经锁定
子goroutine更改条件数值,并发送通知...
子goroutine继续...
子goroutine已经解锁
main被唤醒...
main继续其他工作...
main解锁...
*/


// case:

func main() {
	var count = 4
	var wg sync.WaitGroup
	wg.Add(5)

	// 新建cond
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for i := 0;i < 5; i++ {
		go func(i int) {
			// 争抢互斥锁的锁定
			cond.L.Lock()
			//条件是否达成
			for count > i {
				cond.Wait()
				fmt.Printf("收到第一个通知 goroutine %d\n",i)
			}
			fmt.Printf("goroutine %d 执行结束\n",i)
			cond.L.Unlock()
			wg.Done()
		}(i)
	}
	// 确保所有 goroutine启动完毕
	time.Sleep(20 * time.Millisecond)
	// 锁定一下,改变 count的值
	fmt.Println("broadcast...")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	fmt.Println("第一次,广播结束...")
	cond.L.Unlock()

	time.Sleep(2 * time.Second)
	fmt.Println("------------------------")
	fmt.Println("signal...")
	cond.L.Lock()
	count -= 2
	cond.Signal()
	fmt.Println("第二次,单发通知结束...")
	cond.L.Unlock()

	time.Sleep(2 * time.Second)
	fmt.Println("------------------------")
	fmt.Println("broadcast")
	cond.L.Lock()
	count -= 1
	cond.Broadcast()
	fmt.Println("第三次,广播结束...")
	cond.L.Unlock()

	wg.Wait()

}

/*
输出:
goroutine 4 执行结束
broadcast...
第一次,广播结束...
收到第一个通知 goroutine 1
收到第一个通知 goroutine 0
收到第一个通知 goroutine 2
收到第一个通知 goroutine 3
goroutine 3 执行结束
------------------------
signal...
第二次,单发通知结束...
收到第一个通知 goroutine 1
goroutine 1 执行结束
------------------------
broadcast
第三次,广播结束...
收到第一个通知 goroutine 2
goroutine 2 执行结束
收到第一个通知 goroutine 0
goroutine 0 执行结束

*/


```

### 29. sync的其他操作

```go
// 1. 原子操作
/*
针对数值:通过原子性方法操作某个数值时,其他goroutine不能再访问当前的数值变量.
操作: 原子加/减
	  交换
	  比较并交换:CAS
	  存储
	  加载
原子性操作: sync.atomic
非原子性操作:
*/
var n int64 = 3
fmt.Println(`n的原始数值是:`,n)
//原子加 返回新值
newN := atomic.AddInt64(&n,1)
fmt.Println("新的数据newN:",newN)
fmt.Println("n的新值:",n)

atomic.SwapInt64(&n,9) //更新数据
atomic.CompareAndSwapInt64(&n,4,10) //比较并交换

/*
输出:
n的原始数值是: 3
新的数据newN: 4
n的新值: 4
*/

// 2. 一次操作
/*
一次操作: 某些操作只需要一次
Do(func()) -> 设置多次,也只会执行一次
*/
count := 0
once := sync.Once{}
for i := 1;i <= 10;i++ {
    once.Do(func() {
        count++
    })
}
fmt.Println(count)


// 3. 临时对象池
/*
临时对象池:能够复用某些数据的容器,并可伸缩,同步安全
sync.Pool
	Put(),向pool存储数据
	Get(),获取一个数据,若pool中没有,则调用New对应函数,创建一个
若程序垃圾回收机制运行,则pool中数据都会没有了
*/
var count int64
fun := func() interface{} {
    return atomic.AddInt64(&count,1)
}
pool := sync.Pool{New:fun}
// 获取数据
fmt.Println(pool.Get())
// 存储数据
pool.Put(10)
pool.Put(2)
pool.Put(8)
pool.Put(15)
fmt.Println(pool.Get()) // 有多个数据,任意获取一个,若pool中没有,则调用New对应函数,创建一个

// 执行GC()
runtime.GC()
pool.New = nil
fmt.Println(pool.Get())
```

