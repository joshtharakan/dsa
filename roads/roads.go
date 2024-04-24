package main

import "fmt"

type TreeNode struct {
	city       int
	neighbours []*TreeNode
}

func CreateTreeNode(city int) *TreeNode {
	return &TreeNode{
		city:       city,
		neighbours: make([]*TreeNode, 0),
	}
}

func solution(A []int, B []int) int {

	n := len(A) + 1
	graph := make([]*TreeNode, n)

	for i := 0; i < n; i++ {
		graph[i] = CreateTreeNode(i)
	}

	for i := 0; i < len(A); i++ {
		graph[A[i]].neighbours = append(graph[A[i]].neighbours, graph[B[i]])
		graph[B[i]].neighbours = append(graph[B[i]].neighbours, graph[A[i]])
	}

	// print graph
	for i := 0; i < n; i++ {
		fmt.Printf("City %d: ", i)
		for j := 0; j < len(graph[i].neighbours); j++ {
			fmt.Printf("%d ", graph[i].neighbours[j].city)
		}
		fmt.Println()
		// fmt.Println(graph[i].neighbours)
	}

	visited := make([]bool, n)
	orientations := 0

	var dfs func(*TreeNode)
	dfs = func(currentNode *TreeNode) {
		visited[currentNode.city] = true
		for _, neighbour := range currentNode.neighbours {
			if !visited[neighbour.city] {
				correctlyOriented := false
				for j := 0; j < len(A); j++ {
					if B[j] == currentNode.city && A[j] == neighbour.city {
						correctlyOriented = true
						break
					}
				}
				if !correctlyOriented {
					orientations++
				}
				dfs(neighbour)
			}
		}

	}

	dfs(graph[0])

	return orientations
}

// func main() {
// 	A := []int{1, 1, 3, 3}
// 	B := []int{0, 2, 2, 4}
// 	fmt.Println(solution(A, B))
// 	fmt.Println(solution([]int{0, 2, 2, 3}, []int{1, 1, 4, 4}))             // Should return 2
// 	fmt.Println(solution([]int{1, 6, 6, 3, 0, 5}, []int{6, 2, 0, 0, 4, 0})) // Should return 2
// 	fmt.Println(solution([]int{0, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}))       // Should return 5
// }
