package cd

import (
	"bytes"

	"dreamdump/encoding/bcd"
	"dreamdump/encoding/msf"
	"dreamdump/option"
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

func (sector CdSectorData) HasSyncHeader() bool {
	syncFrame := []uint8{0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00}
	for i, magicNumber := range syncFrame {
		if magicNumber != sector[i] {
			return false
		}
	}
	return true
}

func (dense Dense) NewOffsetManager(lba int32) *OffsetManager {
	syncFrame := []uint8{0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00}
	syncByteIndex := bytes.Index(dense, syncFrame)
	if syncByteIndex == -1 {
		return &OffsetManager{
			SyncByteOffset: 0,
			ByteOffset:     0,
			SampleOffset:   0,
		}
	}

	zeroSectorCount := int32(0)
	syncByteOffset := int32(syncByteIndex)
	if syncByteOffset >= scsi.SECTOR_DATA_SIZE {
		zeroSectorCount = syncByteOffset / int32(scsi.SECTOR_DATA_SIZE)
		syncByteOffset = syncByteOffset % scsi.SECTOR_DATA_SIZE
	}

	correctSector := CdSectorData(dense[syncByteIndex : syncByteIndex+scsi.SECTOR_DATA_SIZE])
	correctSector.Descramble()
	minute := uint64(bcd.ToUint8(correctSector[SECTOR_DATA_MINUTE]))
	second := uint64(bcd.ToUint8(correctSector[SECTOR_DATA_SECOND]))
	frame := uint64(bcd.ToUint8(correctSector[SECTOR_DATA_FRAME]))
	dataFrameLBA := ((minute * msf.MSF_MINUTE) + (second * msf.MSF_SECOND) + (frame * msf.MSF_FRAME)) - 150
	dataFrameLBA -= uint64(zeroSectorCount)

	byteOffset := ((uint64(lba) - dataFrameLBA) * scsi.SECTOR_DATA_SIZE)
	byteOffset += uint64(syncByteOffset)

	sampleOffset := byteOffset / 4

	offsetManager := &OffsetManager{
		SyncByteOffset: uint32(syncByteOffset),
		ByteOffset:     int32(byteOffset),
		SampleOffset:   int32(sampleOffset),
	}

	return offsetManager
}

func (dense *Dense) GetLBA(offsetManager *OffsetManager, lba int32) *CdSectorData {
	lbaStartSize := (lba-option.DC_START)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
	lbaEndSize := ((lba+1)-option.DC_START)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
	var cdSectorData CdSectorData
	copy(cdSectorData[0:scsi.SECTOR_DATA_SIZE], (*dense)[lbaStartSize:lbaEndSize])
	return &cdSectorData
}
