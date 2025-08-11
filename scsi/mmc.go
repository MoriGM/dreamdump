package scsi

const (
	MMC_READ_CD = 0xbe
)

const (
	EXPECTED_SECTOR_TYPE_AUTO = 0x00
	EXPECTED_SECTOR_TYPE_CDDA = 0x04
)

const (
	SUBCODE_RAW = 0x01
	SUBCODE_Q   = 0x02
	SUBCODE_RW  = 0x04
)
