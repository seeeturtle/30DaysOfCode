package main

import (
	"fmt"
	"strconv"
)

func maxConOne(b string) int {
	var max int
	for i := 0; i < len(b); i++ {
		j := i
		length := 0
		if string(b[j]) == "1" {
			for {
				if j == 0 {
					length++
					j++
					continue
				}
				if j+1 > len(b) {
					break
				}
				if j > 0 {
					if string(b[j]) == "1" {
						length++
						j++
						continue
					}
					break
				}
			}
		}
		if length > max {
			max = length
		}
		i = j
	}
	return max
}

func main() {
	var n int64

	fmt.Scanf("%d\n", &n)
	fmt.Println(maxConOne(strconv.FormatInt(n, 2)))
}
