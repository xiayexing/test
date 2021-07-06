// 创建人：lg
// 创建时间：2021/5/19
package main

import "fmt"

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			fmt.Println(&i, i)
		})
	}
	return funs
}

func test2(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
	// 考点：闭包引用相同变量*
	a, b := test2(100)
	a()
	b()
}

// 考点：闭包延迟求值
// for循环复用局部变量i，每一次放入匿名函数的应用都是想一个变量。 结果：
// 0xc042046000 2
// 0xc042046000 2
// 如果想不一样可以改为：
// func test() []func()  {
// 	var funs []func()
// 	for i:=0;i<2 ;i++  {
// 		x:=i
// 		funs = append(funs, func() {
// 			println(&x,x)
// 		})
// 	}
// 	return funs
// }
// 首先得注意 闭包和匿名函数的区别! 另外,看执行顺序,会发现test先执行,但是test()里的闭包并没有执行,只有在range funs的时候 闭包才执行...那时候里面的i最后是啥值,结果就是啥值...
