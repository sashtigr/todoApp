package main

import (
	"github.com/gin-gonic/gin"
	"todoApp/internal/auth"
	"todoApp/internal/database"
	"todoApp/internal/handlers"
)

func main() {
	db := database.InitDB()
	database.MigrateDB(db)

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/tasks", handlers.GetTasks)
		protected.POST("/addTask", handlers.AddTask)
		protected.GET("/tasks/:id", handlers.GetTaskWithId)
		protected.PUT("/tasks/:id", handlers.UpdateTask)
		protected.DELETE("/task/:id", handlers.DeleteTask)
	}

	r.Run(":8080")
}
