// 创建人：lg
// 创建时间：2021/5/19
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	Name string
}

var m map[string]T
var n map[string]*T

func main() {
	m = make(map[string]T)
	n = make(map[string]*T)
	t := T{"LG"}
	m["t"] = t
	n["t"] = &t
	// m["t"].Name = "Lg"  // 编译报错：无法分配给 m["t"].Name
	n["t"].Name = "Lg" // 编译通过
	fmt.Println(m["t"])
	fmt.Println(n["t"])
	f := m["f"] // 访问map中不存在的key返回对应值类型的空值
	fmt.Println(f, reflect.TypeOf(f))
}

// 考点:map
// 编程报错cannot assign to struct field m["t"].Name in map。 因为m[“t”]不是一个普通的指针值，
// map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。 
// 定义的是var m map[string]T，注意哦T不是指针，而且map我们都知道是可以自动扩容的，那么原来的存储t的T可能在地址A，
// 但是如果map扩容了地址A就不是原来的T了，所以go就不允许我们写数据。你改为var m map[string]*T试试看。
