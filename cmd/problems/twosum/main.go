package main

import (
	"fmt"
)

// Problem: Two Sum
// Given an array of integers nums and an integer target, return indices of the two numbers in nums such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

func twoSum(nums []int, target int) []int {
	// Create a map to store number:index pairs
	numMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Calculate the complement needed
		complement := target - num

		// Check if complement exists in map
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}

		// Add current number and its index to map
		numMap[num] = i
	}

	// If no solution is found (shouldn't happen according to problem constraints)
	return []int{}
}

func main() {
	// Test cases
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	// Run test cases
	for i, tc := range testCases {
		result := twoSum(tc.nums, tc.target)

		// Compare results
		if len(result) != len(tc.expected) {
			fmt.Printf("Test case %d failed: expected %v, got %v\n", i+1, tc.expected, result)
			continue
		}

		match := true
		for j := range result {
			if result[j] != tc.expected[j] {
				match = false
				break
			}
		}

		if !match {
			fmt.Printf("Test case %d failed: expected %v, got %v\n", i+1, tc.expected, result)
		} else {
			fmt.Printf("Test case %d passed!\n", i+1)
		}
	}
}
