package models

import (
	"fmt"
	"time"
)

func CreateLog(logs []Log) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	currentTime := time.Now()
	fmt.Println("The time is", currentTime)
	fmt.Println()

	fmt.Printf("%s",
		currentTime.In(loc),
	)
}
