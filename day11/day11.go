package main

import "fmt"

type coordinate struct {
	x int
	y int
}

func getSums(a [][]int) []int {
	var sums []int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			var elements []int
			var sum int

			coordinates := []coordinate{
				coordinate{i, j},
				coordinate{i + 1, j},
				coordinate{i + 2, j},
				coordinate{i + 1, j + 1},
				coordinate{i, j + 2},
				coordinate{i + 1, j + 2},
				coordinate{i + 2, j + 2},
			}
			for _, v := range coordinates {
				e := a[v.y][v.x]
				elements = append(elements, e)
			}
			for _, e := range elements {
				sum += e
			}
			sums = append(sums, sum)
		}
	}
	return sums
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var a [][]int
	var max int

	for i := 0; i < 6; i++ {
		var b, c, d, e, f, g int

		fmt.Scanf("%d %d %d %d %d %d\n", &b, &c, &d, &e, &f, &g)
		row := []int{b, c, d, e, f, g}
		a = append(a, row)
	}
	for i, s := range getSums(a) {
		if i == 0 {
			max = s
			continue
		}
		max = Max(max, s)
	}
	fmt.Println(max)
}
