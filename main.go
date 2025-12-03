package main

import (
	"os"

	"dreamdump/cli"
	"dreamdump/log"
)

const (
	VERSION = "0.2.0"
)

func main() {
	option := cli.SetupOptions()
	log.Setup(&option)
	log.Printf("dreamdump (build: %s)\n\n", VERSION)
	commandFound := cli.ExecuteCommand(&option)
	if !commandFound {
		log.Println(os.Args[0] + " <disc,split> [--drive= --sector-order= --image-path= --image-name= --speed= --read-offset= --cutoff=]")
	}
}
