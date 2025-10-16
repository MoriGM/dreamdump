package drive

import (
	"dreamdump/log"
	"dreamdump/option"
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
	log.Printf("Drive: %s - %s (revision level: %s, vendor specific: %s) Read Offset: %d  Sector Order: %s  Read at once: %d\n", drive.VendorName, drive.ProductInquiryData, drive.RevisionNumber, drive.RevionDate, opt.ReadOffset, sectorOrder, opt.ReadAtOnce)
}
