package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(S string) int {
	// convert the string to array of operations
	operations := strings.Split(S, " ")

	stack := make([]int, 0)

	const MaxUint = 1<<20 - 1 // maximum value of an integer in the stack. 1<<20 is a bit shift operation to get 2^20

	for _, op := range operations {
		// fmt.Println(op)
		switch op {
		case "DUP":
			if len(stack) == 0 {
				return -1
			}
			stack = append(stack, stack[len(stack)-1])
		case "POP":
			if len(stack) == 0 {
				return -1
			}
			stack = stack[:len(stack)-1]
		case "+":
			if len(stack) < 2 {
				return -1
			}
			first, second := stack[len(stack)-1], stack[len(stack)-2] // pop the last two elements
			stack = stack[:len(stack)-2]
			if first+second > MaxUint {
				return -1
			}
			stack = append(stack, first+second)
		case "-":
			if len(stack) < 2 {
				return -1
			}
			first, second := stack[len(stack)-1], stack[len(stack)-2] // pop the last two elements
			stack = stack[:len(stack)-2]
			if first < second {
				return -1
			}
			stack = append(stack, first-second)
		default:
			// check if the operation is a number
			num, err := strconv.Atoi(op)
			if err != nil {
				return -1
			}
			if num > MaxUint {
				return -1
			}
			stack = append(stack, num)
		}
	}
	if len(stack) == 0 { // empty stack
		return -1
	}
	return stack[len(stack)-1]
}

func main() {
	// provided test cases
	fmt.Println(Solution("13 DUP 4 POP 5 DUP + DUP + -"))
	fmt.Println(Solution("5 6 + -"))
	fmt.Println(Solution("3 DUP 5 - -"))
	fmt.Println(Solution("3 5 8 * +"))
	fmt.Println(Solution("3 5 8 * + 1000000000"))
}
