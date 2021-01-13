package main

import (
	"github.com/dgedge/quest/internal/controllers"
	"github.com/dgedge/quest/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	r.GET("/questions", controllers.FindQuestions)

	r.Run()
}
