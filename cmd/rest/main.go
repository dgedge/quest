package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine


func main() {
	router = gin.Default()
	// load templates
	router.LoadHTMLGlob("cmd/rest/templates/*")
	// Initialize the routes
	InitializeRoutes()

	/*
	models.ConnectDataBase()

	router.GET("/questions", controllers.FindQuestions)
	*/
	router.Run()
}
