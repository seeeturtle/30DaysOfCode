package main

import "fmt"

func main() {
	var strsLen int
	fmt.Scanf("%d", &strsLen)

	var strs = make([]string, strsLen)
	var x = make([]interface{}, len(strs))
	for i := range strs {
		x[i] = &strs[i]
	}
	n, _ := fmt.Scan(x...)
	strs = strs[:n]
	for i := strsLen; i > 0; i-- {
		fmt.Printf("%s ", strs[i-1])
	}
}
