package cli

import (
	"os"
	"strings"

	"dreamdump/cd/sections"
	"dreamdump/exit_codes"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/sgio"
)

func FindArgumentString(name string) *string {
	for _, arg := range os.Args {
		if parts := strings.Split(arg, "="); parts[0] == ("--"+name) && len(parts) == 2 {
			return &parts[1]
		}
	}
	return nil
}

func SetupOptions() option.Option {
	opt := option.Option{
		SectorOrder: option.DATA_SUB_C2,
		Device:      "/dev/sr0",
		CutOff:      sections.DC_DEFAULT_CUTOFF,
	}

	device := FindArgumentString("drive")
	if device != nil {
		opt.Device = *device
	}

	dvdDriveDeviceFile, err := sgio.OpenScsiDevice(opt.Device)
	if err != nil {
		log.Println("This drive is unkown")
		os.Exit(exit_codes.UNKOWN_DRIVE)
	}
	opt.Drive = dvdDriveDeviceFile

	sectorOrder := FindArgumentString("sector-order")
	if sectorOrder != nil {
		if *sectorOrder == "DATA_C2" {
			opt.SectorOrder = option.DATA_C2
		}
		if *sectorOrder == "DATA_SUB" {
			opt.SectorOrder = option.DATA_SUB
		}
		if *sectorOrder == "DATA_C2_SUB" {
			opt.SectorOrder = option.DATA_C2_SUB
		}
		if *sectorOrder == "DATA_SUB_C2" {
			opt.SectorOrder = option.DATA_SUB_C2
		}
	}

	if opt.CutOff > sections.DC_END {
		log.Println("Cutoff can not be bigger than the Disc")
		os.Exit(exit_codes.CUTOFF_TO_BIG)
	}

	return opt
}
