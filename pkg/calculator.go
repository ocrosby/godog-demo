// Package pkg provides core domain types for the godog-demo project,
// including Calculator — a simple integer arithmetic engine used by the
// BDD calculator feature scenarios.
package pkg

import "errors"

// Calculator performs basic integer arithmetic and retains the result of the
// most recent operation in its Accumulator field, mirroring the state a
// real pocket calculator would display after each key-press.
type Calculator struct {
	// Accumulator holds the result of the last arithmetic operation.
	// It is zero-valued on construction and updated by every successful call
	// to Add, Subtract, Multiply, or Divide.
	Accumulator int
}

// NewCalculator returns a new Calculator with its Accumulator initialised to zero.
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Add sums a and b, stores the result in the Accumulator, and returns it.
func (c *Calculator) Add(a, b int) int {
	c.Accumulator = a + b
	return c.Accumulator
}

// Subtract computes a minus b, stores the result in the Accumulator, and returns it.
func (c *Calculator) Subtract(a, b int) int {
	c.Accumulator = a - b
	return c.Accumulator
}

// Multiply computes a times b, stores the result in the Accumulator, and returns it.
func (c *Calculator) Multiply(a, b int) int {
	c.Accumulator = a * b
	return c.Accumulator
}

// Divide computes a divided by b using integer (truncating) division, stores the
// result in the Accumulator, and returns it.
// It returns an error when b is zero; in that case the Accumulator is not modified.
func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	c.Accumulator = a / b

	return c.Accumulator, nil
}

// GetAccumulator returns the current value of the Accumulator without modifying it.
func (c *Calculator) GetAccumulator() int {
	return c.Accumulator
}
