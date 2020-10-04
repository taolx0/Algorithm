package main

import "fmt"

func main() {
	ints := []int{99, 68, 3, 89, 33, 55, 11}
	sort := selectionSort(ints)
	fmt.Println(sort)
}

func selectionSort(a []int) []int {
	length := len(a)
	for j := 0; j < length-1; j++ {
		min := j
		for i := j + 1; i < length; i++ {
			if a[i] < a[min] {
				min = i
			}
		}
		a[j], a[min] = a[min], a[j]
	}
	return a
}
