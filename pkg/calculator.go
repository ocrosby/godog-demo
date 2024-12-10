package pkg

import "errors"

// Calculator is a simple calculator that can add, subtract, multiply, and divide two integers.
type Calculator struct {
	Accumulator int
}

// NewCalculator creates a new Calculator.
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add adds two integers and stores the result in the accumulator.
func (c *Calculator) Add(a, b int) int {
	c.Accumulator = a + b
	return c.Accumulator
}

// Subtract subtracts two integers and stores the result in the accumulator.
func (c *Calculator) Subtract(a, b int) int {
	c.Accumulator = a - b
	return c.Accumulator
}

// Multiply multiplies two integers and stores the result in the accumulator.
func (c *Calculator) Multiply(a, b int) int {
	c.Accumulator = a * b
	return c.Accumulator
}

// Divide divides two integers and stores the result in the accumulator.
func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	c.Accumulator = a / b

	return c.Accumulator, nil
}

// GetAccumulator returns the current value of the accumulator.
func (c *Calculator) GetAccumulator() int {
	return c.Accumulator
}
