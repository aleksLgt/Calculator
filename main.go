package main

import (
	"Calculator/math"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type CalcRequest struct {
	firstValue  float64
	secondValue float64
}

func (c *CalcRequest) SetFirstValue(number float64) {
	c.firstValue = number
}

func (c *CalcRequest) SetSecondValue(number float64) {
	c.secondValue = number
}

func (c *CalcRequest) FirstValue() float64 {
	return c.firstValue
}

func (c *CalcRequest) SecondValue() float64 {
	return c.secondValue
}

func main() {
	http.HandleFunc("/calculator/add", addHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func addHandler(writer http.ResponseWriter, request *http.Request) {
	calcRequest := getCalcRequest(request)
	result := fmt.Sprint(math.Add(calcRequest.FirstValue(), calcRequest.SecondValue()))
	_, err := writer.Write([]byte(result))
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getCalcRequest(request *http.Request) CalcRequest {
	firstValue, err := strconv.ParseFloat(request.FormValue("firstValue"), 64)
	check(err)
	secondValue, err := strconv.ParseFloat(request.FormValue("secondValue"), 64)
	check(err)
	calcRequest := CalcRequest{}
	calcRequest.SetFirstValue(firstValue)
	calcRequest.SetSecondValue(secondValue)
	return calcRequest
}
