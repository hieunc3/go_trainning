package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func process(ch chan string) {
	//sleep for 10.5s before start
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successfully"
}

func main() {

	// output1 := make(chan string)
	// output2 := make(chan string)

	// go server1(output1)
	// go server2(output2)
	// /*
	// 	- the select statement blocks until one of its cases is ready.
	// 	- the select statement will block for 3 seconds and will wait for server2 Goroutine to write to the output2 channel then terminate

	// */
	// select {
	// case s1 := <-output1:
	// 	fmt.Println("Get ", s1)
	// case s2 := <-output2:
	// 	fmt.Println("Get ", s2)
	// }

	ch := make(chan string)
	go process(ch)
	for {
		//the for block sleeps for 1s during the start of each iteration
		time.Sleep(1000 * time.Millisecond)
		select {
		/*
			During the first 10.5s the first case of the select statement will not be ready since the process Goroutine
			will write to the channel only after 10.5s
			no value received -> will print 10 times before receive value and return
		*/
		case v := <-ch:
			fmt.Println("receive value:", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}
