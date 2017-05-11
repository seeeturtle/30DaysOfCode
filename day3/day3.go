package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Integer int

func (i Integer) isEven() bool {
	half := i / 2.
	typeOfHalf := reflect.TypeOf(half).String()
	if typeOfHalf == "Integer" {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(str)

	num := Integer(n)

	switch {
	case !num.isEven():
		fmt.Println("Weird")

	case num.isEven() && 2 <= num && num <= 5:
		fmt.Println("Not Weird")

	case num.isEven() && 6 <= num && num <= 20:
		fmt.Println("Weird")

	case num.isEven() && 20 < num:
		fmt.Println("Not Weird")
	}
}
