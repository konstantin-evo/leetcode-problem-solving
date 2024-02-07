package main

import "fmt"

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	runTwoSumExample(nums1, target1)
}

func runTwoSumExample(nums []int, target int) {
	result := twoSum(nums, target)
	fmt.Printf("Input: nums = %v, target = %d\n", nums, target)
	fmt.Printf("Output: %v\n", result)
	fmt.Println("---")
}

func twoSum(nums []int, target int) []int {
	// Map to store the indices of elements
	indexMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Calculate the difference between target and current element
		complement := target - num

		// Check if complement exists in the map
		if index, found := indexMap[complement]; found {
			// Return the indices of the two numbers
			return []int{index, i}
		}

		// Add the current element and its index to the map
		indexMap[num] = i
	}

	// It is guaranteed that there is exactly one solution, so no need for error handling
	return nil
}
