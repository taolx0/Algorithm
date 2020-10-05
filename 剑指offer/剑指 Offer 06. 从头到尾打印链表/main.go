package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	s := []int{}
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	//fmt.Println(s)
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}

func main() {
	var a, b ListNode
	a.Val, b.Val = 1, 2
	a.Next = &b
	fmt.Println(reversePrint(&a))
}
