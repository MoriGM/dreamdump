package cd

import (
	"bytes"

	"dreamdump/encoding/bcd"
	"dreamdump/scsi"
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
	minute := uint64(bcd.ToUint8(correctSector[12]))
	second := uint64(bcd.ToUint8(correctSector[13]))
	frame := uint64(bcd.ToUint8(correctSector[14]))
	dataFrameLBA := ((minute * MSF_MINUTE) + (second * MSF_SECOND) + (frame * MSF_FRAME)) - 150

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
