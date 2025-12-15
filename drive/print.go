package drive

import (
	"dreamdump/log"
	"dreamdump/option"
	"strings"
)

func (drive *Drive) PrintDriveInfo(opt *option.Option) {
	sectorOrder := "DATA"
	switch opt.SectorOrder {
	case option.DATA_C2:
		sectorOrder = "DATA_C2"
	case option.DATA_C2_SUB:
		sectorOrder = "DATA_C2_SUB"
	case option.DATA_SUB:
		sectorOrder = "DATA_SUB"
	case option.DATA_SUB_C2:
		sectorOrder = "DATA_SUB_C2"
	}
	log.Printf("Drive: %s - %s (revision level: %s, vendor specific: %s)  Read Offset: %d  Sector Order: %s  Read at once: %d\n\n", drive.VendorName, drive.ProductInquiryData, drive.RevisionNumber, strings.Trim(string(drive.RevionDate[:]), "\u0000"), opt.ReadOffset, sectorOrder, opt.ReadAtOnce)
}
