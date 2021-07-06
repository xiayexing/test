// 创建人：lg
// 创建时间：2021/7/2
package main

import (
	"test/gotype"
	"testing"
)

func TestRecursiveReverse(t *testing.T) {
	head := &gotype.LNode{}
	CreateNode(head, 10)
	PrintNode("Before: ", head)
	RecursiveReverse(head)
	PrintNode("After : ", head)
}

func TestInsertReverse(t *testing.T) {
	head := &gotype.LNode{}
	CreateNode(head, 10)
	PrintNode("Before: ", head)
	InsertReverse(head)
	PrintNode("After : ", head)
	InsertReverse(head)
}
