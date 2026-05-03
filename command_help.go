package main

import "fmt"

func commandHelp(args []string, commands map[string]Command) {
	fmt.Println("Commands:")
	fmt.Println()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
}
