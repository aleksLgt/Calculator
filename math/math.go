package math

import (
	"Calculator/errors"
	"fmt"
	"strconv"
)

func Add(a float64, b float64) float64 {
	return a + b
}

func Subtract(a float64, b float64) float64 {
	return a - b
}

func Divide(a float64, b float64) float64 {
	result := a / b
	response, err := strconv.ParseFloat(fmt.Sprintf("%.6f", result), 64)
	errors.Check(err)
	return response
}

func Multiply(a float64, b float64) float64 {
	result := a * b
	response, err := strconv.ParseFloat(fmt.Sprintf("%.6f", result), 64)
	errors.Check(err)
	return response
}
