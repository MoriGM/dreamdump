package scsi

import (
	"fmt"
	"os"

	"dreamdump/sgio"
)

func byteFromInt(number int, part byte) byte {
	return (byte)(((number) >> (part * 8)) & 0xFF)
}

func CommandReadCd(dvdDriveDeviceFile *os.File, sector int) (sgio.SgIoHdr, []byte, []byte) {
	block := make([]byte, SECTOR_DATA_C2_SUB_SIZE)
	senseBuf := make([]byte, sgio.SENSE_BUF_LEN)
	cmdBlk := []byte{MMC_READ_CD, EXPECTED_SECTOR_TYPE_CDDA, byteFromInt(sector, 3), byteFromInt(sector, 2), byteFromInt(sector, 1), byteFromInt(sector, 0), 0, 0, 1, 0xfa, 0x02, 0x00}
	fmt.Printf("\rSector: %d", sector)
	sg_io_hdr := sgio.SgIoHdr{
		InterfaceID:    int32('S'),
		CmdLen:         uint8(len(cmdBlk)),
		MxSbLen:        sgio.SENSE_BUF_LEN,
		DxferLen:       SECTOR_DATA_C2_SUB_SIZE,
		DxferDirection: sgio.SG_DXFER_FROM_DEV,
		Cmdp:           &cmdBlk[0],
		Sbp:            &senseBuf[0],
		Dxferp:         &block[0],
		Timeout:        sgio.TIMEOUT_20_SECS,
	}
	sgio.SgioSyscall(dvdDriveDeviceFile, &sg_io_hdr)

	return sg_io_hdr, senseBuf, block
}
