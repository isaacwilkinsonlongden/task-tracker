package main

type Command struct {
	Name        string
	Description string
	Handler     func(args []string, commands map[string]Command)
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Show help",
			Handler:     commandHelp,
		},
	}
}
