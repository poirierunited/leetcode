package searching

// time complexity: O(log n)

func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1

	for low <= high {
		half := low + (high-low)/2

		if target > nums[half] {
			low = half + 1
		} else if target < nums[half] {
			high = half - 1
		} else {
			return half
		}
	}

	return -1
}
