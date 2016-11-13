package main

import (
	"fmt"
	"math"
	"strconv"
)

// simpleEvalFloat calculates the result of two floats and one of the
// following mathematical operators: "+,-,*,/,^".
func simpleEvalFloat(a, b float64, op string) float64 {
	var ans float64
	if op == "+" {
		ans = a + b
	} else if op == "-" {
		ans = a - b
	} else if op == "/" {
		ans = a / b
	} else if op == "*" {
		ans = a * b
	} else if op == "^" {
		ans = math.Pow(a, b)
	}
	return ans
}

func funcEvalFloat(a float64, f string) float64 {
	var ans float64
	if f == "sin" {
		ans = math.Sin(a)
	} else if f == "cos" {
		ans = math.Cos(a)
	} else if f == "tan" {
		ans = math.Tan(a)
	}
	return ans
}

// evalPostfixFloat evaluates a postfix (Reverse Polish Notation)
// expression as a string array of floats and operators.
func evalPostfixFloat(postfix []string) (float64, error) {
	var stack []float64
	for _, val := range postfix {
		if isOp(val) {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			retVal := simpleEvalFloat(a, b, val)
			stack = append(stack[0:len(stack)-2], retVal)
		} else if isFunc(val) {
			a := stack[len(stack)-1]
			retVal := funcEvalFloat(a, val)
			stack = append(stack[0:len(stack)-1], retVal)
		} else {
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
