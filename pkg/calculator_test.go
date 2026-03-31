package pkg

import (
	"math"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	c := NewCalculator()
	if c == nil {
		t.Fatal("NewCalculator returned nil")
	}
	if c.Accumulator != 0 {
		t.Errorf("expected Accumulator = 0, got %d", c.Accumulator)
	}
}

func TestCalculator_Add(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive + positive", 2, 3, 5},
		{"positive + zero", 5, 0, 5},
		{"zero + positive", 0, 5, 5},
		{"zero + zero", 0, 0, 0},
		{"negative + negative", -3, -4, -7},
		{"positive + negative", 10, -3, 7},
		{"negative + positive", -3, 10, 7},
		{"result is negative", -10, 3, -7},
		{"max int overflow (silent wraparound)", math.MaxInt, 1, math.MinInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCalculator()
			got := c.Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
			if c.GetAccumulator() != tt.want {
				t.Errorf("Accumulator = %d after Add(%d, %d), want %d", c.GetAccumulator(), tt.a, tt.b, tt.want)
			}
		})
	}
}

func TestCalculator_Subtract(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive - positive, positive result", 5, 3, 2},
		{"positive - positive, zero result", 5, 5, 0},
		{"positive - positive, negative result", 3, 5, -2},
		{"zero - zero", 0, 0, 0},
		{"zero - positive", 0, 5, -5},
		{"positive - zero", 5, 0, 5},
		{"negative - negative", -3, -5, 2},
		{"negative - positive", -3, 5, -8},
		{"positive - negative", 3, -5, 8},
		{"min int underflow (silent wraparound)", math.MinInt, 1, math.MaxInt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCalculator()
			got := c.Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Subtract(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
			if c.GetAccumulator() != tt.want {
				t.Errorf("Accumulator = %d after Subtract(%d, %d), want %d", c.GetAccumulator(), tt.a, tt.b, tt.want)
			}
		})
	}
}

func TestCalculator_Multiply(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive * positive", 2, 3, 6},
		{"positive * zero", 5, 0, 0},
		{"zero * positive", 0, 5, 0},
		{"zero * zero", 0, 0, 0},
		{"positive * one", 7, 1, 7},
		{"negative * negative", -3, -4, 12},
		{"positive * negative", 3, -4, -12},
		{"negative * positive", -3, 4, -12},
		{"positive * negative one", 5, -1, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCalculator()
			got := c.Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
			if c.GetAccumulator() != tt.want {
				t.Errorf("Accumulator = %d after Multiply(%d, %d), want %d", c.GetAccumulator(), tt.a, tt.b, tt.want)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{"positive / positive, exact", 6, 2, 3, false},
		{"positive / positive, truncated", 7, 2, 3, false},
		{"zero / positive", 0, 5, 0, false},
		{"negative / negative", -6, -2, 3, false},
		{"positive / negative", 6, -2, -3, false},
		{"negative / positive", -6, 2, -3, false},
		{"negative truncation toward zero", -7, 2, -3, false},
		{"positive / one", 5, 1, 5, false},
		{"divide by zero", 6, 0, 0, true},
		{"zero divide by zero", 0, 0, 0, true},
		{"negative divide by zero", -6, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCalculator()
			got, err := c.Divide(tt.a, tt.b)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error, got nil", tt.a, tt.b)
				}
				return
			}

			if err != nil {
				t.Fatalf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Errorf("Divide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
			if c.GetAccumulator() != tt.want {
				t.Errorf("Accumulator = %d after Divide(%d, %d), want %d", c.GetAccumulator(), tt.a, tt.b, tt.want)
			}
		})
	}
}

func TestCalculator_GetAccumulator(t *testing.T) {
	t.Run("initial value is zero", func(t *testing.T) {
		c := NewCalculator()
		if c.GetAccumulator() != 0 {
			t.Errorf("expected 0, got %d", c.GetAccumulator())
		}
	})

	t.Run("reflects last operation result", func(t *testing.T) {
		c := NewCalculator()
		c.Add(3, 4)
		if c.GetAccumulator() != 7 {
			t.Errorf("expected 7, got %d", c.GetAccumulator())
		}
		c.Subtract(10, 1)
		if c.GetAccumulator() != 9 {
			t.Errorf("expected 9, got %d", c.GetAccumulator())
		}
	})
}

func TestCalculator_AccumulatorUnchangedOnDivideByZero(t *testing.T) {
	// A failed Divide must not modify the accumulator — it returns early before the assignment.
	c := NewCalculator()
	c.Add(10, 5)
	if c.GetAccumulator() != 15 {
		t.Fatalf("setup: expected accumulator = 15, got %d", c.GetAccumulator())
	}

	_, err := c.Divide(10, 0)
	if err == nil {
		t.Fatal("expected divide-by-zero error")
	}
	if c.GetAccumulator() != 15 {
		t.Errorf("accumulator = %d after divide-by-zero, want 15 (unchanged)", c.GetAccumulator())
	}
}
