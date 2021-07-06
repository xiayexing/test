// 创建人：lg
// 创建时间：2021/5/19
package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("panic")
}

// 考点：panic仅有最后一个可以被revover捕获
// 触发panic("panic")后顺序执行defer，但是defer中还有一个panic，所以覆盖了之前的panic("panic")