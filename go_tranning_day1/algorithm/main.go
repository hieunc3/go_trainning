package main

import "fmt"

func main() {
	numbers := []int{1, 4, 3, 2, 5, 7, 10, 3, 7, 8}
	//Rank Sort (Depend on sorted slice)
	fmt.Printf("Slice rank: %v \n", rankWithSortedSlice(numbers))
	//Rank Sort
	fmt.Printf("Slice rank: %v \n", rankSortWithOrginSlice(numbers))
}

func rankWithSortedSlice(nums []int) []int {
	var temp int
	sliceRank := make([]int, len(nums))
	//Sort a copy numbers slice
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] >= nums[j] {
				temp = nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	fmt.Printf("Slice after sort: %v \n", nums)
	//Create default rank for the slice
	for i := 0; i < len(nums); i++ {
		sliceRank[i] = i + 1
	}

	//Update the rank for the slice
	for j := 0; j < len(nums)-1; j++ {
		if nums[j] == nums[j+1] {
			sliceRank[j+1] = j + 1
		}
	}

	return sliceRank
}

func rankSortWithOrginSlice(nums []int) []int {

}
