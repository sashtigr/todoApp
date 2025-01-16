package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoApp/internal/database"
)

func GetTasks(c *gin.Context) {
	var tasks []database.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"status": tasks})
}

func AddTask(c *gin.Context) {
	var task database.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
	}
	database.DB.Create(&task)
	c.JSON(http.StatusOK, gin.H{"Status": "Success"})
}

func CreateTask(c *gin.Context) {
	var task database.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}
