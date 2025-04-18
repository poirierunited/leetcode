package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)

	fmt.Println(nums)
}

// time complexity == O(n)
func rotate(nums []int, k int) {

	// if k is 0, return the original array
	if k == 0 {
		return
	}

	arrLen := len(nums)
	// if nums is an empty array, return it
	if arrLen == 0 || arrLen == k {
		return
	}

	var result []int

	// normalize k

	k = k % arrLen
	pivot := len(nums) - k

	// slicing == O(1)
	rightSide := nums[0:pivot]
	leftSide := nums[pivot:]

	//copying all elements, appending = O(n)
	result = append(result, leftSide...)
	result = append(result, rightSide...)

	// coying all elements O(n)
	copy(nums, result)
}
