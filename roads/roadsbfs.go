package main

import "fmt"

func minReorderBFS(n int, connections [][]int) int {

	// create adjascent list as a matrix
	graph := make([][]int, n)
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], -conn[0])
	}

	// create a queue from the above list starting from 0
	queue := []int{0}

	reorientations := 0
	visited := make([]bool, n)
	visited[0] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited[node] = true

		for _, neighbour := range graph[node] {
			if !visited[absValue(neighbour)] {
				if neighbour > 0 {
					reorientations++
				}
				queue = append(queue, absValue(neighbour))
			}
		}
	}

	return reorientations
}

// absValue returns the absolute value of a number
func absValue(num int) int {
	if num >= 0 {
		return num
	}
	return -num

}

func main() {
	fmt.Printf("\nreorientations: %d\n", minReorderBFS(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
}
