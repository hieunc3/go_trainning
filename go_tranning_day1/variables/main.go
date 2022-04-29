package main

import (
	"fmt"
	"strings"
)

func main() {
	//Define variable using var keyword
	var myVariable1 int = 5
	var myVariable2 = "Hello world"
	var myVariable3 = 32.12
	var myVariable4 int

	//Go allows to declare multiple variables of the same type in the single declaration
	var myvariable5, myvariable6, myvariable7 int = 2, 454, 67

	fmt.Printf("The type & value of myVariable1 is: %T - %d \n", myVariable1, myVariable1)
	fmt.Printf("The type & value of myVariable2 is: %T - %s \n", myVariable2, myVariable2)
	fmt.Printf("The type & value of myVariable3 is: %T - %f \n", myVariable3, myVariable3)
	fmt.Printf("The type & value of myVariable4 is: %T - %d \n", myVariable4, myVariable4)
	fmt.Println(myvariable5, myvariable6, myvariable7)

	printBoundary()

	//Define variable using short variable declaration
	myvar1 := 39
	myvar2 := "GeeksforGeeks"
	myvar3 := 34.67
	fmt.Println(myvar1, myvar2, myvar3)
}

func printBoundary() {
	fmt.Println(strings.Repeat("=", 50))
}
