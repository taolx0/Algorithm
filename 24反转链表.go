package main

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
