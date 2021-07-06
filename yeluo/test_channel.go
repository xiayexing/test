// 创建人：lg
// 创建时间：2021/5/19
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if ok {
				fmt.Println("channel closed")
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

// 考点:channel
// 往已经关闭的channel写入数据会panic的。
// 结果：
// panic: send on closed channel
