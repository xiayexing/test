// 创建人：lg
// 创建时间：2021/5/19
package main

import "fmt"

var Ta int64 = aa()

func aa() int64 {
	fmt.Println("calling a()")
	return 2
}

func init() {
	fmt.Println("init in main.go")
}

func main() {
	fmt.Println("calling main")
}

// 考点: Golang加载过程
// 在一个go文件中， 初始化顺序规则：
// (1) 引入的包
// (2) 当前包中的变量常量
// (3) 当前包的init
// (4) main函数
// 注意：
// 1.当前go源文件中， 每一个被Import的包， 按其在源文件中出现顺序初始化。
// 2.如果当前包有多个init在不同的源文件中， 则按源文件名以字典序从小到大排序，小的先被执行到， 同一包且同一源文件中的init,则按其出现在文件中的先后顺序依次初始化； 当前包的package level变量常量也遵循这个规则； 其实准确来说，应是按提交给编译器的源文件名顺序为准，只是在提交编译器之前， go命令行工具对源文件名按字典序排序了。
// 3.init只可以由go runtime自已调用， 我们在代码中不可以显示调用，也不可以被引用，如赋给a function variable。
// 4.包A 引入包B , 包B又引入包C, 则包的初始化顺序为： C -> B -> A
// 5.引入包，必须避免死循环，如 A 引 B , B引C, C引A.
// 6.一个包被其它多个包引入，如A -> B ->C 和 H -> I -> C , C被其它包引了2次， 但是注意包C只被初始化一次。
// 7.另一个大原则， 被依赖的总是先被初始化，当然呀。
// 8.main包总是被最后一个初始化，这很好理解，因为它总是依赖别的包。
// 所以: calling a() init in main.go calling main