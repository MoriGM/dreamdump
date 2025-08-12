package cd

import (
	"dreamdump/scsi"
)

const (
	MSF_FRAME  = 1
	MSF_SECOND = MSF_FRAME * 75
	MSF_MINUTE = MSF_SECOND * 60
)

type Sector struct {
	Data [scsi.SECTOR_DATA_SIZE]uint8
	C2   [scsi.SECTOR_C2_SIZE]uint8
	Sub  [scsi.SECTOR_SUB_SIZE]uint8
}

const (
	TRACK_TYPE_DATA  uint8 = 0x04
	TRACK_TYPE_AUDIO uint8 = 0x01
)

type Track struct {
	Type   uint8
	LBA    int32
	Indexs map[uint8]Index
}

type Index struct {
	Number int8
	LBA    int32
}
