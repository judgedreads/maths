package main

import (
	"fmt"
	"strings"
)

type Yard struct {
	Output []string
	Stack  []string
	Buf    []byte
}

func (y *Yard) flushBuf() {
	if len(y.Buf) == 0 {
		return
	}
	s := strings.ToLower(string(y.Buf))
	if isFunc(s) {
		y.Stack = append(y.Stack, s)
	} else {
		y.Output = append(y.Output, s)
	}
	y.Buf = y.Buf[:0]
	return
}

func (y *Yard) processOp(op string) {
	if len(y.Stack) == 0 {
		y.Stack = append(y.Stack, op)
		return
	}
	// loop backwards here as we need to pop from the back
	for i := len(y.Stack) - 1; i >= 0; i-- {
		p := precedence(op, y.Stack[i])
		// "^" is the only right-associative operator
		if p > 0 || (p == 0 && op == "^") {
			break
		} else {
			y.Output = append(y.Output, y.Stack[i])
			y.Stack = y.Stack[:i]
		}
	}
	y.Stack = append(y.Stack, op)
}

func (y *Yard) closeBracket() error {
	for i := len(y.Stack) - 1; i >= 0; i-- {
		if y.Stack[i] == "(" {
			y.Stack = y.Stack[:i]
			break
		} else {
			if len(y.Stack) == 1 {
				return fmt.Errorf("mismatched parentheses: expected \"(\"")
			}
			y.Output = append(y.Output, y.Stack[i])
			y.Stack = y.Stack[:i]
		}
	}
	return nil
}

// shuntingYard applies the Dijkstra algorithm of the same name to an
// expression in "infix" notation, returning an expression in postfix
// (Reverse Polish Notation).
func shuntingYard(infix string) ([]string, error) {
	yard := Yard{}

	// make sure there are no leading/trailing operators
	first := string(infix[0])
	if isOp(first) {
		if first == "+" || first == "-" {
			infix = "0" + infix
		} else {
			err := fmt.Errorf("cannot start expression with %q", first)
			return yard.Output, err
		}
	}
	last := string(infix[len(infix)-1])
	if isOp(last) {
		err := fmt.Errorf("cannot end expression with %q", last)
		return yard.Output, err
	}

	for i := 0; i < len(infix); i++ {
		t := string(infix[i])
		switch {
		case t == " ":
			continue
		case isOp(t):
			yard.flushBuf()
			yard.processOp(t)
		case t == "(":
			yard.flushBuf()
			yard.Stack = append(yard.Stack, t)
		case t == ")":
			yard.flushBuf()
			err := yard.closeBracket()
			if err != nil {
				return yard.Output, err
			}
		default:
			yard.Buf = append(yard.Buf, infix[i])
		}
	}
	yard.flushBuf()
	for i := len(yard.Stack) - 1; i >= 0; i-- {
		yard.Output = append(yard.Output, yard.Stack[i])
	}
	return yard.Output, nil
}
