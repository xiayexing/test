// 创建人：lg
// 创建时间：2021/5/12
package main

import "fmt"

func main() {
	var c [2]int  // 声明数组未初始化时，是默认数据类型的零值。
	var d [3]int = [3]int{1,2}  // 为初始化的元素的值默认是对应类型的零值。
	fmt.Println(c)  // [0, 0]
	fmt.Println(d)  // [1, 2, 0]
}
