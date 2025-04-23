package main

import "fmt"

func main() {

	adjacencyMatrix := [][]int{{1, 1, 1, 0, 0, 0}, {1, 1, 1, 0, 0, 0},
		{1, 1, 1, 0, 0, 0}, {0, 0, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 0}, {0, 0, 0, 1, 0, 1}}
	numberOfProvinces := findCircleNum(adjacencyMatrix)
	fmt.Println(numberOfProvinces)
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visitedArr := make([]bool, n)
	numOfProvinces := 0

	// time: O(n)
	for i := 0; i < n; i++ {
		if !visitedArr[i] {
			numOfProvinces++
			dfs(isConnected, visitedArr, i)
		}
	}

	return numOfProvinces
}

// time: O(n*2)
func dfs(adjacencyMatrix [][]int, visitedArr []bool, node int) {
	n := len(adjacencyMatrix)
	// set node as visited
	visitedArr[node] = true

	for i := 0; i < n; i++ {
		if adjacencyMatrix[node][i] == 1 && !visitedArr[i] {
			dfs(adjacencyMatrix, visitedArr, i)
		}
	}
}
