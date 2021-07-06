// 创建人：lg
// 创建时间：2021/5/17
package main

import (
	"sync"
	"time"
)

var (
	i  int
	mu sync.Mutex
	ch = make(chan byte, 1)
)

func main() {
	for range [1000]byte{} {
		go add()
	}
	time.Sleep(2 * time.Second)
	print(i)
}

func add() {
	// 使用互斥锁防止读个协程goroutine同时修改共享变量
	mu.Lock()
	defer mu.Unlock()
	i++
	// 使用通道进行同步
	ch <- 0
	i++
	<-ch
}
