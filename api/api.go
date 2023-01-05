package api

import (
	"Calculator/errors"
	"Calculator/math"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
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
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	api := router.Group("/api")
	{
		api.GET("/calculator/add", addHandler)
		api.GET("/calculator/subtract", subtractHandler)
		api.GET("/calculator/divide", divideHandler)
		api.GET("/calculator/multiply", multiplyHandler)
	}

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}

func addHandler(c *gin.Context) {
	calcRequest := getRequestData(c.Request)
	result := math.Add(calcRequest.FirstValue(), calcRequest.SecondValue())
	setResponse(c, result)
}

func subtractHandler(c *gin.Context) {
	calcRequest := getRequestData(c.Request)
	result := math.Subtract(calcRequest.FirstValue(), calcRequest.SecondValue())
	setResponse(c, result)
}

func divideHandler(c *gin.Context) {
	calcRequest := getRequestData(c.Request)
	result := math.Divide(calcRequest.FirstValue(), calcRequest.SecondValue())
	setResponse(c, result)
}

func multiplyHandler(c *gin.Context) {
	calcRequest := getRequestData(c.Request)
	result := math.Multiply(calcRequest.FirstValue(), calcRequest.SecondValue())
	setResponse(c, result)
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

func setResponse(c *gin.Context, responseData float64) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"response": responseData,
	})
}
