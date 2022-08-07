package main

import (
	"container/list"
	"log"
)

//使用stack实现链表反转
func reversePrint1(head *ListNode) []int {
	var s []int
	stack := list.New()
	for head != nil {
		log.Println(head)
		stack.PushBack(head.Val)
		head = head.Next
		log.Println(head)
	}
	for stack.Len() > 0 {
		element := stack.Back()
		stack.Remove(element)
		s = append(s, element.Value.(int))
	}
	return s
}
