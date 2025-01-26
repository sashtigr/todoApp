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
	r.GET("/tasks/:id", handlers.GetTaskWithId)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/task/:id", handlers.DeleteTask)
	r.Run(":8080")
}
