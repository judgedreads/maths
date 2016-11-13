package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// TODO: allow functions in expression
func main() {
	expr := flag.String("expr", "", "expression to evaluate")
	vars := flag.String("vars", "", "comma separated list of var definitions, e.g. x=1,y=2")
	integers := flag.Bool("ints", false, "use integers, not floats, in calculations")
	flag.Parse()
	if *expr == "" {
		err := fmt.Errorf("need an expression to evaluate")
		die(err)
	}
	if *vars != "" {
		v := strings.Split(*vars, ",")
		subbed, err := subVars(*expr, v)
		expr = &subbed
		if err != nil {
			die(err)
		}
	}
	shunt, err := shuntingYard(*expr)
	if err != nil {
		die(err)
	}
	if *integers {
		ans, err := evalPostfixInt(shunt)
		if err != nil {
			die(err)
		}
		fmt.Printf("%s = \n\t%d\n", *expr, ans)
	} else {
		ans, err := evalPostfixFloat(shunt)
		if err != nil {
			die(err)
		}
		fmt.Printf("%s = \n\t%f\n", *expr, ans)
	}
}

// die simply prints the given error to stderr and exits unsuccessfully
func die(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

// subVars takes an array of pairs of the form "var=val" and returns a
// new string equal to expr with all instances of each var set to the
// corresponding val.
func subVars(expr string, vars []string) (string, error) {
	for i := 0; i < len(vars); i++ {
		subs := strings.Split(vars[i], "=")
		if len(subs) != 2 {
			err := fmt.Errorf("invalid var definition: %q", vars[i])
			return "", err
		}
		expr = strings.Replace(expr, subs[0], subs[1], -1)
	}
	return expr, nil
}

// TODO: using pointers seems bad, maybe a struct would be good here?
// Could group the output, stack, and buf onto one struct?
func flushBuf(b []byte, output, stack *[]string) []byte {
	if len(b) == 0 {
		return b
	}
	s := string(b)
	if isFunc(s) {
		*stack = append(*stack, s)
	} else {
		*output = append(*output, s)
	}
	return b[:0]
}

// shuntingYard applies the Dijkstra algorithm of the same name to an
// expression in "infix" notation, returning an expression in postfix
// (Reverse Polish Notation).
func shuntingYard(infix string) ([]string, error) {
	var output []string
	var stack []string
	var buf []byte

	// make sure there are no trailing operators
	first := string(infix[0])
	if isOp(first) {
		if first == "+" || first == "-" {
			infix = "0" + infix
		} else {
			err := fmt.Errorf("cannot start expression with %q", first)
			return output, err
		}
	}
	last := string(infix[len(infix)-1])
	if isOp(last) {
		err := fmt.Errorf("cannot end expression with %q", last)
		return output, err
	}

	// parse infix expression
	for i := 0; i < len(infix); i++ {
		t := string(infix[i])
		if isOp(t) {
			buf = flushBuf(buf, &output, &stack)
			if len(stack) == 0 {
				stack = append(stack, t)
			} else {
				// loop backwards here as we need to pop
				// from the back
				for i := len(stack) - 1; i >= 0; i-- {
					p := precedence(t, stack[i])
					// "^" is the only
					// right-associative operator
					if p > 0 || (p == 0 && t == "^") {
						break
					} else {
						output = append(output, stack[i])
						stack = stack[:i]
					}
				}
				stack = append(stack, t)
			}
		} else if t == "(" {
			buf = flushBuf(buf, &output, &stack)
			stack = append(stack, t)
		} else if t == ")" {
			buf = flushBuf(buf, &output, &stack)
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i] == "(" {
					stack = stack[:i]
					break
				} else {
					if len(stack) == 1 {
						err := fmt.Errorf("mismatched parentheses: expected \"(\"")
						return output, err
					}
					output = append(output, stack[i])
					stack = stack[:i]
				}
			}
		} else {
			buf = append(buf, infix[i])
		}
	}
	buf = flushBuf(buf, &output, &stack)
	for i := len(stack) - 1; i >= 0; i-- {
		output = append(output, stack[i])
	}
	return output, nil
}

// precedence compares the precedence of the mathematical operators:
// "+,-,*,/,^"
func precedence(o1, o2 string) int {
	m := make(map[string]int)
	m["+"] = 2
	m["-"] = 2
	m["*"] = 3
	m["/"] = 3
	m["^"] = 4
	return m[o1] - m[o2]
}

// isOp checks to see if a string is one the mathematical operators:
// "+,-,*,/,^"
func isOp(s string) bool {
	operators := []string{"*", "/", "+", "-", "^"}
	for _, op := range operators {
		if s == op {
			return true
		}
	}
	return false
}

func isFunc(s string) bool {
	funcs := []string{"sin", "cos", "tan"}
	for _, f := range funcs {
		if s == f {
			return true
		}
	}
	return false
}
