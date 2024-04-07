package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Task struct represents a task in the todo list
type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks []Task

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

// CreateTaskHandler creates a new task
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTask.ID = fmt.Sprintf("%d", len(tasks)+1)
	tasks = append(tasks, newTask)
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	// Define routes
	http.HandleFunc("/tasks", GetTasksHandler)       // Handle GET requests to "/tasks"
	http.HandleFunc("/tasks/new", CreateTaskHandler) // Handle POST requests to "/tasks/new"

	// Initialize tasks
	tasks = append(tasks, Task{ID: "1", Title: "Task 1", Completed: false})
	tasks = append(tasks, Task{ID: "2", Title: "Task 2", Completed: true})

	// Start the server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
