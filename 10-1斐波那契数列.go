package main

import "fmt"

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	num := make([]int, n+1)
	num[0] = 0
	num[1] = 1
	for i := 2; i <= n; i++ {
		num[i] = num[i-1] + num[i-2]
	}
	return num[n]
}

func main() {
	fmt.Println(fib(45))
}
