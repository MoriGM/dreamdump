package cd

import (
	"dreamdump/option"
	"dreamdump/scsi"
	"dreamdump/scsi/commands"
	"dreamdump/sgio"
)

func ReadSector(opt *option.Option, lba int32) (Sector, error) {
	sg_io_hdr, senseBuf, block := commands.ReadCd(opt, lba)
	err := sgio.CheckSense(&sg_io_hdr, &senseBuf)
	if err != nil {
		return Sector{}, err
	}

	return ConvertRawToSector(opt, block), nil
}

func ConvertRawToSector(opt *option.Option, block []uint8) Sector {
	sectorContent := Sector{
		Data: [scsi.SECTOR_DATA_SIZE]uint8(block[0:scsi.SECTOR_DATA_SIZE]),
		C2:   [scsi.SECTOR_C2_SIZE]uint8{},
		Sub:  Subchannel{},
	}

	if opt.SectorOrder == option.DATA_C2 {
		sectorContent.C2 = [scsi.SECTOR_C2_SIZE]uint8(block[scsi.SECTOR_DATA_SIZE:scsi.SECTOR_DATA_C2_SIZE])
	}

	if opt.SectorOrder == option.DATA_SUB {
		sectorContent.Sub.Parse([scsi.SECTOR_SUB_SIZE]uint8(block[scsi.SECTOR_DATA_SIZE:scsi.SECTOR_DATA_SUB_SIZE]))
	}

	if opt.SectorOrder == option.DATA_SUB_C2 {
		sectorContent.Sub.Parse([scsi.SECTOR_SUB_SIZE]uint8(block[scsi.SECTOR_DATA_SIZE:scsi.SECTOR_DATA_SUB_SIZE]))
		sectorContent.C2 = [scsi.SECTOR_C2_SIZE]uint8(block[scsi.SECTOR_DATA_SUB_SIZE:scsi.SECTOR_DATA_SUB_C2_SIZE])
	}

	if opt.SectorOrder == option.DATA_C2_SUB {
		sectorContent.C2 = [scsi.SECTOR_C2_SIZE]uint8(block[scsi.SECTOR_DATA_SIZE:scsi.SECTOR_DATA_C2_SIZE])
		sectorContent.Sub.Parse([scsi.SECTOR_SUB_SIZE]uint8(block[scsi.SECTOR_DATA_C2_SIZE:scsi.SECTOR_DATA_C2_SUB_SIZE]))
	}

	return sectorContent
}
