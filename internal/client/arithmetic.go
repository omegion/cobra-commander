package client

import (
	"github.com/omegion/go-cli/internal/controller"
)

// ArithmeticInterface is an interface for Client.
type ArithmeticInterface interface {
	AddNumbers(options NumbersOptions) int
	SubtractNumbers(options NumbersOptions) int
}

// NumbersOptions is options for Arithmetic.
type NumbersOptions struct {
	Number1,
	Number2 int
}

// AddNumbers adds two numbers.
func (c Client) AddNumbers(options NumbersOptions) int {
	return controller.NewArithmetic().Add(options.Number1, options.Number2)
}

// SubtractNumbers subtracts two numbers.
func (c Client) SubtractNumbers(options NumbersOptions) int {
	return controller.NewArithmetic().Subtract(options.Number1, options.Number2)
}
