package models

import "fmt"

// Task represents a unit of work in our system.
// We add "Struct Tags" (the backticks) to tell Go how to translate this into JSON.
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

// DisplayStatus prints the current state of the task.
func (t Task) DisplayStatus() {
	status := "⏳ [IN PROGRESS]"
	if t.IsDone {
		status = "✅ [COMPLETED]"
	}
	fmt.Printf("Task %d: %s %s\n", t.ID, t.Title, status)
}
