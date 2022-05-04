package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	c := make(chan []int)
	//Rank Sort (Depend on sorted slice)
	//fmt.Printf("Slice rank: %v \n", rankWithSortedSlice(numbers))

	//Rank Sort
	// fmt.Println(strings.Repeat("=", 60))
	// go rankSortWithOrginSlice(numbers, c)
	// rankList := <-c
	// printSlice(rankList)

	go rankSortWithMap(numbers, c)
	rankList := <-c
	printSlice(rankList)
}

// func rankWithSortedSlice(nums []int) []int {
// 	var temp int
// 	sliceRank := make([]int, len(nums))
// 	//Sort a copy numbers slice
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] >= nums[j] {
// 				temp = nums[i]
// 				nums[i] = nums[j]
// 				nums[j] = temp
// 			}
// 		}
// 	}
// 	fmt.Printf("Slice after sort: %v \n", nums)
// 	//Create default rank for the slice
// 	for i := 0; i < len(nums); i++ {
// 		sliceRank[i] = i + 1
// 	}

// 	//Update the rank for the slice
// 	for j := 0; j < len(nums)-1; j++ {
// 		if nums[j] == nums[j+1] {
// 			sliceRank[j+1] = j + 1
// 		}
// 	}

// 	return sliceRank
// }

// func rankSortWithOrginSlice(nums []int, rank_list chan []int) {
// 	numsCopy := make([]int, len(nums))
// 	rankArr := make([]int, len(nums))
// 	copy(numsCopy, nums)
// 	sort.Ints(numsCopy)

// 	for i := 0; i < len(nums); i++ {
// 		var count int = 1
// 		for j := 0; j < len(nums); j++ {
// 			if nums[i] > numsCopy[j] {
// 				count++
// 			}
// 		}
// 		rankArr[i] = count
// 	}
//  printSlice(nums)
// 	rank_list <- rankArr
// }

func rankSortWithMap(nums []int, rank_list chan []int) {
	defer close(rank_list)
	arrMap := make(map[int]int)
	sortedArr := make([]int, len(nums))
	arrRank := make([]int, len(nums))
	copy(sortedArr, nums)
	sort.Ints(sortedArr)
	rank := 1
	for _, s := range sortedArr {
		if _, ok := arrMap[s]; !ok {
			arrMap[s] = rank
			rank++
		}
	}
	for i, v := range nums {
		for key, element := range arrMap {
			if v == key {
				arrRank[i] = element
			}
		}
	}
	printSlice(nums)
	rank_list <- arrRank
}

func printSlice(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%6d", arr[i])
	}
	fmt.Println()
}
