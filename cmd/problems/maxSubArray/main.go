package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(input))
}

// time: O(n)
// space: O(2) ==> O(1)
func maxSubArray(nums []int) int {
	//space: O(2)
	currentSum := 0
	maxSum := math.MinInt

	// time: O(n)
	// space: --
	for _, value := range nums {
		currentSum += value
		maxSum = max(currentSum, maxSum)

		if currentSum < 0 {
			currentSum = 0
		}
	}

	return maxSum
}

// time: O(1) --
// space: --
func max(a, b int) int {
	if a < b {
		return b
	}

	if a > b {
		return a
	}
	return a
}
