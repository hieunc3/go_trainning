package main

import "fmt"

func main() {
	//array type
	// - An array is of value type not of reference type
	numbers := [...]int{1, 2, 3}
	fmt.Println("Original array(Before):", numbers)

	new_numbers := numbers

	fmt.Println("New array(before):", new_numbers)

	new_numbers[0] = 500

	fmt.Println("New array(After):", new_numbers)

	fmt.Println("Original array(After):", numbers)

	arr1 := [3]int{9, 7, 6}
	arr2 := [...]int{9, 7, 6}
	arr3 := [3]int{9, 5, 3}

	fmt.Println(arr1 == arr2)
	fmt.Println(arr2 == arr3)
	fmt.Println(arr1 == arr3)

	//- Creating a copy of an array by value
	my_arr1 := [5]string{"Scala", "Perl", "Java", "Python", "Go"}

	my_arr2 := my_arr1

	fmt.Println("Array_1: ", my_arr1)
	fmt.Println("Array_2:", my_arr2)

	my_arr1[0] = "C++"

	fmt.Println("\nArray_1: ", my_arr1)
	fmt.Println("Array_2: ", my_arr2)

	//- Creating a copy of an array by reference
	my_arr3 := [6]int{12, 456, 67, 65, 34, 34}

	my_arr4 := &my_arr3

	fmt.Println("Array_1: ", my_arr3)
	fmt.Println("Array_2:", *my_arr4)

	my_arr3[5] = 1000000

	fmt.Println("\nArray_1: ", my_arr3)
	fmt.Println("Array_2:", *my_arr4)
}
