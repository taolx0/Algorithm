package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var temp *ListNode
	for head != nil {
		temp = &ListNode{
			Val:  head.Val,
			Next: temp,
		}
		head = head.Next
	}
	return temp
}
