package cli

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"dreamdump/drive"
	"dreamdump/exit_codes"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/scsi/driver"
	"dreamdump/scsi/scsi_commands"
)

const (
	CD_SPEED = 176
)

func HasArgumentString(name string) bool {
	return slices.Contains(os.Args, "--"+name)
}

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
		Speed:       0,
		QTocSplit:   false,
	}

	parseSplit(&opt)
	parsePaths(&opt)
	parseDrivePart(&opt)
	initializeDrive(&opt)

	cutoff := FindArgumentString("cutoff")
	if cutoff != nil {
		cutoff, err := strconv.ParseInt(*cutoff, 10, 32)
		if err != nil {
			panic(err)
		}
		if int32(cutoff) > option.DC_END {
			log.Println("Cutoff can not be bigger than the Disc")
			os.Exit(exit_codes.CUTOFF_TO_BIG)
		}
		if int32(cutoff) < option.DC_START {
			log.Println("Cutoff can smaller than the Disc")
			os.Exit(exit_codes.CUTOFF_TO_BIG)
		}
		opt.CutOff = int32(cutoff)
	}

	return opt
}

func parseSplit(opt *option.Option) {
	opt.QTocSplit = HasArgumentString("force-qtoc")
}

func parsePaths(opt *option.Option) {
	imageName := FindArgumentString("image-name")
	if imageName != nil {
		opt.ImageName = *imageName
	}

	pathName := FindArgumentString("image-path")
	if pathName != nil {
		opt.PathName = *pathName
	}
}

func parseDrivePart(opt *option.Option) {
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

	sectorOrder := FindArgumentString("sector-order")
	if sectorOrder != nil {
		opt.SectorOrder = parseSectorOrder(*sectorOrder)
	}

	device := FindArgumentString("drive")
	if device != nil {
		opt.Device = *device
	}
}

func parseSectorOrder(sectorOrder string) int {
	switch sectorOrder {
	case "DATA_C2":
		return option.DATA_C2
	case "DATA_SUB":
		return option.DATA_SUB
	case "DATA_C2_SUB":
		return option.DATA_C2_SUB
	case "DATA_SUB_C2":
		return option.DATA_SUB_C2
	}
	return option.DATA_C2_SUB
}

func initializeDrive(opt *option.Option) {
	driveDeviceFile, err := driver.OpenScsiDevice(opt.Device)
	if err != nil {
		log.Println("This drive is unkown or is missing it gd-rom")
		os.Exit(exit_codes.UNKOWN_DRIVE)
	}

	opt.Drive = driveDeviceFile

	currentDrive := scsi_commands.Inquiry(opt)
	knownDrive := drive.IsKnownDrive(currentDrive)
	if knownDrive != nil {
		log.Println("Good Drive found.")
		opt.SectorOrder = knownDrive.SectorOrder
		opt.ReadOffset = knownDrive.ReadOffset
	}

	currentDrive.PrintDriveInfo(opt)
}
