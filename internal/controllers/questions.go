package controllers


import (
	"github.com/gin-gonic/gin"
	"github.com/dgedge/quest/internal/models"
	"net/http"
)

// GET /questions
// Get all Questions
func FindQuestions(c *gin.Context) {
	var questions []models.Question
	models.DB.Find(&questions)

	c.JSON(http.StatusOK, gin.H{"data": questions})
}