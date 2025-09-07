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
