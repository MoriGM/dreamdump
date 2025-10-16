package cli

import (
	"os"
	"strings"

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
	log.Printf("arguments: %s\n", strings.Join(os.Args[1:], " "))
	for _, command := range cliCommands {
		if command.Name == os.Args[1] {
			command.Function(opt)
			return true
		}
	}
	return false
}
