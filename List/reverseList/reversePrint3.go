package main

//将链表值依次存入切片，利用for循环反转
func reversePrint3(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var s []int
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}
