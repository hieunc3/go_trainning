package main

import "fmt"

func receiver(c <-chan int) { //channel read only
	for v := range c {
		fmt.Println(v)
	}
}

/*
	we create a c int channel and return it from the generator
*/
func generator() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func main() {
	c := generator()
	receiver(c)
}
