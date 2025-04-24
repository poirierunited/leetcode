package main

import "fmt"

func main() {
	input := []int{5, 7, 7, 8, 8, 10}
	target := 10

	firstPosition := extendedBinarySearch(input, target, true)
	lastPosition := extendedBinarySearch(input, target, false)

	fmt.Printf("fist position: %d\nlast position: %d\n",
		firstPosition, lastPosition)
}

func extendedBinarySearch(nums []int, target int, findFirst bool) int {
	result := -1
	l, r := 0, len(nums)-1

	for l <= r {
		half := l + (r-l)/2

		if target > nums[half] {
			l = half + 1
		} else if target < nums[half] {
			r = half - 1
		} else {
			result = half

			if findFirst {
				r = half - 1
			} else {
				l = half + 1
			}
		}
	}

	return result
}
