// 创建人：lg
// 创建时间：2021/5/18
package main

import "fmt"

func main() {
	// i := GetValue()
	type itf interface {
	}
	var i2 itf = 1
	// switch i.(type) {  // 编译失败，因为type只能使用在interface
	switch i2.(type) {  // 编译成功，输出int
	case int:
		fmt.Println("int")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("unknown")
	}
}

func GetValue() int {
	return 1
}
