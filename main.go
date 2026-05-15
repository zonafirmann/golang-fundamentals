package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/zonafirmann/golang-fundamentals/models"
)

// Define a constant for our database file name
const fileName = "tasks.json"

// LoadTasks reads the JSON file and converts it back to a Go slice.
func LoadTasks() []models.Task {
	var tasks []models.Task

	// Attempt to read the physical file
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		// If file doesn't exist yet, return an empty slice safely
		return tasks
	}

	// UNMARSHAL: Translate JSON bytes back into the Go 'tasks' slice
	json.Unmarshal(bytes, &tasks)
	return tasks
}

// SaveTasks converts the Go slice into JSON and writes it to the hard drive.
func SaveTasks(tasks []models.Task) {
	// MARSHAL: Translate Go slice into beautifully formatted JSON bytes
	bytes, _ := json.MarshalIndent(tasks, "", "  ")

	// Write the bytes to a file with standard 0644 read/write permissions
	os.WriteFile(fileName, bytes, 0644)
}

func main() {
	fmt.Println("=== SYSTEM LOG: PERSISTENT CLI ===")

	// 1. Load existing tasks from the hard drive (Storage)
	myTasks := LoadTasks()
	fmt.Printf("[INFO] Loaded %d tasks from database.\n", len(myTasks))

	// 2. Prepare the keyboard reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a new task for today (or press Enter to skip): ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	cleanInput := strings.TrimSpace(input)

	// 3. Only process if the user actually typed something
	if cleanInput != "" {
		newTask := models.Task{
			ID:     len(myTasks) + 1,
			Title:  cleanInput,
			IsDone: false,
		}
		// Add new task to memory
		myTasks = append(myTasks, newTask)

		// Save the updated memory back to the physical file
		SaveTasks(myTasks)
		fmt.Println("[SUCCESS] Task saved permanently to tasks.json")
	}

	fmt.Println("\n--- CURRENT TASK LIST ---")
	for _, task := range myTasks {
		task.DisplayStatus()
	}
}
