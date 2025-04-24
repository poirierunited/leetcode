package main

import "fmt"

func main() {
	input1 := []int{1, 2, 2, 3, 4, 4, 4, 4, 4, 5, 5, 5}
	fmt.Printf("%d\n", findElementWithHighestFrecuency(input1))

}

func findElementWithHighestFrecuency(numbers []int) int {
	freq := make(map[int]int)
	maxFreq := 0
	result := 0

	// time complexity: O(n)
	// space complexity: O(n)
	for _, num := range numbers {
		freq[num]++
		if freq[num] > maxFreq {
			maxFreq = freq[num]
			result = num
		}
	}

	return result
}
