package cd

import (
	"bytes"
	"slices"

	"dreamdump/encoding/bcd"
	"dreamdump/option"
	"dreamdump/scsi"
)

type OffsetManager struct {
	SyncByteOffset    uint32
	DataFrameLBA      int32
	ByteOffset        int32
	ByteOffsetFixed   int32
	SampleOffset      int32
	SampleOffsetFixed int32
}

func GetWriteOffset(opt *option.Option, lba int32, sectors []*Sector) int32 {
	for i := range len(sectors) - 1 {
		sector := sectors[i]
		sectorNext := sectors[i+1]
		syncSize := NewDataSync(opt, lba+int32(i), slices.Concat(sector.Data[:], sectorNext.Data[:]))
		if syncSize == nil {
			continue
		}
	}
	return 0
}

func NewDataSync(opt *option.Option, lba int32, twoSectors []uint8) *OffsetManager {
	syncFrame := []uint8{0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00}
	syncByteOffset := int32(bytes.Index(twoSectors, syncFrame))
	if syncByteOffset == -1 {
		return nil
	}

	correctSector := CdSectorData(twoSectors[syncByteOffset : syncByteOffset+scsi.SECTOR_DATA_SIZE])
	correctSector.Descramble()
	minute := uint64(bcd.ToUint8(correctSector[12]))
	second := uint64(bcd.ToUint8(correctSector[13]))
	frame := uint64(bcd.ToUint8(correctSector[14]))
	dataFrameLBA := ((minute * MSF_MINUTE) + (second * MSF_SECOND) + (frame * MSF_FRAME)) - 150

	byteOffset := ((lba - int32(dataFrameLBA)) * scsi.SECTOR_DATA_SIZE)
	byteOffset += syncByteOffset

	sampleOffset := byteOffset / 4
	byteOffsetFixed := byteOffset - (int32(opt.ReadOffset) * 4)
	sampleOffsetFixed := byteOffsetFixed / 4

	offsetManager := &OffsetManager{
		SyncByteOffset:    uint32(syncByteOffset),
		ByteOffset:        byteOffset,
		ByteOffsetFixed:   byteOffsetFixed,
		SampleOffset:      sampleOffset,
		SampleOffsetFixed: sampleOffsetFixed,
	}

	return offsetManager
}
