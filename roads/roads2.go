package main

import (
	"fmt"
)

type Node struct {
	city       int
	neighbours []*Node
}

func newNode(city int) *Node {
	return &Node{
		city:       city,
		neighbours: make([]*Node, 0),
	}
}

func minReorder(n int, connections [][]int) int {

	graph := make([]*Node, n)
	// create the graph
	for i := 0; i < n; i++ {
		graph[i] = newNode(i)
	}

	// populate the connections
	for _, pair := range connections {
		graph[pair[0]].neighbours = append(graph[pair[0]].neighbours, graph[pair[1]])
		graph[pair[1]].neighbours = append(graph[pair[1]].neighbours, graph[pair[0]])
	}

	for _, node := range graph {
		fmt.Printf("City %d: ", node.city)
		for _, neighbour := range node.neighbours {
			fmt.Printf("%d ", neighbour.city)
		}
		fmt.Println()
	}

	visited := make([]bool, n)
	orientations := 0

	var dfs func(*Node)

	dfs = func(currNode *Node) {
		// mark visited for the current node
		visited[currNode.city] = true

		// go through all the neighbours and check if they are visited already
		for _, neighbour := range currNode.neighbours {
			if !visited[neighbour.city] {
				dfs(neighbour)
				correctlyOriented := false
				for _, conn := range connections {
					// check if the direction is from current to neighbour
					if conn[0] == neighbour.city && conn[1] == currNode.city {
						correctlyOriented = true
						break
					}
				}
				if !correctlyOriented {
					orientations++
				}
			}
		}
	}

	dfs(graph[0])
	return orientations
}

// func main() {
// 	fmt.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
// }
