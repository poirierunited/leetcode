package main

import (
	"fmt"
	"leetcode/pkg/algorithms"
)

func main() {
	input := []int{4, 2, 4, 4, 3}

	algorithms.QuickSort(input)

	fmt.Println(input)
}
