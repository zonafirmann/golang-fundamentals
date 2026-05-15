package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zonafirmann/golang-fundamentals/models"
)

const fileName = "tasks.json"

// LoadTasks reads the JSON database and returns a slice of tasks.
func LoadTasks() []models.Task {
	var tasks []models.Task
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return tasks
	}
	json.Unmarshal(bytes, &tasks)
	return tasks
}

// SaveTasks writes the slice of tasks back to the JSON database.
func SaveTasks(tasks []models.Task) {
	bytes, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, bytes, 0644)
}

// ---------------------------------------------------------
// CONTROLLER: Handles GET (Read) and POST (Create) requests
// ---------------------------------------------------------
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	// Set the HTTP Header to indicate JSON payload for all responses
	w.Header().Set("Content-Type", "application/json")

	// Determine the action based on the HTTP Method
	switch r.Method {
	case http.MethodGet:
		// ACTION: Client wants to read data
		myTasks := LoadTasks()
		json.NewEncoder(w).Encode(myTasks)
		fmt.Println("[LOG] Processed GET request")

	case http.MethodPost:
		// ACTION: Client wants to add new data
		var newTask models.Task

		// Decode the incoming JSON body into our Go struct
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			// If client sends bad JSON, return a 400 Bad Request error
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		// Load existing tasks to calculate the next ID
		myTasks := LoadTasks()
		newTask.ID = len(myTasks) + 1
		newTask.IsDone = false // Default status for new tasks

		// Add the new task and save it permanently
		myTasks = append(myTasks, newTask)
		SaveTasks(myTasks)

		// Respond with 201 Created status and send back the created task
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
		fmt.Printf("[LOG] Processed POST request. Created: %s\n", newTask.Title)

	default:
		// ACTION: Client uses an unsupported method (e.g., PUT, DELETE)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	fmt.Println("=== SYSTEM LOG: RESTFUL API V2 ===")

	// Register the endpoint with the unified handler
	http.HandleFunc("/tasks", tasksHandler)

	port := ":8080"
	fmt.Printf("[INFO] Server is listening on http://localhost%s\n", port)
	fmt.Println("[INFO] Press CTRL+C to stop.")

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("[FATAL ERROR]", err)
	}
}
