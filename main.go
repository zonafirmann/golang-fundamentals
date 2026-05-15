package main

import (
	"fmt"
	// Mengimpor package 'models' yang baru kita buat
	// Format: nama_module_di_go.mod/nama_folder
	"github.com/zonafirmann/golang-fundamentals/models"
)

func main() {
	fmt.Println("=== SYSTEM LOG: MODULAR ARCHITECTURE ===")

	// Memanggil struct Task dari package models
	var myTasks = []models.Task{
		{ID: 1, Title: "Master Workspace & Git Setup", IsDone: true},
		{ID: 2, Title: "Understand Go Packages and Imports", IsDone: false},
		{ID: 3, Title: "Implement Modular Architecture", IsDone: false},
	}

	for _, task := range myTasks {
		// Memanggil method dari package models
		task.DisplayStatus()
	}

	fmt.Println("========================================")
	fmt.Printf("System initialized with %d modular tasks.\n", len(myTasks))
}
