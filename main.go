package main

import (
	"dreamdump/cli"
	"dreamdump/log"
	"os"
)

func main() {
	option := cli.SetupOptions()
	commandFound := cli.ExecuteCommand(&option)
	if !commandFound {
		log.WriteLn(os.Args[0] + "<disc> [--drive= --drive-sector-order=]")
	}
}
