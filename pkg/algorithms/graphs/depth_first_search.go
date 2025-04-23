package graphs

func DFS(adjacency_matrix [][]int, visited []bool, node int) {
	visited[node] = true

	for neighbor := 0; neighbor < len(adjacency_matrix); neighbor++ {
		if adjacency_matrix[node][neighbor] == 1 && !visited[neighbor] {
			DFS(adjacency_matrix, visited, neighbor)
		}
	}
}
