package main

import "fmt"

func fibonacci(ch chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x: //write to channel ch
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// func print(n int, t chan bool, c chan int) {
// 	for i := 0; i < n; i++ {
// 		fmt.Println(<-c)
// 	}
// 	t <- false
// }

func main() {
	n := 10
	ch := make(chan int)
	quit := make(chan bool)
	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch) //read from channel ch
		}
		quit <- false
	}(n)
	// go print(n, quit, ch)
	fibonacci(ch, quit)
}
