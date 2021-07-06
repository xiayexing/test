// 创建人：lg
// 创建时间：2021/5/19
package main

import "fmt"

type configOne struct {
	Daemon string
}

func (c *configOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

func main() {
	c := &configOne{}
	s := c.String()
	fmt.Println(s)
}
// 考点:fmt.Sprintf
// 如果类型实现String()，％v和％v格式将使用String()的值。因此，对该类型的String()函数内的类型使用％v会导致无限递归。 编译报错：
// 结果
// runtime: goroutine stack exceeds 1000000000-byte limit
// runtime: sp=0xc0200e0330 stack=[0xc0200e0000, 0xc0400e0000]
// fatal error: stack overflow