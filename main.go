package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Note  string `json:"note"`
	Done  bool   `json:"done"`
}

const taskFile = "tasks.json"

func main() {
	add := flag.String("add", "", "Task to add")
	list := flag.Bool("list", false, "List all tasks")
	note := flag.String("note", "", "Add task note (optional)")
	done := flag.Int("done", 0, "Mark a task as done by ID")

	flag.Parse()

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		newID := len(tasks) + 1
		tasks = append(tasks, Task{ID: newID, Title: *add, Note: *note, Done: false})
		fmt.Println("Task added:", *add)

	case *list:
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
		}
		for _, t := range tasks {
			status := ">Not done<"
			if t.Done {
				status = ">Done<"
			}

			fmt.Printf("[%d] %s %s\n", t.ID, status, t.Title)

			if t.Note != "" {
				fmt.Printf(" — %s", t.Note)
			}
			fmt.Println()
		}

	case *done > 0:
		found := false
		for i := range tasks {
			if tasks[i].ID == *done {
				tasks[i].Done = true
				found = true
				fmt.Printf("Task %d marked as done!\n", *done)
				break
			}
		}
		if !found {
			fmt.Println("Task not found.")
		}
	default:
		fmt.Println("No command provided. Use --add, --list, or --done")
	}

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(taskFile)
	if os.IsNotExist(err) {
		return []Task{}, nil
	} else if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}
