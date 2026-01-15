package scsi_commands

import (
	"bytes"
	"encoding/binary"

	"dreamdump/drive"
	bigendian "dreamdump/encoding/big_endian"
	"dreamdump/option"
	"dreamdump/scsi"
	"dreamdump/scsi/cbd"
	"dreamdump/scsi/driver"
)

func Inquiry(opt *option.Option) *drive.Drive {
	size := uint16(0x60)
	command := cbd.Inquiry{
		OperationCode:  scsi.COMMON_INQUIRY,
		TransferLength: bigendian.Uint16(size),
	}

	status := driver.Read(opt.Drive, command, uint32(size))
	err := driver.CheckSense(&status)
	if err != nil {
		panic(err)
	}

	drive := new(drive.Drive)
	buf := bytes.NewReader(status.Block[8:46])
	err = binary.Read(buf, binary.BigEndian, drive)
	if err != nil {
		panic(err)
	}

	return drive
}
