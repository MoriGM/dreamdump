package scsi_commands

import (
	bigendian "dreamdump/encoding/big_endian"
	"dreamdump/option"
	"dreamdump/scsi"
	"dreamdump/scsi/cbd"
	"dreamdump/scsi/driver"
)

func ReadCd(opt *option.Option, lba int32) driver.Status {
	size := uint16(scsi.SECTOR_DATA_SIZE)

	readCdCommand := cbd.ReadCD{
		OperationCode:      scsi.MMC_READ_CD,
		ExpectedSectorType: cbd.ReadCD_SECTOR_TYPE_CDDA,
		LBA:                bigendian.Int32(lba),
		MSBTransferLength:  0,
		TransferLength:     bigendian.Uint16(1),
		FlagBits:           cbd.ReadCD_ALL,
		Subchannel:         cbd.ReadCD_SUBCODE_NO,
	}

	if opt.SectorOrder == option.DATA_C2 {
		size = scsi.SECTOR_DATA_C2_SIZE
		readCdCommand.FlagBits |= cbd.ReadCD_C2_ERROR_FLAG
	}

	if opt.SectorOrder == option.DATA_SUB {
		size = scsi.SECTOR_DATA_SUB_SIZE + scsi.SECTOR_PAD_SIZE
		readCdCommand.Subchannel = cbd.ReadCD_SUBCODE_RAW
	}

	if opt.SectorOrder == option.DATA_C2_SUB || opt.SectorOrder == option.DATA_SUB_C2 {
		size = scsi.SECTOR_DATA_C2_SUB_SIZE + scsi.SECTOR_PAD_SIZE
		readCdCommand.FlagBits |= cbd.ReadCD_C2_ERROR_FLAG
		readCdCommand.Subchannel = cbd.ReadCD_SUBCODE_RAW
	}

	return driver.Read(opt.Drive, readCdCommand, size)
}

func SetCDSpeed(opt *option.Option) driver.Status {
	readCdCommand := cbd.Speed{
		OperationCode: scsi.MMC_READ_CD,
		ReadSpeed:     bigendian.Uint16(opt.Speed),
	}

	return driver.Read(opt.Drive, readCdCommand, 0)
}
