// 创建人：lg
// 创建时间：2021/5/18
package main

func testReturnName(x, y int) (sum int, err error) {
	return x+y, nil
}
// 在函数有多个返回值时，只要有一个返回值有指定命名，其他的也必须有命名。
// 如果返回值有有多个返回值必须加上括号；
// 如果只有一个返回值并且有命名也需要加上括号；