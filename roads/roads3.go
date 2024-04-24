package main

import "fmt"

func minReorder3(n int, connections [][]int) int {

	// create a adjascency list as map between connections
	adjList := make(map[int][]int)

	for _, conn := range connections {
		adjList[conn[0]] = append(adjList[conn[0]], conn[1])
		adjList[conn[1]] = append(adjList[conn[1]], -conn[0])
	}

	reorientations := 0
	travelled := make(map[int]bool)

	var dfs func(node int)

	dfs = func(node int) {
		fmt.Println("node:", node)
		// start travelling from 0 node to adjascents and if reorientations required then add no of reorientations
		travelled[node] = true
		neighbours := adjList[node]
		for _, neighbour := range neighbours {
			if !travelled[neighbour] {
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

// func absValue(num int) int {
// 	if num < 0 {
// 		return -num
// 	}
// 	return num

// }

// func main() {
// 	fmt.Printf("\nreorientations: %d\n", minReorder3(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
// }
