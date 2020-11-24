package main

//单独处理头节点值是val的情况
func deleteNode2(head *ListNode, val int) *ListNode {
	resp := head
	if head.Val == val {
		return head.Next
	}
	pre := head
	head = head.Next
	for head != nil {
		if head.Val == val {
			pre.Next = pre.Next.Next
			break
		}
		pre = pre.Next
		head = head.Next
	}
	return resp
}
