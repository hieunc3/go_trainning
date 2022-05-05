package main

import "fmt"

func main() {
	fmt.Println("=============== FUNCTIONS GO LANG ===============")
	fmt.Println()
	//Not allow to write a func inside a function main

	//FUNCTION ARGUMENTS
	//playAroundWithFunctionArguments()

	//VARIADIC FUNCTIONS
	//playAroundWithVariadicFunctions()

	//ANONYMOUS FUNCTION
	//playAroundWithAnonymousFunction()

}

//func playAroundWithAnonymousFunction() {
// func() {
// 	fmt.Println("We are using anonymous function!")
// }()
//pass anonymous func -> run another function
// callBackFunction := func(p, q string) string {
// 	return p + " & " + q
// }
// printString(callBackFunction)

//return an anonymous func from another functon
// value := GFG()
// fmt.Println(value("Welcome ", "to "))
//}

// func printString(i func(p, q string) string) {
// 	fmt.Println(i("Go", "Lang"))
// }

// func GFG() func(i, j string) string {
// 	myf := func(i, j string) string {
// 		return i + j + "Golang World"
// 	}
// 	return myf
// }

// func playAroundWithFunctionArguments() {
// 	// p, q := 10, 20
// 	// swapWithValue(p, q)
// 	// fmt.Printf("p: %v & q: %v", p, q)

// 	// swapWithRef(&p, &q)
// 	// fmt.Printf("p: %v & q: %v", p, q)
// }

// func playAroundWithVariadicFunctions() {
// 	numbers := []int{1, 2, 3, 4, 5}

// 	var sumNumbersSlice float64 = sumOfMultipleArguments(numbers...)
// 	var sumMultipleNumbers float64 = sumOfMultipleArguments(1, 2, 3, 4, 5)
// 	fmt.Printf("Total sum of list number (using slice): %v\n", sumNumbersSlice)
// 	fmt.Printf("Total sum of list number (using numbersg): %v\n", sumMultipleNumbers)
// }

// func sumOfMultipleArguments(nums ...int) float64 {
// 	var sum float64
// 	for _, v := range nums {
// 		sum += float64(v)
// 	}
// 	return sum
// }

// func swapWithValue(a, b int) {
// 	var o int = a
// 	a = b
// 	b = o
// }

// func swapWithRef(a, b *int) {
// 	var o int = *a
// 	*a = *b
// 	*b = o
// }
