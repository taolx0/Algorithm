package main

//可以学习语言特性，不建议实际使用
//利用defer“后进先出”的特性实现链表反转打印
func reversePrint4(head *ListNode) (resp []int) {
	for head != nil {
		temp := head
		defer func() {
			resp = append(resp, temp.Val)
		}()
		head = head.Next
	}
	return resp
}
