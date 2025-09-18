package cli

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"dreamdump/exit_codes"
	"dreamdump/log"
	"dreamdump/option"
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
		Train:       false,
		ReadAtOnce:  26,
	}

	parseSpecial(&opt)
	parsePaths(&opt)
	parseDrivePart(&opt)

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

func parseSpecial(opt *option.Option) {
	opt.QTocSplit = HasArgumentString("force-qtoc")
	opt.Train = HasArgumentString("train")
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

	readAtOnceString := FindArgumentString("read-at-once")
	if readAtOnceString != nil {
		readAtOnce, err := strconv.ParseInt(*readAtOnceString, 10, 8)
		if err != nil {
			panic(err)
		}
		opt.ReadAtOnce = uint8(readAtOnce)
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
