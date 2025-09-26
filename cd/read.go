package cd

import (
	"dreamdump/option"
	"dreamdump/scsi"
	"dreamdump/scsi/driver"
	"dreamdump/scsi/scsi_commands"
	"errors"
)

func ReadSectors(opt *option.Option, lba int32, readAtOnce uint8) ([]*Sector, error) {
	status := scsi_commands.ReadCd(opt, lba, readAtOnce)
	err := driver.CheckSense(&status)
	if err != nil {
		return nil, err
	}
	if status.Status != 0 {
		return nil, errors.New("scsi Error")
	}

	return ConvertRawToSectors(opt, status.Block, readAtOnce), nil
}

func ConvertRawToSectors(opt *option.Option, block []uint8, readAtOnce uint8) []*Sector {
	sectors := make([]*Sector, readAtOnce)
	blockSize := option.DATA

	if opt.SectorOrder == option.DATA_C2 {
		blockSize = scsi.SECTOR_DATA_C2_SIZE
	}

	if opt.SectorOrder == option.DATA_SUB {
		blockSize = scsi.SECTOR_DATA_SUB_SIZE
	}

	if opt.SectorOrder == option.DATA_C2_SUB || opt.SectorOrder == option.DATA_SUB_C2 {
		blockSize = scsi.SECTOR_DATA_C2_SUB_SIZE
	}

	for i := range readAtOnce {
		start := int(i) * blockSize
		end := (int(i) + 1) * blockSize
		sectorBlock := block[start:end]
		sectors[i] = ConvertRawToSector(opt, sectorBlock)
	}

	return sectors
}

func ConvertRawToSector(opt *option.Option, block []uint8) *Sector {
	sectorContent := new(Sector)
	sectorContent.Data = [scsi.SECTOR_DATA_SIZE]uint8(block[0:scsi.SECTOR_DATA_SIZE])
	sectorContent.C2 = [scsi.SECTOR_C2_SIZE]uint8{}
	sectorContent.Sub = Subchannel{}

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
