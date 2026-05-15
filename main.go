package main

import "fmt"

func main() {
	// 1. SLICES: Creating a dynamic list of daily engineering tasks
	tasks := []string{
		"Master Workspace & Git Setup",
		"Understand Go Control Flow",
		"Build a CLI Task Manager",
	}

	fmt.Println("=== SYSTEM LOG: DAILY TASKS ===")

	// 2. IF-ELSE: Logic to check workload capacity
	// len(tasks) is a built-in function to count items in a slice
	if len(tasks) > 5 {
		fmt.Println("[WARNING] CPU Overload: Too many tasks scheduled today!")
	} else if len(tasks) == 0 {
		fmt.Println("[ALERT] Idle status: No tasks assigned.")
	} else {
		fmt.Printf("[INFO] Workload optimal. %d tasks in queue.\n", len(tasks))
	}

	fmt.Println("--------------------------------")

	// 3. FOR LOOP: Iterating through the slice
	// 'range' automatically gives us the Index (0, 1, 2) and the Value (the string)
	for index, task := range tasks {
		// We use index+1 so the display starts at 1 instead of 0
		fmt.Printf("Task %d: %s\n", index+1, task)
	}

	fmt.Println("================================")
}
