package main

func main() {
	/*
		----------ARRAY LİST && ARRAY LENGHT---------
		cities := [5]string{"İzmir", "Ankara", "İstanbul", "Manisa", "Bursa"}
		for i := 0; i < len(cities); i++ {
			fmt.Println(i, ":", cities[i])
		}
		fmt.Println("")
		country := [...]string{"Germany", "Turkey", "Usa", "England", "France"}
		fmt.Println("-------", len(country), "-------")
		fmt.Println("") */

	/*
		------ARRAY OLUŞTUR-------
		var myslc []int
		var eleman int
		fmt.Print("Kaç elemanlı array istiyorsunuz: ")
		fmt.Scan(&eleman)
		myslc = make([]int, eleman)
		var sayi1 int

		for i := 0; i < eleman; i++ {
			fmt.Print("Arrayin ", i, ". elemanını giriniz: ")
			fmt.Scanln(&sayi1)
			myslc[i] = sayi1
		}
		fmt.Println(myslc) */

	/*
		-----SIRALAMA-----
		underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		mySlc := underArray[2:6]
		mySlc1 := underArray[:6]
		mySlc2 := underArray[3:]
		mySlc3 := underArray[:] // == mySlc4 := underArray
		fmt.Println(mySlc)
		fmt.Println(mySlc1)
		fmt.Println(mySlc2)
		fmt.Println(mySlc3) */

	/*
			-----YENİ ARRAY---------
		mySlc := []int{1, 2, 3}
		mySlc1 := append(mySlc, 4, 5, 6, 7, 8, 9, 31, 45)
		fmt.Println(mySlc)
		fmt.Println(mySlc1) */

	/*
			-------APPEND---------
		mySlc := []int{1, 2, 3}
		mySlc2 := []int{4, 5}

		mySlc = append(mySlc, mySlc2...)
		fmt.Println(mySlc) */

	/*
		mySlc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		mySlc = mySlc[:len(mySlc)-4]
		fmt.Println(mySlc) */

	/*
				------INDEX------
		izArr := [5]string{"Bornova", "Buca", "Konak", "Karşıyaka", "Menemen"}

		for index, value := range izArr {
			fmt.Println(index, value)
		}
		fmt.Println("-------------")
		manArr := []string{"Akhisar", "Şehzadeler", "Gördes", "Soma", "Turgutlu"}

		for iasda, v := range manArr {
			fmt.Println(iasda, v)
		} */

	/*
			-------COPY------
		myslc := []int{1, 2, 3}
		myslc2 := make([]int, 2)

		fmt.Println(myslc)
		fmt.Println(myslc2)

		fmt.Println("")
		copy(myslc2, myslc)
		fmt.Println(myslc)
		fmt.Println(myslc2)
	*/

	/* 		-----MAP ILE YAZARLARI VE YAZARLARIN KITAPLARINI GOSTERMEK------
	myMap := map[string][]string{
		"Yaşar Kemal":     []string{"İnce Memed", "Yer Demir Gök Bakır"},
		"Sabahattin Ali":  []string{"Kuyucaklı Yusuf", "Kürk Mantolu Madonna", "Değirmen"},
		"Haruki Murakami": []string{"1Q84", "Dans Dans Dans", "Kumandanı Öldürmek"},
	}

	fmt.Println(myMap)
	fmt.Println(myMap["Yaşar Kemal"])
	fmt.Println(myMap["Haruki Murakami"][2])


	for key, value := range myMap {
		fmt.Println("Yazarımızın Adı: ", key)
		fmt.Println("Bazı Kitapları: ", value)
		for i, v := range value {
			fmt.Println("\t", i+1, v)
		}
	}
	*/
}
