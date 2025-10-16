package cd

import (
	"dreamdump/scsi"
)

type (
	CdSectorData [scsi.SECTOR_DATA_SIZE]byte
	CdSectorC2   [scsi.SECTOR_C2_SIZE]byte
	QChannel     [12]byte
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
	TRACK_TYPE_DATA             uint8 = 0x04
	TRACK_TYPE_AUDIO            uint8 = 0x00
	TRACK_TYPE_DATA_MODE1       uint8 = 0x01
	TRACK_TYPE_DATA_MODE2       uint8 = 0x02
	TRACK_TYPE_DATA_MODE2_FORM1 uint8 = 0x3
	TRACK_TYPE_DATA_MODE2_FORM2 uint8 = 0x4
)

type Track struct {
	Type         uint8
	Lba          int32
	LbaEnd       int32
	TrackNumber  uint8
	IndexNumbers []uint8
	Indexs       map[uint8]*Index
}

type Index struct {
	Lba int32
}

type TrackMeta struct {
	TrackNumber        uint8
	FileName           string
	CRC32              uint32
	MD5                [16]byte
	SHA1               [20]byte
	Size               uint32
	Sectors            uint32
	DataType           uint8
	InvalidSyncSectors uint32
}
