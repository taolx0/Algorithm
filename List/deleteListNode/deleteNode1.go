package main

//假头节点法
//利用假头节点，避免头结点的特殊操作，创建一个前置结点指针用来做删除操作
func deleteNode1(head *ListNode, val int) *ListNode {
	temp := &ListNode{
		Val:  -1,
		Next: head,
	}
	resp := temp
	for head != nil {
		if head.Val == val {
			temp.Next = temp.Next.Next
			break
		}
		head = head.Next
		temp = temp.Next
	}
	return resp.Next
}
