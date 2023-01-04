package api

import (
	"Calculator/errors"
	"Calculator/math"
	"fmt"
	"net/http"
	"strconv"
)

type RequestData struct {
	firstValue  float64
	secondValue float64
}

func (c *RequestData) SetFirstValue(number float64) {
	c.firstValue = number
}

func (c *RequestData) SetSecondValue(number float64) {
	c.secondValue = number
}

func (c *RequestData) FirstValue() float64 {
	return c.firstValue
}

func (c *RequestData) SecondValue() float64 {
	return c.secondValue
}

func InitializeEndpoints() {
	http.HandleFunc("/calculator/add", addHandler)
	http.HandleFunc("/calculator/subtract", subtractHandler)
	http.HandleFunc("/calculator/divide", divideHandler)
	http.HandleFunc("/calculator/multiply", multiplyHandler)
}

func addHandler(writer http.ResponseWriter, request *http.Request) {
	calcRequest := getRequestData(request)
	result := math.Add(calcRequest.FirstValue(), calcRequest.SecondValue())
	setWriter(writer, result)
}

func subtractHandler(writer http.ResponseWriter, request *http.Request) {
	calcRequest := getRequestData(request)
	result := math.Subtract(calcRequest.FirstValue(), calcRequest.SecondValue())
	setWriter(writer, result)
}

func divideHandler(writer http.ResponseWriter, request *http.Request) {
	calcRequest := getRequestData(request)
	result := math.Divide(calcRequest.FirstValue(), calcRequest.SecondValue())
	setWriter(writer, result)
}

func multiplyHandler(writer http.ResponseWriter, request *http.Request) {
	calcRequest := getRequestData(request)
	result := math.Multiply(calcRequest.FirstValue(), calcRequest.SecondValue())
	setWriter(writer, result)
}

func setWriter(writer http.ResponseWriter, result float64) {
	stringValue := fmt.Sprint(result)
	_, err := writer.Write([]byte(stringValue))
	errors.Check(err)
}

func getRequestData(request *http.Request) RequestData {
	firstValue, err := strconv.ParseFloat(request.FormValue("firstValue"), 64)
	errors.Check(err)
	secondValue, err := strconv.ParseFloat(request.FormValue("secondValue"), 64)
	errors.Check(err)
	calcRequest := RequestData{}
	calcRequest.SetFirstValue(firstValue)
	calcRequest.SetSecondValue(secondValue)
	return calcRequest
}
