`automathic` implements Dijkstra's Shunting-Yard algorithm for parsing
mathematical infix expressions. The resultant postfix expression is
calculated and printed. Algebraic operations are supported by passing
variable definitions (run `automathic -h` for syntax).

Support is planned for functions in expressions, such as trig and most
other functions in Go's math package.

This tool is intended to facilitate tools for other algorithms for
calculating derivatives, integrals, etc. for mathematical formulas.

Just run `go build *.go` for now, proper building will come with time.
