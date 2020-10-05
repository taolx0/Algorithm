package main

import (
	"fmt"
	"sort"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		i := sort.SearchInts(nums, target)
		//fmt.Println(i)
		if i < len(nums) && target == nums[i] {
			return true
		}
	}
	return false
}

func main() {
	a := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	//fmt.Println(len(a))
	fmt.Println(findNumberIn2DArray(a, 18))
}
