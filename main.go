package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zonafirmann/golang-fundamentals/models"
)

func main() {
	fmt.Println("=== SYSTEM LOG: INTERACTIVE CLI ===")

	// Initialize slice with existing data
	var myTasks = []models.Task{
		{ID: 1, Title: "Master Workspace & Git Setup", IsDone: true},
		{ID: 2, Title: "Understand Go Packages", IsDone: true},
	}

	// Initialize standard input reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a new task for today: ")

	// Read user input until the Enter key is pressed
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Sanitize the input
	cleanInput := strings.TrimSpace(input)

	// Create a new Task object
	newTask := models.Task{
		ID:     len(myTasks) + 1,
		Title:  cleanInput,
		IsDone: false,
	}

	// Dynamically add the new task to the slice
	myTasks = append(myTasks, newTask)

	fmt.Println("\n--- UPDATED TASK LIST ---")
	for _, task := range myTasks {
		task.DisplayStatus()
	}
}
