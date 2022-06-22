package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func numbers(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}
func alphabets(wg *sync.WaitGroup) {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func main() {
	go say("world")
	// Step 1: Create variable has type sync.WaitGroup: wg
	// Step 2: Add numbers of goroutine
	// Step 3: Each goroutine call wg.Done() before return
	// Step 4: Call method wg.wait()

	var wg sync.WaitGroup
	wg.Add(2)
	go numbers(&wg)
	go alphabets(&wg)
	wg.Wait()
	fmt.Println("main terminated")
}
