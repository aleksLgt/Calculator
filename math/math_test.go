package math

import "testing"

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
