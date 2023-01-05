package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(10.6, 23.9) != 34.5 {
		t.Error("10.6 + 23.9 did not equal 34.5")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(25.3, 11.4) != 13.9 {
		t.Error("25.3 - 11.4 did not equal 13.9")
	}
}

func TestDivide(t *testing.T) {
	if Divide(10.23, 0.2) != 51.15 {
		t.Error("10.23 / 2 did not equal 51.15")
	}
}

func TestDivideByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	if Divide(10.23, 0) != 51.15 {
		t.Error("10.23 / 2 did not equal 51.15")
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(10.23, 0.2) != 2.046 {
		t.Error("10.23 * 0.2 did not equal 2.046")
	}
}
