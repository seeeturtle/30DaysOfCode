package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		p.age = 0
		return p
	}

	p.age = initialAge
	return p
}

func (p person) amIOld() {
	switch {
	case p.age < 13:
		fmt.Println("You are young.")
	case 13 <= p.age && p.age < 18:
		fmt.Println("You are a teenager.")
	default:
		fmt.Println("You are old.")
	}
}

func (p person) yearPasses() person {
	p.age++

	return p
}

func main() {
	var T, age int
	var ages []int

	// fmt.Scan(&T)
	ages = []int{-1, 10, 16, 18}
	T = 4

	for i := 0; i < T; i++ {
		// fmt.Scan(&age)
		age = ages[i]
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
