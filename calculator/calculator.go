package calculator

// ArithmaticOperation struct
type ArithmaticOperation struct {
	OperationType string
	A, B          float64
}

// Sum of a and b
func Sum(a float64, b float64) float64 {
	return a + b
}

// Subtract of a and b
func Subtract(a float64, b float64) float64 {
	return a - b
}

// Multiply a by b
func Multiply(a float64, b float64) float64 {
	return a * b
}

// Divide a by b
func Divide(a float64, b float64) float64 {
	return a / b
}
