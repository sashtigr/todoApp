package main

import (
	"github.com/gin-gonic/gin"
	"todoApp/internal/database"
	"todoApp/internal/handlers"
)

func main() {
	db := database.InitDB()
	database.MigrateDB(db)

	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)
	r.POST("/addTask", handlers.AddTask)
	r.Run(":8080")
}
