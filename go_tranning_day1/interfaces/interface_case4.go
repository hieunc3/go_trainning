package main

import "fmt"

func findType(i interface{}) {
	switch value := i.(type) {
	case string:
		fmt.Println("String: ", value)
	case int:
		fmt.Println("Int: ", value)
	default:
		fmt.Println("Unknown type")
	}
}

func exampleInterface4() {
	//A type switch is used to compare the concrete type of an
	//interface against multiple types specified in various case
	//statement
	//-> help to check type of struct if it belong to one specific interface
	findType("Hieunguyen")
	findType(99)
	findType(99.9)
}
