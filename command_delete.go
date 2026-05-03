package main

import (
	"fmt"
	"strconv"
)

func commandDelete(args []string, commands map[string]Command) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: task-cli delete <task ID>")
	}

	if err := ensureFileExists(); err != nil {
		return err
	}

	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	target, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	found := false

	for i := range tasks {
		if tasks[i].ID == target {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", target)
	}

	if err := saveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task deleted successfully (ID: %d)\n", target)

	return nil
}
