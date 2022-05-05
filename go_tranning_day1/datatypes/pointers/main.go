package main

import "fmt"

func main() {
	x := 7
	//we set y equal to pointer of x
	y := &x
	//&x where is the value 7 stored
	fmt.Println(y, x)
	//change the value
	*y = 8
	fmt.Println(y, x)

	toChange := "hello"
	var pointer *string = &toChange
	//valueStr will be the value of pointer variable
	// valueStr := *(*&pointer)
	// fmt.Println(pointer, valueStr)
	fmt.Println(pointer)
}
