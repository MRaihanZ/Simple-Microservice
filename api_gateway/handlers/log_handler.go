package handlers

import "github.com/gin-gonic/gin"

func GetLogs(c *gin.Context) {
	id := c.Param("id")

	//return response JSON
	c.JSON(200, gin.H{
		"message": "Hello World!" + id,
	})
}
