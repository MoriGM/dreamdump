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
	CdSectorData [scsi.SECTOR_DATA_SIZE]uint8
	CdSectorC2   [scsi.SECTOR_C2_SIZE]uint8
	QChannel     [12]uint8
)

type Subchannel struct {
	Qchannel QChannel
}

type Sector struct {
	Data CdSectorData
	C2   CdSectorC2
	Sub  Subchannel
}

const (
	TRACK_TYPE_DATA  uint8 = 0x04
	TRACK_TYPE_AUDIO uint8 = 0x01
)

type Track struct {
	Type        uint8
	Lba         int32
	LbaEnd      int32
	TrackNumber uint8
	Indexs      map[uint8]*Index
}

type Index struct {
	Lba int32
}

type TrackMeta struct {
	FileName string
	CRC32    uint32
	MD5      [16]byte
	SHA1     [20]byte
	Size     uint32
	Sectors  uint32
}
