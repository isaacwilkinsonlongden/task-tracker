package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: task-cli <command>")
		os.Exit(1)
	}

	commands := getCommands()
	userCommand := os.Args[1]

	cmd, ok := commands[userCommand]
	if !ok {
		fmt.Fprintln(os.Stderr, "Command not found")
		os.Exit(1)
	}

	if err := cmd.Handler(os.Args[2:], commands); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
