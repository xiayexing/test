// 创建人：lg
// 创建时间：2020/7/15
package main

import "fmt"

func bubbleSort(params []int) []int {
	// 切片传参只有头部信息是值传递，切片所引用的底层数组不会复制。
	for i := 1; i < len(params); i++ {
		for j := 0; j < len(params)-i; j++ {
			if params[j] > params[j+1] {
				params[j], params[j+1] = params[j+1], params[j]
			}
		}
	}
	return params  // 是否有返回值都无所谓，因为排序是在原数组的内存空间进行的，所以返回的还是原数组的内存地址
}

func main() {
	p := []int{1, 2, 4, 6, 3, 7, 9, 5, 8}
	p1 := bubbleSort(p)
	fmt.Println(p)
	println(p)
	println(p1) // p和p1地址相同
}
