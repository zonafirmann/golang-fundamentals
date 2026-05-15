// Baris pertama WAJIB nama folder tempat file ini berada
package models

import "fmt"

// Task represents a unit of work in our system.
// Notice that 'Task' starts with a CAPITAL letter.
type Task struct {
	ID     int
	Title  string
	IsDone bool
}

// DisplayStatus prints the current state of the task.
// Method name also starts with a CAPITAL letter.
func (t Task) DisplayStatus() {
	status := "⏳ [IN PROGRESS]"
	if t.IsDone {
		status = "✅ [COMPLETED]"
	}
	fmt.Printf("Task %d: %s %s\n", t.ID, t.Title, status)
}
