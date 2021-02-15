package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
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



}
