package math

import (
	"Calculator/errors"
	"fmt"
	"strconv"
)

func Add(a float64, b float64) float64 {
	result := a + b
	return prepareResponse(result)
}

func Subtract(a float64, b float64) float64 {
	result := a - b
	return prepareResponse(result)

}

func Divide(a float64, b float64) float64 {
	if b == 0 {
		panic(fmt.Errorf("cannot be divided by zero"))
	}
	result := a / b
	return prepareResponse(result)

}

func Multiply(a float64, b float64) float64 {
	result := a * b
	return prepareResponse(result)
}

func prepareResponse(result float64) float64 {
	response, err := strconv.ParseFloat(fmt.Sprintf("%.9f", result), 64)
	errors.Check(err)
	return response
}
