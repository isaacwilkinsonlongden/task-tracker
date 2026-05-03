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
			Description: "Update a tasks description by ID",
			Handler:     commandUpdate,
		},
		"delete": {
			Name:        "delete",
			Description: "Delete a task by ID",
			Handler:     commandDelete,
		},
		"mark-in-progress": {
			Name:        "mark-in-progress",
			Description: "Update task progress by ID",
			Handler:     commandMarkInProgress,
		},
	}
}
