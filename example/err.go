package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := evenNum(3)
	if err != nil {
		fmt.Println(err, result)
	} else {
		fmt.Println("Girilen çift sayi: ", result)
	}

}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return num, errors.New("HATA tek sayı girdiniz.")
	}
	return num, nil
}
