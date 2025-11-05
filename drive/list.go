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
	ReadAtOnce         uint8
	Train              bool
}

var DriveList []*DriveOption

func init() {
	DriveList = []*DriveOption{
		// Good
		{"TSSTcorp", "DVD-ROM SH-D162C", "DC02", option.DATA_SUB_C2, +6, 12, false},
		{"TSSTcorp", "DVD-ROM SH-D163A", "DC02", option.DATA_SUB_C2, +6, 12, false},
		{"LITE-ON", "DVD D LH-16D1P", "ZZ00", option.DATA_SUB_C2, +6, 20, false},
		{"TSSTcorp", "DVD-ROM SH-D163B", "ZZ01", option.DATA_C2_SUB, +6, 12, true},
		{"TSSTcorp", "DVD-ROM SH-D163B", "ZZ02", option.DATA_C2_SUB, +6, 12, true},
		// Meh
		{"LITE-ON", "DVD D DH-16D3S", "ZZ00", option.DATA_C2_SUB, +6, 32, true},
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
