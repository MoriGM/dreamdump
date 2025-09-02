package drive

import (
	"strings"

	"dreamdump/option"
)

type DriveOption struct {
	VendorName         string
	ProductInquiryData string
	RevisionNumber     string
	SectorOrder        int
	ReadOffset         int16
}

var DriveList []*DriveOption

func init() {
	DriveList = []*DriveOption{
		{"TSSTcorp", "DVD-ROM SH-D162C", "DC02", option.DATA_SUB_C2, +6},
		{"TSSTcorp", "DVD-ROM SH-D163A", "DC02", option.DATA_SUB_C2, +6},
		{"LITE-ON", "DVD D LH-16D1P", "ZZ00", option.DATA_SUB_C2, +6},
	}
}

func IsKnownDrive(drive *Drive) *DriveOption {
	for _, knownDrive := range DriveList {
		if knownDrive.VendorName != strings.Trim(string(drive.VendorName[:]), " ") {
			continue
		}
		if knownDrive.ProductInquiryData != strings.Trim(string(drive.ProductInquiryData[:]), " ") {
			continue
		}
		if knownDrive.RevisionNumber != strings.Trim(string(drive.RevisionNumber[:]), " ") {
			continue
		}
		return knownDrive
	}
	return nil
}
