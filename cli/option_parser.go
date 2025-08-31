package cli

import (
	"os"
	"strconv"
	"strings"

	"dreamdump/drive"
	"dreamdump/exit_codes"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/scsi/scsi_commands"
	"dreamdump/sgio"
)

const (
	CD_SPEED = 176
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
		CutOff:      option.DC_DEFAULT_CUTOFF,
		ImageName:   "Game",
		PathName:    "./Game",
		ReadOffset:  0,
		Speed:       0xFFFF,
	}

	device := FindArgumentString("drive")
	if device != nil {
		opt.Device = *device
	}

	dvdDriveDeviceFile, err := sgio.OpenScsiDevice(opt.Device)
	if err != nil {
		log.Println("This drive is unkown or is missing it gd-rom")
		os.Exit(exit_codes.UNKOWN_DRIVE)
	}
	opt.Drive = dvdDriveDeviceFile

	currentDrive := scsi_commands.Inquiry(&opt)
	knownDrive := drive.IsKnownDrive(currentDrive)
	if knownDrive != nil {
		log.Println("Good Drive found.")
		opt.SectorOrder = knownDrive.SectorOrder
		opt.ReadOffset = knownDrive.ReadOffset
	}
	sectorOrder := FindArgumentString("sector-order")
	if sectorOrder != nil {
		if (*sectorOrder) == "DATA_C2" {
			opt.SectorOrder = option.DATA_C2
		}
		if (*sectorOrder) == "DATA_SUB" {
			opt.SectorOrder = option.DATA_SUB
		}
		if (*sectorOrder) == "DATA_C2_SUB" {
			opt.SectorOrder = option.DATA_C2_SUB
		}
		if (*sectorOrder) == "DATA_SUB_C2" {
			opt.SectorOrder = option.DATA_SUB_C2
		}
	}
	currentDrive.PrintDriveInfo(&opt)

	imageName := FindArgumentString("image-name")
	if imageName != nil {
		opt.ImageName = *imageName
	}

	pathName := FindArgumentString("image-path")
	if pathName != nil {
		opt.PathName = *pathName
	}

	readOffset := FindArgumentString("read-offset")
	if readOffset != nil {
		offset, err := strconv.ParseInt(*readOffset, 10, 16)
		if err != nil {
			panic(err)
		}
		opt.ReadOffset = int16(offset)
	}

	speed := FindArgumentString("speed")
	if speed != nil {
		speed, err := strconv.ParseInt(*speed, 10, 16)
		if err != nil {
			panic(err)
		}
		if speed > 48 {
			speed = 48
		}
		opt.Speed = uint16(speed) * CD_SPEED
	}

	if opt.CutOff > option.DC_END {
		log.Println("Cutoff can not be bigger than the Disc")
		os.Exit(exit_codes.CUTOFF_TO_BIG)
	}

	return opt
}
