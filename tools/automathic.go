package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

// map contant keywords to library constants
var CONSTS = map[string]float64{
	"pi": math.Pi,
	"e":  math.E,
}

// map function keywords to library calls
var FUNCS = map[string]func(float64) float64{
	"sqrt":    math.Sqrt,
	"cbrt":    math.Cbrt,
	"abs":     math.Abs,
	"sin":     math.Sin,
	"cos":     math.Cos,
	"tan":     math.Tan,
	"arcsin":  math.Asin,
	"arccos":  math.Acos,
	"arctan":  math.Atan,
	"sinh":    math.Sinh,
	"cosh":    math.Cosh,
	"tanh":    math.Tanh,
	"arcsinh": math.Asinh,
	"arccosh": math.Acosh,
	"arctanh": math.Atanh,
	"ln":      math.Log,
	"log":     math.Log,
	"log10":   math.Log10,
	"log2":    math.Log2,
}

// map operators to their precedence
var OPS = map[string]int{
	"+": 2,
	"-": 2,
	"*": 3,
	"/": 3,
	"^": 4,
}

func main() {
	vars := flag.String("vars", "", "comma separated list of var definitions, e.g. x=1,y=2")
	integers := flag.Bool("ints", false, "use integers, not floats, in calculations")
	flag.Parse()
	expr := flag.Arg(0)
	if expr == "" {
		err := fmt.Errorf("need an expression to evaluate")
		die(err)
	}
	if *vars != "" {
		v := strings.Split(*vars, ",")
		expr, err := subVars(expr, v)
		if err != nil {
			die(err)
		}
	}
	shunt, err := shuntingYard(expr)
	if err != nil {
		die(err)
	}
	if *integers {
		ans, err := evalPostfixInt(shunt)
		if err != nil {
			die(err)
		}
		fmt.Printf("%s = \n\t%d\n", expr, ans)
	} else {
		ans, err := evalPostfixFloat(shunt)
		if err != nil {
			die(err)
		}
		fmt.Printf("%s = \n\t%f\n", expr, ans)
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

// precedence compares the precedence of the mathematical operators:
// "+,-,*,/,^"
func precedence(o1, o2 string) int {
	return OPS[o1] - OPS[o2]
}

func isOp(s string) bool {
	_, exists := OPS[s]
	return exists
}

func isFunc(s string) bool {
	_, exists := FUNCS[s]
	return exists
}

func isConst(s string) bool {
	_, exists := CONSTS[s]
	return exists
}
