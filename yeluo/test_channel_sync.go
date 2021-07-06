// 创建人：lg
// 创建时间：2021/5/19
package main

import "fmt"

func main() {
	chanNum := make(chan bool)
	chanChar := make(chan bool, 1)
	done := make(chan struct{})

	go func() {
		for i := 1; i < 11; i += 2 {
			<-chanChar
			fmt.Println(i)
			fmt.Println(i + 1)
			chanNum <- true
		}
	}()

	go func() {
		charSeq := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
		for i := 0; i < 10; i += 2 {
			<-chanNum
			fmt.Println(charSeq[i])
			fmt.Println(charSeq[i+1])
			chanChar <- true
		}
		done <- struct{}{}
	}()
	chanChar <- true
	<-done
}

// 考点: chan的用法 两goroutine里的chan都没有值,一直在等,到倒数第二行,chan_c 赋值了,这时上面那个 goroutine开始执行,
// 然后给 chan_n 赋值,此时 chan_c 已经没有值了,又在等待,但是第二个 goroutine的chan_n有了值,然后第二个开始执行,
// 然后后面又给 chan_c 赋值.. 如此循环..打印所有
