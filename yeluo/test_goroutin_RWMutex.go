// 创建人：lg
// 创建时间：2021/5/19
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			m[rand.Int()] = rand.Int()
		}()
	}
	wg.Wait()
	fmt.Println(len(m))
}

// 考点:
// map并发读写的问题 并发访问map并不安全,会出现未定义行为,会导致程序退出,如果希望在多协程里并发访问map,
// 必须提供某种同步机制,通常可以使用读写锁 sync.RWMutex实现 对map的并发访问控制!也可以使用go提供的线程安全的map，即sync.Map
// fatal error: concurrent map writes
