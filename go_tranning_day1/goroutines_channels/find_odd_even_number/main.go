package main

import "fmt"

func main() {
	numbers := []int{91, 42, 23, 14, 15, 76, 87, 28, 18, 95}
	chOdd := make(chan int)
	chEven := make(chan int)

	go odd(chOdd)
	go even(chEven)

	for _, value := range numbers {
		if value%2 != 0 {
			chOdd <- value
		} else {
			chEven <- value
		}
	}

	close(chEven)
	close(chOdd)
	fmt.Println("Run")
}

func odd(ch <-chan int) {
	//channel read only
	for v := range ch {
		fmt.Println("ODD: ", v)
	}
}

//ch chan<- int channel write only
func even(ch <-chan int) {
	for v := range ch {
		fmt.Println("EVEN:", v)
	}
}
