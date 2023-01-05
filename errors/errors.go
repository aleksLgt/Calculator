package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReportPanic(c *gin.Context) {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"response": err.Error(),
		})
	} else {
		panic(p)
	}
}
