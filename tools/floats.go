package main

import (
	"fmt"
	"math"
	"strconv"
)

// simpleEvalFloat calculates the result of two floats and one of the
// following mathematical operators: "+,-,*,/,^".
func simpleEvalFloat(a, b float64, op string) float64 {
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
		return math.Pow(a, b)
	default:
		return 0.0
	}
}

// evalPostfixFloat evaluates a postfix (Reverse Polish Notation)
// expression as a string array of floats and operators.
func evalPostfixFloat(postfix []string) (float64, error) {
	var stack []float64
	for _, val := range postfix {
		switch {
		case isOp(val):
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			retVal := simpleEvalFloat(a, b, val)
			stack = append(stack[0:len(stack)-2], retVal)
		case isFunc(val):
			a := stack[len(stack)-1]
			retVal := FUNCS[val](a)
			stack = append(stack[0:len(stack)-1], retVal)
		case isConst(val):
			num := CONSTS[val]
			stack = append(stack, num)
		default:
			num, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return 0.0, err
			}
			stack = append(stack, num)
		}
	}
	if len(stack) != 1 {
		err := fmt.Errorf("error calculating postfix")
		return 0.0, err
	}
	return stack[0], nil
}
