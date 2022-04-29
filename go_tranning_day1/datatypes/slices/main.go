package main

import "fmt"

func main() {
	mySlice1 := []int{1, 2, 3}
	fmt.Println(mySlice1)
	// Growing slices
	// - Using copy func
	mySlice2 := make([]int, len(mySlice1), (cap(mySlice1)+1)*2)
	copy(mySlice2, mySlice1)
	fmt.Println(mySlice2)
	fmt.Println(mySlice1)

	// - Using append func
	a := make([]int, 1, 6)
	a = append(a, 1, 2, 3)
	fmt.Printf("Slice : %v \n", a)
	// Creating an array
	arr := [4]string{"Geeks", "for", "Geeks", "GFG"}

	var my_slice_11 = arr[1:2]
	fmt.Println(my_slice_11)
}
