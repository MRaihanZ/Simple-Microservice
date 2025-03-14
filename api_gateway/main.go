package main

import (
	"github.com/MRaihanZ/Simple-Microservice/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	//inisialiasai Gin
	router := gin.Default()

	router.GET("/api/v1/users", handlers.GetUsers)

	//membuat route dengan method GET
	router.GET("/api/v1/users/:id", handlers.GetUser)

	router.POST("/api/v1/users", handlers.CreateUser)

	router.GET("/api/v1/logs", handlers.GetLogs)

	//mulai server dengan port 3000
	router.Run(":3000")
}
