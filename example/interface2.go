package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee
	hasDegree bool
}

func main() {
	e1 := employee{
		name:      "John",
		age:       24,
		isMarried: false,
	}
	fmt.Println(e1)

	e2 := employee{
		name:      "Elissa",
		age:       24,
		isMarried: true,
	}
	fmt.Println(e2)

	m1 := manager{
		employee: employee{
			name:      "Ayşe",
			age:       45,
			isMarried: false,
		},
		hasDegree: true,
	}

	fmt.Println(m1)
}
