package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var a, b, c ListNode
	a = ListNode{Val: 1, Next: &b}
	b = ListNode{Val: 2, Next: &c}
	c = ListNode{Val: 3, Next: nil}
	ints := reversePrint5(&a)
	fmt.Println(ints)
}
