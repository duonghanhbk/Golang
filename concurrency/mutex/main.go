package main

import (
	"fmt"
	"sync"
)

var sumMutex = 0
var sumChannel = 0

// Solving the race condition using a mutex
func incrementMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	sumMutex += 1
	m.Unlock()
	wg.Done()
}

func incrementChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	sumChannel += 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementMutex(&w, &m)
	}
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementChannel(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of sumMutex", sumMutex)
	fmt.Println("final value of sumChannel", sumChannel)
}
