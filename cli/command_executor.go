package cli

import (
	"os"

	"dreamdump/exit_codes"
	"dreamdump/log"
	"dreamdump/option"
)

type Command struct {
	Name     string
	Function func(opt *option.Option)
}

var commands []*Command

func init() {
	commands = append(commands, &Command{
		Name:     "disc",
		Function: DreamDumpDisc,
	})
}

func ExecuteCommand(opt *option.Option) bool {
	if len(os.Args) < 2 {
		log.Println("Missing command argument")
		os.Exit(exit_codes.MISSING_COMMAND_ARGUMENTS)
	}
	for _, command := range commands {
		if command.Name == os.Args[1] {
			command.Function(opt)
			return true
		}
	}
	return false
}
