package main

import "fmt"

func main() {

	for i := 0; i <= 45; i++ {
		if i%2 == 0 {
			fmt.Println("2'ye tam bölünen: ", i)
		} else if i%3 == 0 {
			fmt.Println("3'e tam bölünen: ", i)
		} else if i%5 == 0 {
			fmt.Println("5'e tam bölünen: ", i)
		} else if i%i == 0 {
			fmt.Println("Asal sayı: ", i)
		}
	}
}
