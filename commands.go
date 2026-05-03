package main

type Command struct {
	Name        string
	Description string
	Handler     func(args []string, commands map[string]Command) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Show help",
			Handler:     commandHelp,
		},
		"add": {
			Name:        "add",
			Description: "Add a new task",
			Handler:     commandAdd,
		},
		"update": {
			Name:        "update",
			Description: "Update a tasks description",
			Handler:     commandUpdate,
		},
	}
}
