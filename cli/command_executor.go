package cli

import (
	"os"

	"dreamdump/cli/commands"
	"dreamdump/exit_codes"
	"dreamdump/log"
	"dreamdump/option"
)

type Command struct {
	Name     string
	Function func(opt *option.Option)
}

var cliCommands []*Command

func init() {
	cliCommands = append(cliCommands, &Command{
		Name:     "disc",
		Function: commands.DreamDumpDisc,
	}, &Command{
		Name:     "split",
		Function: commands.DreamDumpSplit,
	})
}

func ExecuteCommand(opt *option.Option) bool {
	if len(os.Args) < 2 {
		log.Println("Missing command argument")
		os.Exit(exit_codes.MISSING_COMMAND_ARGUMENTS)
	}
	for _, command := range cliCommands {
		if command.Name == os.Args[1] {
			command.Function(opt)
			return true
		}
	}
	return false
}
