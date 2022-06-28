package main

import "fmt"

func main() { /*
		isMarried := map[string]int{
			"Ahmet":  50,
			"Mehmet": 40,
			"Ali":    30,
			"Efe":    20,
		}
		fmt.Println(isMarried["Ahmet"], isMarried["Ali"])
	*/
	isMarried := map[string]bool{
		"Ahmet":  true,
		"Mehmet": false,
		"Ali":    false,
		"Efe":    true,
	}
	fmt.Println(isMarried)
	delete(isMarried, "Ahmet")
	fmt.Println(isMarried)
	fmt.Println(len(isMarried))
}
