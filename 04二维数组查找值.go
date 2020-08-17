package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {

}

func main() {
	a := [][]int{
		{0, 1, 2, 4},   //index is 0
		{4, 5, 6, 7},   //index is 1
		{8, 9, 10, 11}, //index is 2
	}
	fmt.Println(findNumberIn2DArray(a, 10))
}
