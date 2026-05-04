package main

import (
	"fmt"
	"slices"
)

func commandList(args []string, commands map[string]Command) error {
	correctArgs := []string{"todo", "in-progress", "done"}

	if len(args) > 0 {
		if !slices.Contains(correctArgs, args[0]) {
			return fmt.Errorf("invalid status '%s'; expected todo, in-progress, or done", args[0])
		}
	}

	if err := ensureFileExists(); err != nil {
		return err
	}

	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		for _, task := range tasks {
			fmt.Printf("[ID: %d] %s (%s)\n", task.ID, task.Description, task.Status)
		}
	} else {
		for _, task := range tasks {
			if task.Status == args[0] {
				fmt.Printf("[ID: %d] %s (%s)\n", task.ID, task.Description, task.Status)
			}
		}
	}

	return nil
}
