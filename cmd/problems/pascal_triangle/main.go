package main

import "fmt"

func main() {
	input := 5

	rows := generate(input)

	for _, row := range rows {
		fmt.Println(row)
	}
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		// set values of edges
		row[0], row[i] = 1, 1

		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}

		result[i] = row
	}

	return result
}
