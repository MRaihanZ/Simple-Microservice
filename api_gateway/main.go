package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//inisialiasai Gin
	router := gin.Default()

	//membuat route dengan method GET
	router.GET("/api/v1/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!" + id,
		})
	})

	router.GET("/api/v1/users/:id/", func(c *gin.Context) {
		id := c.Param("id")
		start := c.Query("start")
		end := c.Query("end")

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!" + id + " | " + start + " | " + end,
		})
	})

	router.POST("/api/v1/users", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/api/v1/logs", func(c *gin.Context) {
		id := c.Param("id")

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!" + id,
		})
	})

	//mulai server dengan port 3000
	router.Run(":3000")
}
