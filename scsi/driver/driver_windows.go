//go:build windows

package driver

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"dreamdump/scsi/driver/win32"

	"golang.org/x/sys/windows"
)

const (
	MAX_READ_AT_ONCE = 20
)

func OpenScsiDevice(fname string) (any, error) {
	return win32.OpenScsiDevice(fname)
}

func Read(fileHandle any, cmd any, size uint32) Status {
	driveDeviceFile, ok := fileHandle.(windows.Handle)
	if !ok {
		panic("Error while casting")
	}
	var cmdBlk bytes.Buffer
	err := binary.Write(&cmdBlk, binary.LittleEndian, cmd)
	if err != nil {
		panic(err)
	}
	if cmdBlk.Len() < 16 {
		err := binary.Write(&cmdBlk, binary.LittleEndian, make([]byte, 16-cmdBlk.Len()))
		if err != nil {
			panic(err)
		}
	}

	block := make([]byte, size)

	sptdSd := &win32.SPTD_SD{
		SPTD: win32.SCSI_PASS_THROUGH_DIRECT_STRUCT{
			TargetId:           0,
			PathId:             0,
			Lun:                0,
			CdbLength:          uint8(cmdBlk.Len()),
			Cdb:                [16]uint8(cmdBlk.Bytes()),
			DataIn:             win32.SCSI_IOCTL_DATA_IN,
			DataTransferLength: size,
			TimeOutValue:       30,
		},
		SD: win32.SENSE_DATA{},
	}
	sptdSd.SPTD.Length = uint16(unsafe.Sizeof(sptdSd.SPTD))
	sptdSd.SPTD.SenseInfoLength = uint8(unsafe.Sizeof(sptdSd.SD))
	sptdSd.SPTD.SenseInfoOffset = uint32(unsafe.Offsetof(sptdSd.SD))
	if size > 0 {
		sptdSd.SPTD.DataBuffer = &block[0]
	}

	err = win32.ScsiCall(driveDeviceFile, sptdSd)
	if err != nil {
		return Status{
			Status: 0xFF,
			Key:    0xFF,
			Asc:    0xFF,
			AscQ:   0xFF,
		}
	}

	status := Status{
		Status: sptdSd.SPTD.ScsiStatus,
		Key:    sptdSd.SD.SenseKey,
		Asc:    sptdSd.SD.AdditionalSenseCode,
		AscQ:   sptdSd.SD.AdditionalSenseCodeQualifier,
		Block:  block,
	}
	return status
}
