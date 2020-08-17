package main

import (
	"fmt"
	"sort"
)

func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	for i, numSize := 0, len(nums); i < numSize; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 1, 0, 2, 5, 2}
	num := findRepeatNumber(nums)
	fmt.Println("重复的一个数字为: ", num)
}
