package cd

import (
	"bytes"

	"dreamdump/encoding/bcd"
	"dreamdump/encoding/msf"
	"dreamdump/scsi"
)

const (
	SECTOR_DATA_MINUTE = 12
	SECTOR_DATA_SECOND = 13
	SECTOR_DATA_FRAME  = 14
)

type OffsetManager struct {
	SyncByteOffset uint32
	DataFrameLBA   int32
	ByteOffset     int32
	SampleOffset   int32
}

func (dense Dense) NewOffsetManager(lba int32) *OffsetManager {
	syncFrame := []uint8{0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00}
	syncByteOffset := int32(bytes.Index(dense, syncFrame))
	if syncByteOffset == -1 {
		return &OffsetManager{
			SyncByteOffset: 0,
			ByteOffset:     0,
			SampleOffset:   0,
		}
	}

	correctSector := CdSectorData(dense[syncByteOffset : syncByteOffset+scsi.SECTOR_DATA_SIZE])
	correctSector.Descramble()
	minute := uint64(bcd.ToUint8(correctSector[SECTOR_DATA_MINUTE]))
	second := uint64(bcd.ToUint8(correctSector[SECTOR_DATA_SECOND]))
	frame := uint64(bcd.ToUint8(correctSector[SECTOR_DATA_FRAME]))
	dataFrameLBA := ((minute * msf.MSF_MINUTE) + (second * msf.MSF_SECOND) + (frame * msf.MSF_FRAME)) - 150

	byteOffset := ((lba - int32(dataFrameLBA)) * scsi.SECTOR_DATA_SIZE)
	byteOffset += syncByteOffset

	sampleOffset := byteOffset / 4

	offsetManager := &OffsetManager{
		SyncByteOffset: uint32(syncByteOffset),
		ByteOffset:     byteOffset,
		SampleOffset:   sampleOffset,
	}

	return offsetManager
}
