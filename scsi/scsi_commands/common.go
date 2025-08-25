package scsi_commands

import (
	"bytes"
	"encoding/binary"

	"dreamdump/drive"
	bigendian "dreamdump/encoding/big_endian"
	"dreamdump/option"
	"dreamdump/scsi"
	"dreamdump/scsi/cbd"
	"dreamdump/sgio"
)

func Inquiry(opt *option.Option) *drive.Drive {
	size := uint16(0x60)
	inquiryCommand := cbd.Inquiry{
		OperationCode:  scsi.COMMON_INQUIRY,
		TransferLength: bigendian.Uint16(size),
	}

	sg_io_hdr, senseBuf, block := scsi.Read(opt.Drive, inquiryCommand, size)
	err := sgio.CheckSense(&sg_io_hdr, &senseBuf)
	if err != nil {
		panic(err)
	}

	drive := new(drive.Drive)
	buf := bytes.NewReader(block[8:46])
	err = binary.Read(buf, binary.LittleEndian, drive)
	if err != nil {
		panic(err)
	}

	return drive
}
