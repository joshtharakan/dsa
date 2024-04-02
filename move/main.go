package main

import (
	"fmt"
)

func Solution1(S string) int {
	length := len(S)
	moves := 0
	notOccupied := make([]bool, length)

	for i, move := range S {
		target := i
		switch move {
		case '>':
			target = i + 1
		case '<':
			target = i - 1
		case '^', 'v':
			moves++
			notOccupied[i] = true
			continue
		}

		if target < 0 || target >= length {
			moves++
			notOccupied[i] = true
			continue
		} else if notOccupied[target] {
			moves++
			notOccupied[i] = true
			continue
		}
	}
	return moves
}

func main() {
	fmt.Println(Solution1("><^v"))
	fmt.Println(Solution1("<<^<v>>"))
}
