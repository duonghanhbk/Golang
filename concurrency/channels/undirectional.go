package main

import "fmt"

func sendData(data chan<- int) {
	data <- 10
}

func main() {
	sendCh := make(chan int)
	go sendData(sendCh)
	fmt.Println(<-sendCh)
}
