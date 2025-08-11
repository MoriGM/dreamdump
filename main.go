package main

import (
	"os"

	"dreamdump/cd"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/scsi"
	"dreamdump/sections"
	"dreamdump/sgio"
)

func readDisc(option option.Option) {
	for _, section := range sections.GetSectionMap() {
		log.WriteLineNew("{} {}", section.StartSector, section.EndSector)
	}
	sector, _ := cd.ReadSector(option, 0)
	log.WriteCleanLine(sector)
}

func setupOptions() option.Option {
	if len(os.Args) != 2 {
		panic(os.Args[0] + " <drive>")
	}

	option := option.Option{
		SectorOrder: scsi.DATA_C2_SUB,
		Device:      os.Args[1],
		CutOff:      sections.DC_CUTOFF,
	}

	dvdDriveDeviceFile, err := sgio.OpenScsiDevice(option.Device)
	if err != nil {
		log.WriteLineNew("This drive is unkown")
	}
	option.Drive = dvdDriveDeviceFile

	return option
}

func main() {
	option := setupOptions()

	readDisc(option)
}
