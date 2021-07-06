// 创建人：lg
// 创建时间：2021/5/19
package main

type S struct {
}

func f(x interface{}) {

}

func g(x *interface{}) {

}

func main() {
	s := S{}
	p := &s
	f(s)
	// g(s)  // 编译报错
	f(p)
	// g(p)  // 编译报错
	var i interface{}
	g(&i)

}

// 考点:interface
//
// 看到这道题需要第一时间想到的是Golang是强类型语言，interface是所有golang类型的父类，类似Java的Object。
// 函数中func f(x interface{})的interface{}可以支持传入golang的任何类型，包括指针，
// 但是函数func g(x *interface{})只能接受*interface{}.
