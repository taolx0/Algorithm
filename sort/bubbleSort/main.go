package main

import "fmt"

func main() {
	ints := []int{1, 68, 3, 89, 33, 55, 11}
	sort := bubbleSort(ints)
	fmt.Println(sort)
}

func bubbleSort(a []int) []int {
	length := len(a)
	for j := 0; j < length-1; j++ {
		for i := 0; i < length-j-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}
