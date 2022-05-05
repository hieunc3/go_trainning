package main

import "fmt"

// func hello() {
// 	fmt.Println("Hello world")
// }

// func hello(done chan bool) {
// 	fmt.Println("Hello world goroutine")
// 	//send data to channel
// 	done <- true
// }

// func hello(done chan bool) {
// 	fmt.Println("hello goroutine is going to sleep 4s")
// 	time.Sleep(4 * time.Second)
// 	fmt.Println("hello goroutine awake after 4s and going to write to done")
// 	done <- true
// }

// func run() {
// 	time.Sleep(1 * time.Second)
// 	fmt.Print("Running")
// }

// func numbers() {
// 	for i := 1; i <= 5; i++ {
// 		time.Sleep(250 * time.Millisecond)
// 		fmt.Printf("%d", i)
// 	}
// }

// func alphabets() {
// 	for i := 'a'; i <= 'e'; i++ {
// 		time.Sleep(400 * time.Millisecond)
// 		fmt.Printf("%c", i)
// 	}
// }

//main Goroutine
func main() {

	//1> starting a Goroutines
	//create new goroutine run func hello()
	//go hello()
	/*
		- When a new Goroutine is started, the goroutine call returns immediatedly.
		  . Unlike functions, the control does not wait for the Goroutine to finish executing.
		- The control returns immediatedly to the next line of code after the Goroutine call
		  any return values from the Goroutine are ignored.
		- The main Goroutine should be running for any other Goroutines to run. If
		  the main Goroutine terminates then the program will be terminated and no
		  other Goroutine will run
	*/
	//we want main Goroutine sleep for 1s so that go hello() has enough time
	//to execute before the main Goroutine terminates.
	// time.Sleep(1 * time.Second)

	//2> starting multiple Goroutines
	// go numbers()
	// go alphabets()
	// time.Sleep(3000 * time.Millisecond)
	/*
		- numbers Goroutine sleeps initially for 250 milliseconds then
		  prints 1, then sleep again and prints 2 and the same cycle
		  happens till it prints 5.
		- alphabets Goroutine prints alphabets from a to e and has 400
		  milliseconds of sleep time
		- main Goroutine initiates the numbers and alphabets Goroutine,
		  sleeps for 3000 milliseconds and then terminates.
	*/
	/* Describe how it work?
	numbers Goroutine: sleep 250ms before start
	 0ms	250ms		500ms		750ms		1000ms		1250ms
			  1  	  	  2  	  	  3  	  	  4  	  	  5
	alphabets Goroutine: sleep 400ms before start
	 0ms			400ms				800ms				1200ms				1600ms				2000ms
	 				  a  				  b  				  c  				  d  				  e
	main Goroutine:
	------------------------------------------------------------------------------------------------------------------------> 3000ms (main terminate)
	output:
	0ms     250ms    400ms  500ms     750ms   800ms     1000ms    1200ms  1250ms       1600ms      2000ms            3000ms
			  1  	   a	  2  		3  		4  		  4  		c  		5			  d			  e				main terminate
	*/
	//fmt.Println("main terminated")

	//3> Goroutines & Channels
	/*
		- we use a sleep in the first section to make the main Goroutine
		wait for the hello Goroutine to finish.
		- isDone := <-done is blocking which means that unitl some
		Goroutine writes data to the done channel, the control will
		not moce to the next line of code
		- Hence this eliminates the need for the time.Sleep which was
		present in the original program to prevent the main Goroutine
		from exiting
		- now we have our main Goroutine blocked waiting for data
		on the done channel. The hello Goroutine receives this channel
		as a parameter, prints something and then writes to the done
		channel. When this write is complete, the main Goroutine receives
		the data from the done channel, it's unblocked and then the text
		main func is printed.
	*/
	//done := make(chan bool)
	// go hello(done)
	// isDone := <-done
	// if isDone {
	// 	fmt.Println("Run")
	// }

	// fmt.Println("Main going to call hello goroutine")
	// go hello(done)
	//run goroutine will sleep 1s before start
	// go run()
	// <-done
	// fmt.Println("main terminated")

	//4> another example
	/*
		input: 123
			squares = (1 * 1) + (2 * 2) + (3 * 3)
			cubes = (1 * 1 * 1) + (2 * 2 * 2) + (3 * 3 * 3)
			output = squares + cubes = 50

		we will structure the program such that the squares are
		calculated in a separate Goroutine and final summation happens
		in the main Goroutine
	*/
	// number := 123
	// sqrch := make(chan int)
	// cubech := make(chan int)
	// go calcSquares(number, sqrch)
	// go calcCubes(number, cubech)
	//main Goroutine wait for data from both these channels
	//once the data is received from both the channels -> print output
	// squares, cubes := <-sqrch, <-cubech
	// fmt.Println("Final output:", squares+cubes)

	//5> closing channels and for range loops on channels
	//senders have the ability to close the channel to notify receivers
	//that no more data will be sent on the channel
	// ch := make(chan int)
	// go producer(ch)
	// for {
	// 	v, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("Received:", v, ok)
	// }

	// for v := range ch {
	// 	fmt.Println("Received:", v)
	// }

	//6> another example
	number := 123
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

// func producer(chnl chan int) {
// 	defer close(chnl)
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// }

// func calcSquares(number int, squareop chan int) {
// 	sum := 0
// 	for number != 0 {
// 		digit := number % 10
// 		sum += digit * digit
// 		number /= 10
// 	}
// 	//we send the sum result to the channel
// 	squareop <- sum
// }

// func calcCubes(number int, cubeop chan int) {
// 	sum := 0
// 	for number != 0 {
// 		digit := number % 10
// 		sum += digit * digit * digit
// 		number /= 10
// 	}
// 	//we send the sum result to the channel
// 	cubeop <- sum
// }

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	defer close(dchnl)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	//we send the sum result to the channel
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	//we send the sum result to the channel
	cubeop <- sum
}
