package main

import (
	"fmt"
)

// Write a function that takes in a non-empty array of distinct integers and an integer representing a target sum.
// If any two numbers in the input array sum up to the target sum, the function should return them in an array,
// in any order. If no two numbers sum up to the target sum, the function should return an empty array.

func TwoNumberSum(array []int, target int) []int {
	// Declaring a new map
	nums := map[int]bool{}

	for index, num := range array {
		fmt.Println(index)
		potentialMatch := target - num
		//Assesing the map,getting the key, validating and returning when found->true
		if key, found := nums[potentialMatch]; found {
			fmt.Println(key)
			return []int{potentialMatch, num}
		}
		nums[num] = true
	}
	return []int{}
}
