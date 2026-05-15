package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zonafirmann/golang-fundamentals/models"
)

// Define a constant for our database file name
const fileName = "tasks.json"

// LoadTasks reads the JSON database and returns a slice of tasks.
func LoadTasks() []models.Task {
	var tasks []models.Task
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return tasks
	}
	// Translate JSON bytes back into the Go slice
	json.Unmarshal(bytes, &tasks)
	return tasks
}

// ---------------------------------------------------------
// CONTROLLER: Handles incoming HTTP requests for tasks
// ---------------------------------------------------------
func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Retrieve data from the persistent storage
	myTasks := LoadTasks()

	// 2. Set the HTTP Header to indicate JSON payload
	w.Header().Set("Content-Type", "application/json")

	// 3. Encode the Go slice into JSON and send it as the response
	json.NewEncoder(w).Encode(myTasks)
	
	fmt.Println("[LOG] GET request received at /tasks endpoint!")
}

func main() {
	fmt.Println("=== SYSTEM LOG: INITIALIZING RESTFUL API ===")

	// 1. Register the route (Endpoint) and its corresponding handler
	http.HandleFunc("/tasks", getTasksHandler)

	// 2. Define the port for the local web server
	port := ":8080"
	fmt.Printf("[INFO] Server is running on http://localhost%s\n", port)
	fmt.Println("[INFO] Press CTRL+C in terminal to stop the server.")

	// 3. Start the HTTP server to listen for incoming network requests
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("[FATAL ERROR] Server failed to start:", err)
	}
}