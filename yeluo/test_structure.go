// 创建人：lg
// 创建时间：2021/5/12
package main

type (
	A struct {
		v int
	}

	// 定义结构体B，嵌入结构体A作为匿名字段
	B struct {
		A
	}

	// 定义结构体C，嵌入结构体A的指针作为匿名字段
	C struct {
		*A
	}
)

func (a *A) setV(v int) {
	a.v = v
}

func (a A) getV() int {
	return a.v
}

func (b B) getV() int {
	return b.v
}

func (c *C) getV() int {
	return c.v
}

func main() {
	a := A{}
	b := B{}    // 初始化结构体B，其内匿名字段A默认零值是A{}
	c := C{&A{}}    // 初始化结构体C，其内匿名指针字段*A默认零值是nil，需要初始化赋值

	println(a.v)

	// 结构体A嵌入B，A内字段自动提升到B
	println(b.v)

	// 结构体指针*A嵌入C，*A对应结构体内字段自动提升到C
	println(c.v)

	a.setV(3)
	b.setV(5)
	c.setV(7)
	println(a.getV(), b.A.getV(), c.A.getV())
	println(a.getV(), b.getV(), c.getV())
}
