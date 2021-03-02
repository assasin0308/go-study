### Golang

### 1. beego开发环境

```go
/*
安装go
vim /etc/profile
export PATH=$PATH:/usr/local/go/bin
source /etc/profile

go mod init gin
go mod edit -require github.com/gin-gonic/gin@latest
```

### 2.  RPC

```json
// RPC: remote procedure call protocol,远程过程调用协议,这种调用的过程跨越了物理服务器的限制,在网络中完成,再调用过程中,本地程序等待返回结果,知道远程程序执行结束,将结果返回到本地,最终完成一次完整的调用.远程过程调用指的是调用远程服务器上的程序的方法的整个过程.
// 技术架构:
// 客户端: 服务调用发起方,又称为服务消费者
// 服务端: 远端计算机上运行的程序,其中包含客户端要调用和访问的方法
// 客户缓存根: 存放服务器的地址,端口消息,将客户的请求参数打包成网络消息,发送到服务方,接受服务方返回的数据包.该段程序运行在客户端.
// 服务端存根: 接受客户端发送的数据包,解析数据包,调用具体的服务方法.将调用结果打包发送给客户端一方.该段程序运行在服务端.
Golang-100-days
```

### 3.  RPC 使用

```go
// server

package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

// 数学计算
type MathUtil struct {
}
// 该方法向外暴露:提供计算圆形面积的的服务
func (mu *MathUtil) CalculateCircleArea(req float32,resp *float32) error {
	*resp = math.Pi * req * req
	return nil
}

func main() {
	// 1. 初始化指针数据类型
	mathUtil := new(MathUtil) // 初始化指针数据类型
	// 2. 调用net/rpc包的功能将服务对象进行注册
	err := rpc.Register(mathUtil)
    //err := rpc.RegisterName("MathUtil",mathUtil)  指定服务名称
	if err != nil {
		panic(err.Error())
	}
	// 3. 通过函数把MathUtil中提供的服务注册到http协议上,方便调用者可以利用http的方法进行数据传递
	rpc.HandleHTTP()

	// 4. 在特定的端口进行监听
	listen,err := net.Listen("tcp",":8081")
	if err != nil {
		panic(err.Error())
	}
    http.Serve(listen,nil)
    //go http.Serve(listen,nil)


}

```

```go
// client

package main

import (
	"fmt"
	"net/rpc"
)

// 客户端实践
func main() {
	client,err := rpc.DialHTTP("tcp","localhost:8081")
	if err != nil {
		panic(err.Error())
	}
	var req float32 // 请求值
	req = 3

	var resp *float32
    // 同步的调用方式
	err = client.Call("MathUtil.CalculateCircleArea",req,&resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*resp)

	//异步调用
	//var respSync *float32
	//syncCall := client.Go("MathUtil.CalculateCircleArea",req,&respSync,nil)
	//replayDone := <- syncCall.Done
	//fmt.Println(replayDone)
	//fmt.Println(*respSync)
	
}

```

### 4. RPC与Protobuf

```go
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/golang/protobuf/proto
// https://github.com/assasin0308/protobuf.git
// 执行 protoc --go_out=. *.proto 生成 test.pb.go 文件


// 需求: 假设在一个系统中,有订单模块(Order),其他模块想要实现RPC的远程工具调用,根据订单ID和时间戳可以获取订单信息.如果查询获取成功就返回相应的订单信息,如果查询不到返回失败信息.

// 1. 数据定义,根据需求,定义message.proto 文件

syntax = "proto3";

package message;
// 订单参数请求
message OrderRequest {
    string orderId = 1;
    int64 timeStamp = 2;
}
        
//订单信息
message OrderInfo {
    string OrderId = 1;
    string OrderName = 2;
    string OrderStatus = 3;
}

```

### 5. RPC框架

```go
// https://grpc.io/
// https://github.com/assasin0308/grpc.git

go get -u google.golang.org/grpc

// 执行 protoc --go_out=. *.proto 生成 test.pb.go 文件
// gRpc支持: protoc --go_out=plugins=grpc:. *.proto
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