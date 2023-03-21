package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int) []int {
	result := make([]int, 0, n)

	result = append(result, 0, 1)

	for i := 2; i < n; i++ {
		result = append(result, result[i-1]+result[i-2])
	}

	return result
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	result := fibonacci(n)

	fmt.Println(result)
}
