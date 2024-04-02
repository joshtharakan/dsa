package move

import (
	"fmt"
)

func Solution(S string) int {
	length := len(S)
	moves := 0
	occupied := make([]bool, length)
	for i := range occupied {
		occupied[i] = true
	}

	for i, move := range S {
		switch move {
		case '>':
			if i+1 == length {
				moves++
				occupied[i] = false
				continue
			}
			if !occupied[i+1] {
				moves++
				occupied[i+1] = true
				occupied[i] = false
			}
		case '<':
			if i-1 < 0 {
				moves++
				occupied[i] = false
				continue
			} else if !occupied[i-1] {
				moves++
				occupied[i-1] = true
				occupied[i] = false
			}
		case '^':
			fmt.Printf("i: %d, char: %s\n", i, string(move))
			moves++
			occupied[i] = false
		case 'v':
			fmt.Printf("i: %d, char: %s\n", i, string(move))
			moves++
			occupied[i] = false
		}

	}
	return moves
}
