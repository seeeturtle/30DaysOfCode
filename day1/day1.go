package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	var j uint64
	var e float64
	var t string
	var texts []string

	for scanner.Scan() {
		text := scanner.Text()
		texts = append(texts, text)
		if len(texts) == 3 {
			break
		}
	}
	j, _ = strconv.ParseUint(texts[0], 10, 64)

	e, _ = strconv.ParseFloat(texts[1], 64)

	t = texts[2]

	i += j
	d += e
	s += t

	fmt.Printf("%d\n", i)
	fmt.Printf("%.1f\n", d)
	fmt.Printf("%s\n", s)
}
