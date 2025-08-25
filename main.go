package main

import (
	"os"

	"dreamdump/cli"
	"dreamdump/log"
)

func main() {
	option := cli.SetupOptions()
	commandFound := cli.ExecuteCommand(&option)
	if !commandFound {
		log.Println(os.Args[0] + "<disc> [--drive= --drive-sector-order=]")
	}
}
