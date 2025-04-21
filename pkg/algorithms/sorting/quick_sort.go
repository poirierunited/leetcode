package sorting

// QuickSort sorts an integer slice in ascending order using the quicksort algorithm.
// It modifies the slice in-place.
func QuickSort(array []int) {
	if len(array) <= 1 {
		return
	}
	quickSort(array, 0, len(array)-1)
}

// quickSort is the internal recursive implementation of the quicksort algorithm
func quickSort(array []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	pivot := partition(array, startIndex, endIndex)
	quickSort(array, startIndex, pivot-1) // Sort elements before pivot
	quickSort(array, pivot+1, endIndex)   // Sort elements after pivot
}

// partition divides the array into two sub-arrays and returns the pivot index
func partition(arr []int, startIndex, endIndex int) int {
	pivot := arr[endIndex]
	i := startIndex - 1

	for j := startIndex; j < endIndex; j++ {
		if arr[j] <= pivot {
			i++
			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
		}
	}

	temp := arr[endIndex]
	arr[endIndex] = arr[i+1]
	arr[i+1] = temp

	return i + 1
}
