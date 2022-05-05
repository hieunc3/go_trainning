package main

import "fmt"

func main() {
	//1> create & initlize slice
	slice1 := []int{1, 2, 3, 4, 5}
	arrStr := [7]string{"This", "is", "the", "tutorial", "of", "Go", "language"}
	slice2 := arrStr[1:]
	/*
		- We create the slice2 from the given arrStr
		- The pointer of slice2 pointed to index 1 because the lower
		  bound of the slice is set to 1(index 1 of the arrStr) so it
		  starts accessing elements from index 1.
		- The length of the slice is
		- Any changes made in slice2 will reflect to arrStr
	*/
	//slice2[0] = "are"
	fmt.Println("My slice2:", slice2)
	fmt.Println("My arrStr:", arrStr)
	fmt.Println("My slice1:", slice1)

	//2> append to a slice
	cars := []string{"BMW", "Honda", "Ferrari", "Ford"}
	newLineCars := []string{"Vinfast", "Tesla"}

	fmt.Println("cars:", cars, "has old length", len(cars))
	fmt.Println(&cars)
	/*
		- how a slice become a dynamic length whereas slice are
		backed by arrays and arrays themeselves are of fixed length?
		- when new elements are appended to the slice, a new array
		is created.
		- the elements existing array are copied to this new array
		and a new slice reference for this new array is returned.
		- the capacity of the new slice is now twice that of the old
		slice.
	*/
	cars = append(cars, "Toyota")
	//append using ... operator
	cars = append(cars, newLineCars...)
	fmt.Println("cars:", cars, "has new length", len(cars))
	fmt.Printf("%v\n", &cars)

	//3> pass a slice to a function
	/*
		- a slice is passed to a func, even though it's passed by
		value, the pointer variable will refer to the same underlying
		array.
		-> Any changes made inside the func are visible outside the
		func too
	*/
	fmt.Println("Slice: slice1 before func call", slice1)
	subtactOne(slice1)
	fmt.Println("Slice: slice1 after func call", slice1)

	//4> memory optimisation
	/*
		- Slices hold a reference to the underlying array. As long
		as the slices is in memory, the arr cannot be garbage collected.
		-> This might be of concern when is comes to memory management.
		- lets assume that we have a very large arr and we interested in
		processing only a small part of it.
		- we create a slice from that array and start processing the
		slice. The important thing to be noted here is that the arr
		will still be in memory since the slice references it.
		-> using copy func to solve this problem (this way can use the new
		slice and the original arr can be garbage collected)
	*/
	needCars := cars[:len(cars)-1]
	carsCpy := make([]string, len(needCars))
	copy(carsCpy, needCars)
	fmt.Println(carsCpy)
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 1
	}
}
