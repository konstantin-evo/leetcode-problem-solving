package main

import "fmt"

func main() {
	input1 := "abcabcbb"
	result := lengthOfLongestSubstring(input1)
	fmt.Println("Length of the longest substring without repeating characters in", input1, "is", result)

	input2 := "bbbbb"
	result = lengthOfLongestSubstring(input2)
	fmt.Println("Length of the longest substring without repeating characters in", input2, "is", result)

	input3 := "pwwkew"
	result = lengthOfLongestSubstring(input3)
	fmt.Println("Length of the longest substring without repeating characters in", input3, "is", result)
}

func lengthOfLongestSubstring(s string) int {
	// 1. Initialize the variables
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	// 2. Iterate over each character s[i] in the input string s.
	for i := 0; i < len(s); i++ {
		// 2.1 Check if the current character s[i] has been encountered before (exists in charIndex).
		if lastIndex, found := charIndex[s[i]]; found && lastIndex >= start {
			// 2.2 Update the start index to exclude the previous occurrence of s[i] (start = lastIndex + 1).
			start = lastIndex + 1
		}
		// 2.3 Update the last index of s[i] in charIndex to the current index i.
		// + Calculate the length of the current substring (length = i - start + 1).
		charIndex[s[i]] = i
		length := i - start + 1
		// 2.4 If the length of the current substring is greater than the current maxLength,
		// update maxLength with the new length.
		if length > maxLength {
			maxLength = length
		}
	}

	// 3. Return the final maxLength as the result.
	return maxLength
}
