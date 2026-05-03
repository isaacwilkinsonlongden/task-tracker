package main

import "fmt"

func commandList(args []string, commands map[string]Command) error {
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
