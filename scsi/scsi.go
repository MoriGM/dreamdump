package scsi

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"

	"dreamdump/sgio"
)

func Read(dvdDriveDeviceFile *os.File, cmd interface{}, size uint16) (sgio.SgIoHdr, []byte, []byte) {
	var cmdBlk bytes.Buffer
	cmdBlkEncoder := gob.NewEncoder(&cmdBlk)
	cmdBlkEncoder.Encode(cmd)

	block := make([]byte, size)
	senseBuf := make([]byte, sgio.SENSE_BUF_LEN)
	sg_io_hdr := sgio.SgIoHdr{
		InterfaceID:    int32('S'),
		CmdLen:         uint8(cmdBlk.Len()),
		MxSbLen:        sgio.SENSE_BUF_LEN,
		DxferLen:       uint32(size),
		DxferDirection: sgio.SG_DXFER_FROM_DEV,
		Cmdp:           &cmdBlk.Bytes()[0],
		Sbp:            &senseBuf[0],
		Dxferp:         &block[0],
		Timeout:        sgio.TIMEOUT_20_SECS,
	}
	fmt.Println(sg_io_hdr)
	os.Exit(1)

	sgio.SgioSyscall(dvdDriveDeviceFile, &sg_io_hdr)

	return sg_io_hdr, senseBuf, block
}
