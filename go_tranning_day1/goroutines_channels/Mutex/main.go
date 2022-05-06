package main

import (
	"fmt"
	"sync"
)

var x = 0

func incrementMuxtex(wg *sync.WaitGroup, m *sync.Mutex) {
	/*
		the code is void of any race conditions since onlly one Goroutine is allowed to execute this piece of code at any
		point in time
	*/
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func incrementChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	//var m sync.Mutex
	/*
		- We have created a buffered channel of capacity 1 and this is passed to the increment Goroutine.
		- The buffered channel is used to ensure that only one Goroutine access the critical section of code which increments x.
		- This is done by passing true to the buffered channel before x is incremented. Since the buffered channel has a capacity of 1,
		all other Goroutines trying to write to this channel are blocked until the value is read from this channel -> allow only one Goroutine to access the critical section

	*/
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		//go incrementMuxtex(&w, &m)
		go incrementChannel(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x:", x)
}
