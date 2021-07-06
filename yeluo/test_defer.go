// 创建人：lg
// 创建时间：2021/5/18
package main

func main() {
	println(deferFunc1(1))
	println(deferFunc2(1))
	println(deferFunc3(1))
}

func deferFunc1(i int) (t int){  // 命名返回值，return操作会直接返回t，defer操作会影响t的值。
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func deferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
		println("fun2:in", t)
	}()
	println("fun2:out", t)
	return t  // 首先创建临时变量并赋值：s:=t,此时t=1；再执行defer操作：t=4；最后返回s的值
}

func deferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2  // 首先对t进行赋值：t = 2；然后执行defer操作t=3；最后再返回t的值
}

// 考点:defer和函数返回值
// 需要明确一点是defer需要在函数结束前执行。 函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数
// DeferFunc1有函数返回值t作用域为整个函数，在return之前defer会被执行，所以t会被修改，返回4;
// DeferFunc2函数中t的作用域为函数，返回1;
// DeferFunc3返回3

// 理解return 返回值的运行机制:
// 为了弄清上述两种情况的区别，我们首先要理解return 返回值的运行机制:
// return 并非原子操作，分为赋值，和返回值两步操作
// deferFunc2 : 实际上return 执行了两步操作，因为返回值没有命名，所以
// return 默认指定了一个返回值（假设为s），首先将i赋值给s,后续
// 的操作因为是针对i,进行的，所以不会影响s, 此后因为s不会更新，所以
// return s 不会改变
// 相当于：
// var i int
// s := i
// return s
// deferFunc1、deferFunc3 : 同上，s 就相当于 命名的变量i, 因为所有的操作都是基于
// 命名变量i(s),返回值也是i, 所以每一次defer操作，都会更新
// 返回值i