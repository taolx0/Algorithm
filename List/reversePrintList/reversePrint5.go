package main

//反转链表后，依次打印值
func reversePrint5(head *ListNode) []int {
	var resp []int
	var temp *ListNode
	for head != nil {
		node := head.Next //这句放在for循环开头，为了捕捉链表末尾的节点，结束for循环
		head.Next = temp
		temp = head
		head = node
	}
	for temp != nil {
		resp = append(resp, temp.Val)
		temp = temp.Next
	}
	return resp
}
