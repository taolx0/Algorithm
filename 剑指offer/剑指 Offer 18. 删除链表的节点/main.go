package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{Val: 0, Next: head}
	first := dummyNode
	second := dummyNode.Next
	for {
		if second.Val == val {
			first.Next = second.Next
			break
		}
		first = first.Next
		second = second.Next
	}
	return dummyNode.Next
}
