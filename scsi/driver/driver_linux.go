//go:build linux

package driver

import (
	"bytes"
	"encoding/binary"
	"os"

	"dreamdump/scsi/driver/sgio"
)

func OpenScsiDevice(fname string) (any, error) {
	return sgio.OpenScsiDevice(fname)
}

func Read(fileHandle any, cmd interface{}, size uint32) Status {
	driveDeviceFile, ok := fileHandle.(*os.File)
	if !ok {
		panic("Error while casting")
	}
	var cmdBlk bytes.Buffer
	err := binary.Write(&cmdBlk, binary.LittleEndian, cmd)
	if err != nil {
		panic(err)
	}

	block := make([]byte, size)
	senseBuf := make([]byte, sgio.SENSE_BUF_LEN)
	sg_io_hdr := sgio.SgIoHdr{
		InterfaceID:    int32('S'),
		CmdLen:         uint8(cmdBlk.Len()),
		MxSbLen:        sgio.SENSE_BUF_LEN,
		DxferLen:       size,
		DxferDirection: sgio.SG_DXFER_FROM_DEV,
		Cmdp:           &cmdBlk.Bytes()[0],
		Sbp:            &senseBuf[0],
		Timeout:        sgio.TIMEOUT_20_SECS,
	}

	if size > 0 {
		sg_io_hdr.Dxferp = &block[0]
	}

	err = sgio.SgioSyscall(driveDeviceFile, &sg_io_hdr)
	if err != nil {
		panic(err)
	}

	status := Status{
		Status: sg_io_hdr.Status,
		Key:    senseBuf[2] & 0x0F,
		Asc:    senseBuf[12],
		AscQ:   senseBuf[13],
		Block:  block,
	}

	return status
}
