package models

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MRaihanZ/Simple-Microservice/entities"
	"google.golang.org/protobuf/encoding/protojson"
)

// Define Response struct
type Response struct {
	Code int            `json:"code"`
	Body []entities.Log `json:"body"`
}

// Function to get users and return Response struct
// func GetUsers() ([]User, error) {
// 	// Read JSON file
// 	data, err := os.ReadFile("../db/db.json")
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading file: %w", err)
// 	}

// 	// Parse JSON into slice of User structs
// 	var users []User
// 	if err := json.Unmarshal(data, &users); err != nil {
// 		return nil, fmt.Errorf("error parsing JSON: %w", err)
// 	}

// 	return users, nil
// }

func CreateLog() {
	jsonData, err := protojson.Marshal(logs)
	if err != nil {
		fmt.Println("Error marshaling proto to JSON:", err)
		return
	}

	// Write JSON data to file
	fileName := "logs.json"
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func main() {
	// Call function and handle error
	response, err := GetLog()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print response as JSON
	jsonResponse, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(jsonResponse))
}
