package handlers

import (
	"fmt"

	"github.com/MRaihanZ/Simple-Microservice/clients"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

func GetUsersHandler(c *gin.Context) {
	fmt.Println("Connecting...")
	conn, client := clients.ConnectUserService()
	fmt.Println("Connected")
	response := clients.GetUsers(conn, client)
	result, err := protojson.Marshal(response)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to serialize response"})
		return
	}
	//return response JSON
	c.Data(200, "application/json", result)
}

// func GetUserHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	//return response JSON
// 	c.JSON(200, gin.H{
// 		"message": "Hello World!" + id,
// 	})
// }

// func CreateUserHandler(c *gin.Context) {
// 	name := c.PostForm("name")
// 	email := c.PostForm("email")
// 	password := c.PostForm("password")

// 	conn, client := clients.ConnectUserService()

// 	//return response JSON
// 	c.JSON(200, gin.H{
// 		"message": "Hello World! POST",
// 	})
// }
