// 创建人：lg
// 创建时间：2020/9/14
package main

import (
	"fmt"
	"test/gotype"
)

func SituReverse(node *gotype.LNode) {
	var pre, next *gotype.LNode
	if node == nil || node.Next == nil {return}
	cur := node.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	node.Next = pre
}

func main() {
	head := &gotype.LNode{}
	fmt.Println("Reverse")
	CreateNode(head, 8)
	PrintNode("Before: ", head)
	SituReverse(head)
	PrintNode("After: ", head)
}
