package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var a, b, c, d ListNode
	var node *ListNode
	a = ListNode{Val: 1, Next: &b}
	b = ListNode{Val: 2, Next: &c}
	c = ListNode{Val: 3, Next: &d}
	d = ListNode{Val: 4, Next: nil}
	node = deleteNode2(&a, 3)
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
}
