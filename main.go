package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	commands := getCommands()
	userCommand := os.Args[1]

	cmd, ok := commands[userCommand]
	if !ok {
		fmt.Fprintln(os.Stderr, "Command not found")
		os.Exit(1)
	}

	cmd.Handler(os.Args[2:], commands)
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: task-cli <command>")
	os.Exit(1)
}
