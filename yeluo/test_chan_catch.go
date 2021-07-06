// 创建人：lg
// 创建时间：2021/5/17
package main

import (
	"fmt"
	"sync"
	"time"
)

type threadSafeSet struct {
	sync.RWMutex  // 这里不用读写锁也可以，因为不涉及到并发修改的问题
	s []interface{}
}

func main() {
	th := threadSafeSet{
		s: []interface{}{"1", "2", 3, 4, 5, true},
	}
	ch := th.Iter()
	// fmt.Println(fmt.Sprintf("%s: %v", "ch", <-ch)) // 这里程序会打印一个直接退出
	for v := range ch {  // 可以用循环获取所有的值，也可以使用带缓存的channel
		time.Sleep(time.Second)
		fmt.Println(fmt.Sprintf("%s: %v", "ch", v))
	}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	// ch := make(chan interface{})  // 不带缓存的是等待管道中的内容被读取后才能写入下一条
	ch := make(chan interface{}, len(set.s))  // 可以一次性写入所有数据
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- value
			fmt.Println("Iter: ", elem, value)
		}
		close(ch)
		set.RUnlock()
	}()
	time.Sleep(1*time.Second)
	return ch
}
