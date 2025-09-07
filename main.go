package main

import (
	"os"

	"dreamdump/cli"
	"dreamdump/log"
)

func main() {
	option := cli.SetupOptions()
	log.Setup(&option)
	commandFound := cli.ExecuteCommand(&option)
	if !commandFound {
		log.Println(os.Args[0] + " <disc> [--drive= --sector-order= --image-path= --image-name= --speed= --read-offset= --cutoff=]")
	}
}
