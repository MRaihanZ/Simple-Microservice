package main

import (
	"github.com/MRaihanZ/Simple-Microservice/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/v1/users", handlers.GetUsersHandler)

	// router.GET("/api/v1/users/:id", handlers.GetUserHandler)

	// router.POST("/api/v1/users", handlers.CreateUserHandler)

	router.GET("/api/v1/logs", handlers.GetLogsHandler)

	router.Run(":3000")
}
