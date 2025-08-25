package cbd

const (
	ReadCD_SECTOR_TYPE_AUTO = 0x00
	ReadCD_SECTOR_TYPE_CDDA = 0x04
)

const (
	ReadCD_SUBCODE_NO  = 0x00
	ReadCD_SUBCODE_RAW = 0x01
	ReadCD_SUBCODE_Q   = 0x02
	ReadCD_SUBCODE_RW  = 0x04
)

const (
	ReadCD_C2_ERROR_FLAG       = 0x02
	ReadCD_C2_BLOCK_ERROR_FLAG = 0x04
)

const (
	ReadCD_USER_DATA   = 0x10
	ReadCD_HEADER      = 0x20
	ReadCD_SYNC        = 0x80
	ReadCD_SYNC_HEADER = 0xE0
	ReadCD_ALL         = 0xF8
)

type ReadCD struct {
	OperationCode      uint8
	ExpectedSectorType uint8
	LBA                int32
	MSBTransferLength  uint8
	TransferLength     uint16
	FlagBits           uint8
	Subchannel         uint8
	Reserved           uint8
}

type Inquiry struct {
	OperationCode  uint8
	Reserved       uint8
	PageCode       uint8
	TransferLength uint16
	Control        uint8
}
