package main

import "fmt"

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}
func main() {
	var oddIndexeds, evenIndexeds [][]byte
	var strs []string
	var numStr int

	fmt.Scan(&numStr)
	for i := 0; i < numStr; i++ {
		s := ""
		fmt.Scan(&s)
		strs = append(strs, s)
	}

	for _, str := range strs {
		evenIndexed := []byte{}
		oddIndexed := []byte{}
		for i := 0; i < len(str); i++ {
			if isEven(i) {
				evenIndexed = append(evenIndexed, str[i])
				continue
			}
			oddIndexed = append(oddIndexed, str[i])
			continue
		}
		evenIndexeds = append(evenIndexeds, evenIndexed)
		oddIndexeds = append(oddIndexeds, oddIndexed)
	}
	for i := 0; i < len(oddIndexeds); i++ {
		str := ""
		str += string(evenIndexeds[i])
		str += " "
		str += string(oddIndexeds[i])
		fmt.Println(str)
	}
}
