package main

import (
	"dreamdump/cd"
	"dreamdump/cd/sections"
	"dreamdump/cli"
	"dreamdump/log"
	"dreamdump/option"
)

func readDisc(option option.Option) {
	for _, section := range sections.GetSectionMap(option) {
		log.WriteLn("{} {}", section.StartSector, section.EndSector)
	}
	sector, _ := cd.ReadSector(&option, 0)
	log.WriteCleanLine(sector)
}

func main() {
	option := cli.SetupOptions()
	cli.ExecuteCommand(&option)
}
