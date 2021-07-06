// 创建人：lg
// 创建时间：2020/7/15
package main

import "fmt"

func selectSort(param []int) []int {
	for i := 0; i < len(param)-1; i++ {
		minIndex := i
		for j := i+1; j < len(param); j++ {
			if param[j] < param[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			param[minIndex], param[i] = param[i], param[minIndex]
		}
	}
	return param
}

func main() {
	p := []int{1, 2, 4, 6, 3, 7, 9, 5, 8}
	p1 := selectSort(p)
	fmt.Println(p)
	fmt.Println(p1)
	println(p)
	println(p1)  // p和p1地址相同
}


