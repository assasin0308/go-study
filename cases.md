### 1. 函数式编程

```go
// 1. 斐波那契数列:
// 1 ,2,,3,5,8,13,21......
func fibobacci() func() int {
	a,b := 0,1
	return func() int {
		a,b = b,a + b
		return a
	}
}
func main() {
	f := fibobacci()
	fmt.Println(f())
	// ......
}

// 2. 优化: 为函数实现接口
func fibonacci() intGen {
	a, b := 0,1
	return func() int {
		a, b = b, a + b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int,err error){
	next := g()
	if next > 1000 {
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	//TODO: incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	printFileContents(f)

}

```

### 2. go文档

```go
// 使用go doc/ godoc 生成/查看文档
go doc 
godock -http :6060
// 示例代码:
在测试中Example关键字
```

### 3. 表格驱动测试

```go
// 命令行测试
go test .

```

### 4. 

```go

```

### 5. 

```go

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