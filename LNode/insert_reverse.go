// 创建人：lg
// 创建时间：2021/7/2
package main

import "test/gotype"

func InsertReverse(node *gotype.LNode) {
	if node == nil || node.Next== nil {
		return
	}
	var next *gotype.LNode
	cur := node.Next.Next  // 从第二个节点开始
	node.Next.Next = nil  // 把第一个节点的Next指向nil
	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}

