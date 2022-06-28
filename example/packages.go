package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Sayı gir: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	fmt.Println("Girilen sayi: ", value)
}
