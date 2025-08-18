package cli

import (
	"os"

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

func ExecuteCommand(opt *option.Option) error {
	if len(os.Args) < 2 {
		log.WriteLn("Missing command argument")
		os.Exit(4)
	}
	for _, command := range commands {
		if command.Name == os.Args[1] {
			command.Function(opt)
		}
	}
	return nil
}
