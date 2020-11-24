package main

import "fmt"

func main() {
	ints := []int{99, 68, 3, 89, 33, 55, 11}
	sort := selectionSort(ints)
	fmt.Println(sort)
}

func selectionSort(a []int) []int {
	length := len(a)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[min], a[i] = a[i], a[min]
	}
	return a
}
