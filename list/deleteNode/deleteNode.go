package deleteNode

type ListNode struct {
	Val  int
	Next *ListNode
}

//假头节点法
//利用假头节点，避免头结点的特殊操作，创建一个前置节点指针用来做删除操作
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
