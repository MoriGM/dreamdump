//go:build windows

package driver

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"dreamdump/scsi/driver/win32"

	"golang.org/x/sys/windows"
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

	sptd_sd := &win32.SPTD_SD{
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
	sptd_sd.SPTD.Length = uint16(unsafe.Sizeof(sptd_sd.SPTD))
	sptd_sd.SPTD.SenseInfoLength = uint8(unsafe.Sizeof(sptd_sd.SD))
	sptd_sd.SPTD.SenseInfoOffset = uint32(unsafe.Offsetof(sptd_sd.SD))
	if size > 0 {
		sptd_sd.SPTD.DataBuffer = &block[0]
	}

	err = win32.ScsiCall(driveDeviceFile, sptd_sd)
	if err != nil {
		return Status{
			Status: 0xFF,
			Key:    0xFF,
			Asc:    0xFF,
			AscQ:   0xFF,
		}
	}

	status := Status{
		Status: sptd_sd.SPTD.ScsiStatus,
		Key:    sptd_sd.SD.SenseKey,
		Asc:    sptd_sd.SD.AdditionalSenseCode,
		AscQ:   sptd_sd.SD.AdditionalSenseCodeQualifier,
		Block:  block,
	}
	return status
}
