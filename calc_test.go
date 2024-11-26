package calc

import (
	"fmt"
	"testing"
)

func TestAddition_Calculate(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{a: 1, b: 2, c: 3},
		{a: 0, b: 0, c: 0},
		{a: 0, b: 1, c: 1},
		{a: 1, b: 0, c: 1},
		{a: 1, b: 1, c: 2},
		{a: 12345, b: 67890, c: 80235},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d+%d=%d", tt.a, tt.b, tt.c), func(t *testing.T) {
			this := &Addition{}
			if got := this.Calculate(tt.a, tt.b); got != tt.c {
				t.Errorf("Calculate() = %v, want %v", got, tt.c)
			}
		})
	}
}
func TestSubtraction_Calculate(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{a: 1, b: 2, c: -1},
		{a: 0, b: 0, c: 0},
		{a: 0, b: 1, c: -1},
		{a: 1, b: 0, c: 1},
		{a: 1, b: 1, c: 0},
		{a: 12345, b: 67890, c: -55_545},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%d=%d", tt.a, tt.b, tt.c), func(t *testing.T) {
			this := &Subtraction{}
			if got := this.Calculate(tt.a, tt.b); got != tt.c {
				t.Errorf("Calculate() = %v, want %v", got, tt.c)
			}
		})
	}
}
func TestMultiplication_Calculate(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{a: 1, b: 2, c: 2},
		{a: 0, b: 0, c: 0},
		{a: 0, b: 1, c: 0},
		{a: 1, b: 0, c: 0},
		{a: 1, b: 1, c: 1},
		{a: 12345, b: 67890, c: 838_102_050},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d*%d=%d", tt.a, tt.b, tt.c), func(t *testing.T) {
			this := &Multiplication{}
			if got := this.Calculate(tt.a, tt.b); got != tt.c {
				t.Errorf("Calculate() = %v, want %v", got, tt.c)
			}
		})
	}
}
func TestDivision_Calculate(t *testing.T) {
	tests := []struct {
		a, b, c int
	}{
		{a: 4, b: 2, c: 2},
		{a: 0, b: 1, c: 0},
		{a: 10, b: 3, c: 3},
		{a: 654, b: 321, c: 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d/%d=%d", tt.a, tt.b, tt.c), func(t *testing.T) {
			this := &Division{}
			if got := this.Calculate(tt.a, tt.b); got != tt.c {
				t.Errorf("Calculate() = %v, want %v", got, tt.c)
			}
		})
	}
}
