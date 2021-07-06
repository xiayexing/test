// 创建人：lg
// 创建时间：2021/7/5
package main

import (
	"fmt"
	"test/gotype"
)

// 定义用于算法测试的公用方法

// 创建链表
func CreateNode(node *gotype.LNode, max int) {
	cur := node
	for i := 1; i <= max; i++ {
		cur.Next = &gotype.LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

// 打印链表
func PrintNode(info string, node *gotype.LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

