package main

import (
	"fmt"
	"strings"
	"time"
)

func commandAdd(args []string, commands map[string]Command) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: task-cli add <description>")
	}

	if err := ensureFileExists(); err != nil {
		return err
	}

	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		ID:          len(tasks) + 1,
		Description: strings.Join(args, " "),
		Status:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	if err = saveTasks(tasks); err != nil {
		return err
	}

	return nil
}
