package main

import "fmt"

func main() {
	name := "Richard"
	x := 22
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", &x, &x)
	fmt.Println("-----------------------------")

	y := &x
	fmt.Println(y)
	fmt.Println(&y)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", &y, &y)
	fmt.Println("------------------------------")

	newName := &name
	fmt.Println(newName)

}
