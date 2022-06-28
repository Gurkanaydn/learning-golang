package main

import "fmt"

/*
 func main() {
	slices := [3]int{1, 10, 100}
	fmt.Println(slices)
	fmt.Println("-----------")
	double(slices)
	fmt.Println("-----------")

}

func double(slc [3]int) {
	for i := 0; i < len(slc); i++ {
		slc[i] *= 2

	}
	fmt.Println(slc)
} */

func main() {
	x := 5
	fmt.Println(x)
	double(&x)
	fmt.Println(x)
}
func double(num *int) {
	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)
}
