package main

func minReorder4(n int, connections [][]int) int {

	// create adjascent list as a matrix
	graph := make([][]int, n)
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], -conn[0])
	}
	reorientations := 0
	visited := make([]bool, n)

	var dfs func(node int)

	// dfs function to traverse the graph
	dfs = func(node int) {
		visited[node] = true
		for _, neighbour := range graph[node] {
			if !visited[absValue(neighbour)] {
				dfs(absValue(neighbour))
				if neighbour > 0 {
					reorientations++
				}
			}
		}
	}
	dfs(0)
	return reorientations
}

// // absValue returns the absolute value of a number
// func absValue(num int) int {
// 	if num >= 0 {
// 		return num
// 	}
// 	return -num

// }

// func main() {
// 	fmt.Printf("\nreorientations: %d\n", minReorder4(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
// }
