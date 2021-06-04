package controller

// Arithmetic is a struct for arithmetic operations.
type Arithmetic struct{}

// NewArithmetic is a factory for Arithmetic.
func NewArithmetic() *Arithmetic {
	return &Arithmetic{}
}

// Add adds two numbers.
func (c Arithmetic) Add(num1, num2 int) int {
	return num1 + num2
}

// Subtract subtract two numbers.
func (c Arithmetic) Subtract(num1, num2 int) int {
	return num1 - num2
}
