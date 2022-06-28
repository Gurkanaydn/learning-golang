package main

import "fmt"

func main() {
	/* var (
		firstName   string  = "Richard"
		lastName    string  = "Hendricks"
		height      float32 = 1.75
		age, weight int     = 18, 65

		isMarried bool = false
	) */

	firstName, lastName, height, age, weight, isMarried := "Richard", "Hendricks", 1.75, 18, 65, false
	/*
		var x int16 = 160
		var y int32
		y = int32(x)
		fmt.Println(y)
		fmt.Printf("")
		fmt.Printf("%T\n", y) */

	fmt.Println("Name: ", firstName, lastName)
	fmt.Println("Age: ", age)
	fmt.Println("Height: ", height)
	fmt.Println("weight: ", weight)
	fmt.Println("isMarried: ", isMarried)

	fmt.Println("")

	fmt.Printf("%T\n", firstName)
	fmt.Printf("%T\n", lastName)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", weight)
	fmt.Printf("%T\n", isMarried)
}
