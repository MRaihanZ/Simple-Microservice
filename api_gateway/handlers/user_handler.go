package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	//return response JSON
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	//return response JSON
	c.JSON(200, gin.H{
		"message": "Hello World!" + id,
	})
}

func CreateUser(c *gin.Context) {

	//return response JSON
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
