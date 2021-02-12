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

# 基本数据类型: bool,string,int,float
# 复合数据类型: array,slice,map,struct,pointer,interface,function......

# 值类型数据: bool,string,int,float,array,struct
# 引用类型数据: slice,map,function,pointer

# 浅拷贝: 拷贝的是数据的内存地址,多个变量指向同一内存地址,引用类型的数据默认都是浅拷贝: slice,map,function,
# 深拷贝: 拷贝的是数据本身(副本),对原始数据诶呦影响,值类型的数据都是深拷贝: array,int,float,string,bool,struct



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
第1个学生的信息:
学生姓名:王大狗,年龄:19
		第 1 本书名:孙子兵法,价格:120.80
		第 2 本书名:天局,价格:28.30
		第 3 本书名:易经经,价格:88.90
第2个学生的信息:
学生姓名:李小花,年龄:18
		第 1 本书名:红露梦,价格:77.80
		第 2 本书名:西游记,价格:98.20



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
```

### 16. error处理

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

### 24.  封装与包

```go
// 名字一般使用CamelCase 大驼峰
// 首字母大写: public
// 首字母小写: private

// 每个目录就是一个包,main包包含可执行入口
// 为结构定的方法必须放在同一个包内,可以是不同的文件
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

### 20. 依赖管理

```go
go mod init
go build ./...
go install ./...


go get -u go.uber.org/zap
```

