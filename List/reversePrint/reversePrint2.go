package main

//使用temp方法实现链表反转打印
func reversePrint2(head *ListNode) []int {
	var resp []int
	var temp *ListNode
	for head != nil {
		temp = &ListNode{
			Val:  head.Val,
			Next: temp,
		}
		head = head.Next
	}
	for temp != nil {
		resp = append(resp, temp.Val)
		temp = temp.Next
	}
	return resp
}
