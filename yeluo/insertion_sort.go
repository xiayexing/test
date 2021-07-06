// 创建人：lg
// 创建时间：2020/7/15
package main

import "fmt"

func insertionSort(param []int) []int {
	for i := 1; i < len(param); i++ {
		preIndex := i - 1
		currentData := param[i] //这里要用一个变量接收当前值
		for ; preIndex >= 0 && param[preIndex] > currentData; preIndex-- {
			param[preIndex+1] = param[preIndex] //如果不保存当前值，此处前一个值会把当前值覆盖掉
		}
		if preIndex != i+1 {  // 如果preIndex没有发生移动(即 param[preIndex] > currentData 条件不满足，上面的循环没有执行)，则不进行插入操作
			param[preIndex+1] = currentData
		}
	}
	return param
}

func main() {
	p := []int{1, 2, 4, 6, 3, 7, 9, 5, 8}
	p1 := insertionSort(p)
	fmt.Println(p)
	fmt.Println(p1)
	println(p)
	println(p1) // p和p1地址相同
}
