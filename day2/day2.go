package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var mealCost, tipPercent, taxPercent float64
	var texts []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		texts = append(texts, text)
		if len(texts) >= 3 {
			break
		}
	}
	mealCost, _ = strconv.ParseFloat(texts[0], 64)

	tipPercent, _ = strconv.ParseFloat(texts[1], 64)

	taxPercent, _ = strconv.ParseFloat(texts[2], 64)

	tip := mealCost * (tipPercent / 100)
	tax := mealCost * (taxPercent / 100)

	totalCost := round(mealCost + tip + tax)

	fmt.Printf("The total meal cost is %d dollars.", totalCost)
}

func round(aFloat float64) int {
	if aFloat < 0 {
		return int(aFloat - 0.5)
	}
	return int(aFloat + 0.5)
}
