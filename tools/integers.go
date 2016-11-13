package main

import (
	"fmt"
	"strconv"
)

// simpleEvalInt calculates the result of two ints and one of the
// following mathematical operators: "+,-,*,/,^".
func simpleEvalInt(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "*":
		return a * b
	case "^":
		return power(a, b)
	default:
		return 0
	}
}

// evalPostfixInt evaluates a postfix (Reverse Polish Notation)
// expression as a string array of ints and operators.
func evalPostfixInt(postfix []string) (int, error) {
	var stack []int
	for _, val := range postfix {
		if isOp(val) {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			retVal := simpleEvalInt(a, b, val)
			stack = append(stack[0:len(stack)-2], retVal)
		} else {
			num, err := strconv.Atoi(val)
			if err != nil {
				return 0, nil
			}
			stack = append(stack, num)
		}
	}
	if len(stack) != 1 {
		err := fmt.Errorf("error calculating postfix")
		return 0, err
	}
	return stack[0], nil
}

// power raises a to the power of b
func power(a, b int) int {
	ans := 1
	for i := 0; i < b; i++ {
		ans = a * ans
	}
	return ans
}
