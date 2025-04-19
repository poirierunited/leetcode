package main

import "fmt"

func main() {
	/* input := []int{1, 2, 2, 3, 4, 4, 4, 4, 4, 5, 5, 5} */
	input2 := []int{}
	result := findElementWithHighestFrecuency(input2)

	fmt.Println(result)
}

// time: O(n) + O(M) == O(M+n), where worst case M == n, so O(2n) == O(n)
// space: O(M) + O(1) == O(M)
func findElementWithHighestFrecuency(elements []int) int {

	if len(elements) == 0 {
		return -1
	}

	frecuencyMap := createFrecuencyMap(elements)
	return findMostRepeatedElement(frecuencyMap)
}

// time: O(n)
// space: O(M)
func createFrecuencyMap(elements []int) map[int]int {
	frecuencyMap := make(map[int]int)

	for _, element := range elements {
		if _, found := frecuencyMap[element]; found {
			frecuencyMap[element]++
		} else {
			frecuencyMap[element] = 1
		}
	}

	return frecuencyMap
}

// O(M) - where M = number of elements without its repetitions
// space: O(1)
/* return the first element with the highest frecuency */
func findMostRepeatedElement(frecuencyMap map[int]int) int {
	result, frecuencyFlag := int(0), int(0)

	for key, value := range frecuencyMap {
		if value >= frecuencyFlag {
			result, frecuencyFlag = key, value
		}
	}

	return result
}
