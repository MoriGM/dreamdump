package cd

import (
	"dreamdump/scsi"
)

const (
	MSF_FRAME  = 1
	MSF_SECOND = MSF_FRAME * 75
	MSF_MINUTE = MSF_SECOND * 60
)

type (
	CdSectorData       [scsi.SECTOR_DATA_SIZE]uint8
	CdSectorC2         [scsi.SECTOR_C2_SIZE]uint8
	CdSectorSubchannel [scsi.SECTOR_SUB_SIZE]uint8
)

type Sector struct {
	Data CdSectorData
	C2   CdSectorC2
	Sub  CdSectorSubchannel
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
	LBA int32
}
