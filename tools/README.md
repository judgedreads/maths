##About
`automathic` implements Dijkstra's Shunting-Yard algorithm for parsing
mathematical infix expressions. The resultant postfix expression is
calculated and printed. Algebraic operations are supported by passing
variable definitions (run `automathic -h` for syntax).

##Functions and constants
The following functions and constants are supported (case insensitive).
###Constants
- pi
- e

###Functions
- sqrt    (math.Sqrt)
- cbrt    (math.Cbrt)
- abs     (math.Abs)
- sin     (math.Sin)
- cos     (math.Cos)
- tan     (math.Tan)
- arcsin  (math.Asin)
- arccos  (math.Acos)
- arctan  (math.Atan)
- sinh    (math.Sinh)
- cosh    (math.Cosh)
- tanh    (math.Tanh)
- arcsinh (math.Asinh)
- arccosh (math.Acosh)
- arctanh (math.Atanh)
- ln      (math.Log)
- log     (math.Log)
- log10   (math.Log10)
- log2    (math.Log2)

##Building
`go build -o automathic`
