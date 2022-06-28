package main

import "fmt"

func main() {
	fmt.Println(topla(2, 5))
	fmt.Println(yaz("Hayırlı ", "Cumalar"))
}

func topla(x int, y int) int {
	sum := x + y
	return sum
}
func yaz(text1 string, text2 string) string {
	message := (text1 + text2 + " World")
	return message

}
