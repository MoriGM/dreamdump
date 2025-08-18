package cli

import (
	"os"
	"strings"

	"dreamdump/cd/sections"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/sgio"
)

func FindArgumentString(name string) *string {
	for _, arg := range os.Args {
		if parts := strings.Split(arg, "="); parts[0] == name && len(parts) == 2 {
			return &parts[0]
		}
	}
	return nil
}

func SetupOptions() option.Option {
	opt := option.Option{
		SectorOrder: option.DATA_C2_SUB,
		Device:      "/dev/sr0",
		CutOff:      sections.DC_DEFAULT_CUTOFF,
	}

	device := FindArgumentString("--drive")
	if device != nil {
		opt.Device = *device
	}

	sectorOrder := FindArgumentString("--sector-order")
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

	dvdDriveDeviceFile, err := sgio.OpenScsiDevice(opt.Device)
	if err != nil {
		log.WriteLn("This drive is unkown")
		os.Exit(1)
	}
	opt.Drive = dvdDriveDeviceFile

	return opt
}
