package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func commandUpdate(args []string, commands map[string]Command) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: task-cli update <task ID> <description>")
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
			tasks[i].Description = strings.Join(args[1:], " ")
			tasks[i].UpdatedAt = time.Now()
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

	fmt.Printf("Task updated successfully (ID: %d)\n", target)

	return nil
}
