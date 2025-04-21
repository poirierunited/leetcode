package main

import (
	"fmt"
	"leetcode/pkg/algorithms/sorting"
)

func main() {
	input := []int{4, 2, 4, 4, 3}

	sorting.QuickSort(input)

	fmt.Println(input)
}
