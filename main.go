package main

import (
	"dreamdump/cli"
)

func main() {
	option := cli.SetupOptions()
	cli.ExecuteCommand(&option)
}
