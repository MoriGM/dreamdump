package main

import (
	"os"

	"dreamdump/cd"
	"dreamdump/cd/sections"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/sgio"
)

func readDisc(option option.Option) {
	for _, section := range sections.GetSectionMap(option) {
		log.WriteLineNew("{} {}", section.StartSector, section.EndSector)
	}
	sector, _ := cd.ReadSector(&option, 0)
	log.WriteCleanLine(sector)
}

func setupOptions() option.Option {
	if len(os.Args) != 2 {
		panic(os.Args[0] + "[OPTIONS] <drive>")
	}

	option := option.Option{
		SectorOrder: option.DATA_C2_SUB,
		Device:      os.Args[1],
		CutOff:      sections.DC_DEFAULT_CUTOFF,
	}

	dvdDriveDeviceFile, err := sgio.OpenScsiDevice(option.Device)
	if err != nil {
		log.WriteLineNew("This drive is unkown")
		os.Exit(1)
	}
	option.Drive = dvdDriveDeviceFile

	return option
}

func main() {
	option := setupOptions()

	readDisc(option)
}
