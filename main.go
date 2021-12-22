package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InterpretationResponse struct {
	Result string

	Err bool

	ErrMessage string
}

func main() {

	router := gin.Default()

	router.GET("/interpret", Interpret)

	router.Run("localhost:8080")
}

func Interpret(c *gin.Context) {

	response := InterpretationResponse{Result: "1", Err: false, ErrMessage: "None"}

	c.IndentedJSON(http.StatusOK, response)
}
