package main

import (
	"fmt"
	"strings"
)

func main() {

	//goVariable()
	goDataTypes()
}

func goVariable() {
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
func goDataTypes() {

	//bool type
	// bool1 := "Hello world"
	// bool2 := "Hello world"
	// bool3 := "hello world"

	// result1 := bool2 == bool3
	// result2 := bool1 == bool2

	// fmt.Printf("The type of result1 is %T \n", result1)
	// fmt.Printf("The type of result1 is %T \n", result2)

	//printBoundary()

	//strings type
	// mysliceByte := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
	// mysliceRune := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0073}
	// myString1 := string(mysliceByte)
	// myString2 := string(mysliceRune)

	// fmt.Println("Create string with slice of byte: ", myString1)
	// fmt.Println("Create string with slice of rune: ", myString2)

	// myString3 := "ｿｹｯﾄを作成する"
	// myString3Runes := []rune(myString3)
	// myString4 := "abcdefghi"

	// fmt.Println("String1: ", myString3)
	// fmt.Println("Length1: ", utf8.RuneCountInString(myString3))
	// printBytes(myString3)
	// printCharRune(myString3Runes)
	// printCharRune(myString3Runes[1:])

	// fmt.Println()

	// fmt.Println("String2: ", myString4)
	// fmt.Println("Length2: ", len(myString4))
	// printBytes(myString4)
	// printChars(myString4)
	// fmt.Println("After substring from 0:", myString4[1:])

	// printBoundary()

	//array type
	// - An array is of value type not of reference type
	// numbers := [...]int{1, 2, 3}
	// fmt.Println("Original array(Before):", numbers)

	// new_numbers := numbers

	// fmt.Println("New array(before):", new_numbers)

	// new_numbers[0] = 500

	// fmt.Println("New array(After):", new_numbers)

	// fmt.Println("Original array(After):", numbers)

	// arr1 := [3]int{9, 7, 6}
	// arr2 := [...]int{9, 7, 6}
	// arr3 := [3]int{9, 5, 3}

	// fmt.Println(arr1 == arr2)
	// fmt.Println(arr2 == arr3)
	// fmt.Println(arr1 == arr3)

	//- Creating a copy of an array by value
	// my_arr1 := [5]string{"Scala", "Perl", "Java", "Python", "Go"}

	// my_arr2 := my_arr1

	// fmt.Println("Array_1: ", my_arr1)
	// fmt.Println("Array_2:", my_arr2)

	// my_arr1[0] = "C++"

	// fmt.Println("\nArray_1: ", my_arr1)
	// fmt.Println("Array_2: ", my_arr2)

	//- Creating a copy of an array by reference
	// my_arr1 := [6]int{12, 456, 67, 65, 34, 34}

	// my_arr2 := &my_arr1

	// fmt.Println("Array_1: ", my_arr1)
	// fmt.Println("Array_2:", *my_arr2)

	// my_arr1[5] = 1000000

	// fmt.Println("\nArray_1: ", my_arr1)
	// fmt.Println("Array_2:", *my_arr2)

	//Slice type
	mySlice1 := make([]int, 3, 3)
	mySlice1 = []int{1, 2, 3}

	// Growing slices
	// - Using copy func
	mySlice2 := make([]int, len(mySlice1), (cap(mySlice1)+1)*2)
	copy(mySlice2, mySlice1)
	fmt.Println(mySlice2)
	fmt.Println(mySlice1)

	// - Using append func
	a := make([]int, 1)
	a = append(a, 1, 2, 3)

	// Creating an array
	arr := [4]string{"Geeks", "for", "Geeks", "GFG"}

	var my_slice_11 = arr[1:2]
	fmt.Println(my_slice_11)

}

//Utils
func printBoundary() {
	fmt.Println(strings.Repeat("=", 50))
}

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
}

func printCharRune(runes []rune) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println()
}
