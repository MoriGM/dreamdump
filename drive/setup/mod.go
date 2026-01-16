package setup

import (
	"os"

	"dreamdump/drive"
	"dreamdump/exit_codes"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/scsi/driver"
	"dreamdump/scsi/scsi_commands"
)

func InitializeDrive(opt *option.Option) {
	driveDeviceFile, err := driver.OpenScsiDevice(opt.Device)
	if err != nil {
		log.Println("This drive is unknown or is missing a gd-rom")
		os.Exit(exit_codes.UNKNOWN_DRIVE)
	}

	opt.Drive = driveDeviceFile

	currentDrive := scsi_commands.Inquiry(opt)
	knownDrive := drive.IsKnownDrive(currentDrive)
	if knownDrive != nil {
		log.Println("Known Drive found.")
		opt.SectorOrder = knownDrive.SectorOrder
		opt.ReadOffset = knownDrive.ReadOffset
		if opt.ReadAtOnce == driver.MAX_READ_AT_ONCE {
			opt.ReadAtOnce = knownDrive.ReadAtOnce
		}
		if opt.ReadAtOnce > driver.MAX_READ_AT_ONCE {
			opt.ReadAtOnce = driver.MAX_READ_AT_ONCE
		}
	}

	currentDrive.PrintDriveInfo(opt)
}
