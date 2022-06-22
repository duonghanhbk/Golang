package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(11 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(11 * time.Second)
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case s1 := <-output1:
			fmt.Println(s1)
			return
		case s2 := <-output2:
			fmt.Println(s2)
			return
		default:
			fmt.Println("no value received")
		}
	}

}
