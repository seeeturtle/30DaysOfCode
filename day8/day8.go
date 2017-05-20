package main

import (
	"fmt"
	"log"
)

func readIntoMap(n int) (map[string]int, error) {
	m := map[string]int{}
	for i := 0; i < n; i++ {
		var k string
		var v int
		_, err := fmt.Scanf("%s %d\n", &k, &v)
		if err != nil {
			return map[string]int{}, err
		}
		m[k] = v
	}
	return m, nil
}

func readIntoSlice() ([]string, error) {
	var strs []string
	for {
		var str string

		_, err := fmt.Scanf("%s\n", &str)
		if err != nil {
			break
		}
		strs = append(strs, str)
	}
	return strs, nil
}

func query(m map[string]int, queries ...string) {
	for _, q := range queries {
		if val, exist := m[q]; exist {
			fmt.Printf("%s=%d\n", q, val)
		} else {
			fmt.Println("Not found")
		}
	}
}

func main() {
	var n int

	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	phoneBook, err := readIntoMap(n)
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	queries, err := readIntoSlice()
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	query(phoneBook, queries...)
}
