package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
	task := Task{
		ID:          1,
		Description: "Complete the project documentation",
		Status:      "In Progress",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
	fmt.Println(task)
}
