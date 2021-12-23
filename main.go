package main

import (
	"net/http"
	"strconv"
	"symgolic/evaluation"
	"symgolic/parsing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type InterpretationRequest struct {
	Expression string `json:"expression"`
}

type InterpretationResponse struct {
	Result string `json:"result"`

	Err bool `json:"error"`
}

type User struct {
	Id int

	Username string

	Email string

	Password string
}

func main() {

	router := gin.Default()

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowAllOrigins = true

	router.Use(cors.New(corsConfig))

	router.GET("/getusers", GetUsers)

	router.GET("/getuser/:id", GetUserById)

	router.POST("/parse", Parse)

	router.POST("/command/sumliketerms", SumLikeTerms)
	router.POST("/command/factor", Factor)

	router.Run("localhost:8000")
}

func GetUsers(c *gin.Context) {

	users := []User{
		{Id: 1, Username: "michael", Email: "m.tracey@gmail.com", Password: "password"},
		{Id: 2, Username: "alex", Email: "a.smith@gmail.com", Password: "password"},
	}
	c.IndentedJSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	users := []User{
		{Id: 1, Username: "michael", Email: "m.tracey@gmail.com", Password: "password"},
		{Id: 2, Username: "alex", Email: "a.smith@gmail.com", Password: "password"},
	}
	for _, user := range users {

		if user.Id == id {

			c.IndentedJSON(http.StatusOK, user)

			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, "no user matches that Id")
}

func Parse(c *gin.Context) {

	var request InterpretationRequest

	if err := c.BindJSON(&request); err != nil {

		return
	}
	result := parsing.ParseExpression(request.Expression)

	response := InterpretationResponse{Result: result.ToString(), Err: false}

	c.IndentedJSON(http.StatusOK, response)
}

func SumLikeTerms(c *gin.Context) {

	var request InterpretationRequest

	if err := c.BindJSON(&request); err != nil {

		return
	}
	result := parsing.ParseExpression(request.Expression)

	evaluation.EvaluateAndReplace(result.GetRoot(), &result, evaluation.EvaluateLikeTerms)

	response := InterpretationResponse{Result: result.ToString(), Err: false}

	c.IndentedJSON(http.StatusOK, response)
}

func Factor(c *gin.Context) {

	var request InterpretationRequest

	if err := c.BindJSON(&request); err != nil {

		return
	}
	result := parsing.ParseExpression(request.Expression)

	evaluation.EvaluateAndReplace(result.GetRoot(), &result, evaluation.EvaluateFactorisation)

	response := InterpretationResponse{Result: result.ToString(), Err: false}

	c.IndentedJSON(http.StatusOK, response)
}

func IsEqual(c *gin.Context) {

}
