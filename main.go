package main

import (
	"os"

	"dreamdump/cli"
	"dreamdump/log"
)

const (
	VERSION = "0.0.0-rc1"
)

func main() {
	option := cli.SetupOptions()
	log.Setup(&option)
	log.Println("Version: " + VERSION)
	commandFound := cli.ExecuteCommand(&option)
	if !commandFound {
		log.Println(os.Args[0] + " <disc> [--drive= --sector-order= --image-path= --image-name= --speed= --read-offset= --cutoff=]")
	}
}
