package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d\n", &n)

	fmt.Println(factorial(n))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
