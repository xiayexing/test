// 创建人：lg
// 创建时间：2021/7/2
package main

import "test/gotype"

func RecursiveReverseChild(node *gotype.LNode) *gotype.LNode {
	if node == nil || node.Next == nil {
		return node
	}
	newNode := RecursiveReverseChild(node.Next)
	node.Next.Next = node // 5->6, node:5, node.Next:6, node.Next.Next:6.Next   ==> 6->5
	node.Next = nil       // 5->nil
	return newNode
}

func RecursiveReverse(node *gotype.LNode) {
	firstNode := node.Next
	newFirstNode := RecursiveReverseChild(firstNode)
	node.Next = newFirstNode
}

func main() {
	head := &gotype.LNode{}
	CreateNode(head, 1000)
	PrintNode("Before: ", head)
	RecursiveReverse(head)
	PrintNode("After: ", head)
}
