// 创建人：lg
// 创建时间：2021/5/19
package main

import "fmt"

var a = func(i int) {
	fmt.Println("x")
}

func main() {
	a := func(i int) {
		fmt.Println(i)
		if i > 0 {
			a(i-1)
		}
	}
	a(10)
}

// 考点: 匿名函数的调用 匿名函数是不能自己调用自己的, a(i - 1)其实是上面var定义的 a
// 结果:
//    10 x
