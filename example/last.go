package main

import "fmt"

func hello(chan1 chan string) {
	chan1 <- "Hello From channel"
}

func main() {
	myChannel := make(chan string)
	go hello(myChannel)
	fmt.Println(myChannel)
	fmt.Println(<-myChannel)
}
