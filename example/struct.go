package main

import (
	"fmt"
)

func main() {

	/* var employee struct {
		firstName string
		lastName  string
		age       int
		isMarried bool
	}

	fmt.Println(employee)
	fmt.Printf("%#v\n", employee)
	fmt.Println(employee.age)
	fmt.Println("")
	employee.firstName = "Richard"
	employee.lastName = "Hendricks"
	employee.age = 25
	employee.isMarried = false

	fmt.Println(employee)
	fmt.Printf("%#v\n", employee) */

	/* type employee struct {
		firstName string
		lastName  string
		age       int
		isMarried bool
	}
	var e1 employee
	e1.firstName = "Richard"
	e1.lastName = "Hendricks"
	e1.age = 25
	e1.isMarried = false

	fmt.Println(e1)
	fmt.Println(e1.firstName, e1.lastName)
	fmt.Println("")

	var e2 employee
	e2.firstName = "Richard"
	e2.lastName = "Hendricks"
	e2.age = 39
	e2.isMarried = true
	fmt.Println(e2)
	fmt.Println(e2.firstName, e2.lastName)
	fmt.Println("")

	e3 := employee{
		firstName: "Richard",
		lastName:  "Hendricks",
		age:       19,
		isMarried: false,
	}
	fmt.Println(e3)
	fmt.Println(e3.firstName, e3.lastName)
	fmt.Println("")

	fmt.Println(e1, e2, e3) */

	type employee struct { // underlying type
		name      string
		age       int
		isMarried bool
	}

	type manager struct {
		employee
		hasDegree bool
	}

	e1 := employee{
		name:      "Ahmet",
		age:       40,
		isMarried: true,
	}
	fmt.Println(e1)

	m1 := manager{
		employee: employee{
			name:      "Ali",
			age:       30,
			isMarried: false,
		},
		hasDegree: true}

	/* m1 := manager{}
	m1.name = "Ayşe"
	m1.age = 28
	m1.isMarried = false
	m1.hasDegree = true */

	fmt.Println(m1)

}
