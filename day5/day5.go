package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scanf("%d", &num)

	for i := 1; i < 11; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
