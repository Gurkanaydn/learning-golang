package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--TOPLAMA--")
	topla()
	fmt.Println("--CIKART--")
	cikart()
	fmt.Println("--CARP--")
	carp()
	fmt.Println("--BOL--")
	bol()
	fmt.Println("--KARESINIAL--")
	karesinial()

}
func topla() {
	var sayi1 float64
	var sayi2 float64
	fmt.Print("İlk sayıyı gir: ")
	fmt.Scanln(&sayi1)
	fmt.Print("İkinci sayıyı gir: ")
	fmt.Scanln(&sayi2)
	var sonuc float64 = sayi1 + sayi2
	fmt.Println("Toplam sonuc: ", sonuc)

}
func cikart() {
	var sayi1 float64
	var sayi2 float64
	fmt.Print("İlk sayıyı gir: ")
	fmt.Scanln(&sayi1)
	fmt.Print("İkinci sayıyı gir: ")
	fmt.Scanln(&sayi2)
	var sonuc float64 = sayi1 - sayi2
	fmt.Println("Cikan sonuc: ", sonuc)

}
func carp() {
	var sayi1 float64
	var sayi2 float64
	fmt.Print("İlk sayıyı gir: ")
	fmt.Scanln(&sayi1)
	fmt.Print("İkinci sayıyı gir: ")
	fmt.Scanln(&sayi2)
	var sonuc float64 = sayi1 * sayi2
	fmt.Println("Carpim sonuc: ", sonuc)
}
func bol() {
	var sayi1 float64
	var sayi2 float64
	fmt.Print("İlk sayıyı gir: ")
	fmt.Scanln(&sayi1)
	fmt.Print("İkinci sayıyı gir: ")
	fmt.Scanln(&sayi2)
	var sonuc float64 = sayi1 / sayi2
	fmt.Println("Bolum sonuc: ", sonuc)
}

func karesinial() {
	var sayi float64
	fmt.Print("Sayi giriniz: ")
	fmt.Scan(&sayi)
	sonuc := math.Pow(sayi, sayi)
	fmt.Println("Sayiların karesi: ", sonuc)
}
