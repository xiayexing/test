// 创建人：lg
// 创建时间：2021/5/19
package main

import "sync"

const N = 10

var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < N; i++ {
		go func(i int) {
			wg.Add(i)
			print(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()

	// 可以正常执行
	// for i := 0; i < N; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		println(i)
	// 		defer wg.Done()
	// 	}(i)
	// }
	// wg.Wait()

}

// 1024657839fatal error: all goroutines are asleep - deadlock!
// 输出都会不同甚至又出现报错的问题。 这是因为go执行太快了，导致wg.Add(1)还没有执行main函数就执行完毕了。
