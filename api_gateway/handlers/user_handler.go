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
