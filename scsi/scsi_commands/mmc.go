package scsi_commands

import (
	bigendian "dreamdump/encoding/big_endian"
	"dreamdump/option"
	"dreamdump/scsi"
	"dreamdump/scsi/cbd"
	"dreamdump/scsi/driver"
)

func ReadCd(opt *option.Option, lba int32, readAtOnce uint8) driver.Status {
	size := scsi.SECTOR_DATA_SIZE

	readCdCommand := cbd.ReadCD{
		OperationCode:      scsi.MMC_READ_CD,
		ExpectedSectorType: cbd.ReadCD_SECTOR_TYPE_CDDA,
		LBA:                bigendian.Int32(lba),
		MSBTransferLength:  0,
		TransferLength:     bigendian.Uint16(uint16(readAtOnce)),
		FlagBits:           cbd.ReadCD_ALL,
		Subchannel:         cbd.ReadCD_SUBCODE_NO,
	}

	if opt.SectorOrder == option.DATA_C2 {
		size = scsi.SECTOR_DATA_C2_SIZE
		readCdCommand.FlagBits |= cbd.ReadCD_C2_ERROR_FLAG
	}

	if opt.SectorOrder == option.DATA_SUB {
		size = scsi.SECTOR_DATA_SUB_SIZE
		readCdCommand.Subchannel = cbd.ReadCD_SUBCODE_RAW
	}

	if opt.SectorOrder == option.DATA_C2_SUB || opt.SectorOrder == option.DATA_SUB_C2 {
		size = scsi.SECTOR_DATA_C2_SUB_SIZE
		readCdCommand.FlagBits |= cbd.ReadCD_C2_ERROR_FLAG
		readCdCommand.Subchannel = cbd.ReadCD_SUBCODE_RAW
	}

	size *= int(readAtOnce)
	if (size & 3) > 0 {
		size += 4 - (size & 3)
	}

	return driver.Read(opt.Drive, readCdCommand, uint32(size))
}

func SetCDSpeed(opt *option.Option) driver.Status {
	readCdCommand := cbd.Speed{
		OperationCode: scsi.MMC_SET_CD_SPEED,
		ReadSpeed:     bigendian.Uint16(opt.Speed),
	}

	return driver.Read(opt.Drive, readCdCommand, 0)
}
