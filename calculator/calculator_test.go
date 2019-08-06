package calculator

import (
	"testing"
)

func TestSum(t *testing.T) {
	want := 43.41 + 33.52
	if got := Sum(43.41, 33.52); got != want {
		t.Errorf("Sum(43.41, 33.52) = %f want %f", got, want)
	}
}

func TestSubtract(t *testing.T) {
	want := 43.41 - 33.52
	if got := Subtract(43.41, 33.52); got != want {
		t.Errorf("Subtract(43.41, 33.52) = %f want %f", got, want)
	}
}
func TestMultiply(t *testing.T) {
	want := 43.41 * 33.52
	if got := Multiply(43.41, 33.52); got != want {
		t.Errorf("Multiply(43.41, 33.52) = %f want %f", got, want)
	}
}
func TestDivide(t *testing.T) {
	want := 43.41 / 33.52
	if got := Divide(43.41, 33.52); got != want {
		t.Errorf("Divide(43.41, 33.52) = %f want %f", got, want)
	}
}
