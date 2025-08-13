package cd

import (
	"dreamdump/option"
	"dreamdump/scsi"
	"dreamdump/sgio"
)

func ReadSector(option option.Option, sector int32) (Sector, error) {
	sg_io_hdr, senseBuf, block := scsi.CommandReadCd(option.Drive, sector)
	err := sgio.CheckSense(&sg_io_hdr, &senseBuf)
	if err != nil {
		return Sector{}, err
	}

	return ConvertRawToSector(option, block), nil
}

func ConvertRawToSector(option option.Option, block []uint8) Sector {
	sectorContent := Sector{
		Data: [scsi.SECTOR_DATA_SIZE]uint8(block[0:scsi.SECTOR_DATA_SIZE]),
		C2:   [scsi.SECTOR_C2_SIZE]uint8{},
		Sub:  [scsi.SECTOR_SUB_SIZE]uint8{},
	}

	if option.SectorOrder == scsi.DATA_C2 {
		sectorContent.C2 = [scsi.SECTOR_C2_SIZE]uint8(block[scsi.SECTOR_DATA_SIZE:scsi.SECTOR_DATA_C2_SIZE])
	}

	if option.SectorOrder == scsi.DATA_SUB {
		sectorContent.Sub = [scsi.SECTOR_SUB_SIZE]uint8(block[scsi.SECTOR_DATA_SIZE:scsi.SECTOR_DATA_SUB_SIZE])
	}

	if option.SectorOrder == scsi.DATA_SUB_C2 {
		sectorContent.Sub = [scsi.SECTOR_SUB_SIZE]uint8(block[scsi.SECTOR_DATA_SIZE:scsi.SECTOR_DATA_SUB_SIZE])
		sectorContent.C2 = [scsi.SECTOR_C2_SIZE]uint8(block[scsi.SECTOR_DATA_SUB_SIZE:scsi.SECTOR_DATA_SUB_C2_SIZE])
	}

	if option.SectorOrder == scsi.DATA_C2_SUB {
		sectorContent.C2 = [scsi.SECTOR_C2_SIZE]uint8(block[scsi.SECTOR_DATA_SIZE:scsi.SECTOR_DATA_C2_SIZE])
		sectorContent.Sub = [scsi.SECTOR_SUB_SIZE]uint8(block[scsi.SECTOR_DATA_C2_SIZE:scsi.SECTOR_DATA_C2_SUB_SIZE])
	}

	return sectorContent
}
