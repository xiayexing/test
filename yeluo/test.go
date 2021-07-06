// 创建人：lg
// 创建时间：2020/7/23
package main

import "fmt"

func insert(p []int) {
	for i := 1; i < len(p); i++ {
		pre := i-1
		cu := p[i]
		for ; pre>0 &&  p[pre] > cu; pre-- {
			p[pre+1] = p[pre]
		}
		if pre != i -1 {
			p[pre+1] = cu
		}
	}
}

func main() {
	l := []int{1,3,5,2,4,7,6,9,8}
	insert(l)
	fmt.Println(l)
}