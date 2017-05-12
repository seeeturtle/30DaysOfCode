package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Integer int

func (i Integer) isEven() bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func (i Integer) isWeird() string {
	switch {
	case !i.isEven():
		return "Weird"

	case i.isEven() && 2 <= i && i <= 5:
		return "Not Weird"

	case i.isEven() && 6 <= i && i <= 20:
		return "Weird"

	case i.isEven() && 20 < i:
		return "Not Weird"
	}
	return ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(str)
	num := Integer(n)
	fmt.Println(num.isWeird())
}
