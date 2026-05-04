package main

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ensureFileExists() error {
	_, err := os.Stat("tasks.json")

	if os.IsNotExist(err) {
		file, err := os.Create("tasks.json")
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := file.Write([]byte("[]")); err != nil {
			return err
		}
	}

	return nil
}

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}

	var tasks []Task

	if err = json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	if err = os.WriteFile("tasks.json", data, 0644); err != nil {
		return err
	}

	return nil
}
