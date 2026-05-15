package main

import "fmt"

// 1. STRUCT: This is the Blueprint (Cetakan) of our data.
// In the future, this is exactly how you will map JSON data from an API.
type Task struct {
	ID     int
	Title  string
	IsDone bool
}

// 2. FUNCTION (Method): Giving behavior to our Struct.
// This function specifically belongs to the 'Task' struct.
func (t Task) DisplayStatus() {
	status := "⏳ [IN PROGRESS]"
	if t.IsDone {
		status = "✅ [COMPLETED]"
	}
	fmt.Printf("Task %d: %s %s\n", t.ID, t.Title, status)
}

func main() {
	fmt.Println("=== SYSTEM LOG: ADVANCED TASK MANAGER ===")

	// 3. SLICE OF STRUCTS: Creating a list of complex objects
	var myTasks = []Task{
		{ID: 1, Title: "Master Workspace & Git Setup", IsDone: true},
		{ID: 2, Title: "Understand Go Control Flow", IsDone: true},
		{ID: 3, Title: "Implement Structs and Functions", IsDone: false},
		{ID: 4, Title: "Build RESTful API", IsDone: false},
	}

	// 4. CLEAN EXECUTION: The main function becomes very clean and easy to read
	for _, task := range myTasks {
		// Calling the method attached to the struct
		task.DisplayStatus()
	}

	fmt.Println("=========================================")
	fmt.Printf("Total tasks loaded into memory: %d\n", len(myTasks))
}
