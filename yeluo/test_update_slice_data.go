// 创建人：lg
// 创建时间：2021/5/10
package main

import "fmt"

type Label struct {
	Amount int
}

func main() {
	var res []Label
	res = append(res, Label{Amount: 1}) // 如果想要修改slice中对象的值，要把res中的对象定义为指针类型，即：[]*Label
	for _, r := range res {
		r.Amount = 2  // 这里的r是局部变量，对r的修改不影响res中的对象
	}
	fmt.Println(res[0].Amount)
}
