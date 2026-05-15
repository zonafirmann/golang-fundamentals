package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv" // Package to convert string to integer

	"github.com/zonafirmann/golang-fundamentals/models"
)

const fileName = "tasks.json"

func LoadTasks() []models.Task {
	var tasks []models.Task
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return tasks
	}
	json.Unmarshal(bytes, &tasks)
	return tasks
}

func SaveTasks(tasks []models.Task) {
	bytes, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, bytes, 0644)
}

// ---------------------------------------------------------
// CONTROLLER: Handles Full CRUD (GET, POST, PUT, DELETE)
// ---------------------------------------------------------
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// READ: Return all tasks
		myTasks := LoadTasks()
		json.NewEncoder(w).Encode(myTasks)
		fmt.Println("[LOG] Processed GET request")

	case http.MethodPost:
		// CREATE: Add a new task
		var newTask models.Task
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		myTasks := LoadTasks()
		newTask.ID = len(myTasks) + 1
		newTask.IsDone = false

		myTasks = append(myTasks, newTask)
		SaveTasks(myTasks)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
		fmt.Printf("[LOG] Processed POST request. Created: %s\n", newTask.Title)

	case http.MethodPut:
		// UPDATE: Mark a task as completed based on ID in URL parameter (?id=1)
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || idStr == "" {
			http.Error(w, "Invalid or missing task ID", http.StatusBadRequest)
			return
		}

		myTasks := LoadTasks()
		taskFound := false

		// Loop through tasks to find the matching ID
		for i, task := range myTasks {
			if task.ID == id {
				myTasks[i].IsDone = true // Mark as done
				taskFound = true
				break
			}
		}

		if !taskFound {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		SaveTasks(myTasks)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Task successfully marked as completed"}`))
		fmt.Printf("[LOG] Processed PUT request. Task %d updated.\n", id)

	case http.MethodDelete:
		// DELETE: Remove a task based on ID in URL parameter (?id=1)
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || idStr == "" {
			http.Error(w, "Invalid or missing task ID", http.StatusBadRequest)
			return
		}

		myTasks := LoadTasks()
		taskFound := false

		for i, task := range myTasks {
			if task.ID == id {
				// The Go magic to delete an item from a slice
				myTasks = append(myTasks[:i], myTasks[i+1:]...)
				taskFound = true
				break
			}
		}

		if !taskFound {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		SaveTasks(myTasks)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Task successfully deleted"}`))
		fmt.Printf("[LOG] Processed DELETE request. Task %d removed.\n", id)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	fmt.Println("=== SYSTEM LOG: FULL CRUD RESTFUL API ===")
	http.HandleFunc("/tasks", tasksHandler)
	port := ":8080"
	fmt.Printf("[INFO] Server is listening on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("[FATAL ERROR]", err)
	}
}
