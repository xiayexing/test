// 创建人：lg
// 创建时间：2021/5/18
package main

import "fmt"

func main() {
	// list := new([]int)  // new()申请内存，并返回对应类型0值的指针对象。
	// list = append(list, 1) // 报错：无法将 'list' (类型 *[]int) 用作类型 []Type
	list := make([]int, 2) // make申请内存，并返回对应类型0值的对象。这里初始化为[0,0]
	list = append(list, 1)
	fmt.Println(list) // 这里输出的是[0,0,1]

	s := make([]int, 3, 8)
	a := s[:9]  // panic: runtime error: slice bounds out of range [:9] with capacity 8
	fmt.Println(s, a)
}

// 考点: make的第三个参数 make第一个参数是类型,第二个参数长度,第三个是预留分配空间
// 在长度为3的范围内正常赋值,若想使用到长度为8的空间,需要重新切片!但是切片长度不能超过8
