package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define struct for User
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Define Response struct
type Response struct {
	Code int    `json:"code"`
	Body []User `json:"body"`
}

// Function to get users and return Response struct
func GetUsers() ([]User, error) {
	// Read JSON file
	data, err := os.ReadFile("../db/db.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Parse JSON into slice of User structs
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	// Return Response struct
	return users, nil
}

func main() {
	// Call function and handle error
	response, err := GetUsers()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print response as JSON
	jsonResponse, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(jsonResponse))
}
