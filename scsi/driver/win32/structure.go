//go:build windows

package win32

const (
	SCSI_PASS_THROUGH_DIRECT = 0x4D014
	SCSI_IOCTL_DATA_OUT      = 0
	SCSI_IOCTL_DATA_IN       = 1
)

type SENSE_DATA struct {
	ErroCode                     uint8
	SegmentNumber                uint8
	SenseKey                     uint8
	Information                  [4]uint8
	AdditionalSenseLength        uint8
	CommandSpecificInformation   [4]uint8
	AdditionalSenseCode          uint8
	AdditionalSenseCodeQualifier uint8
	FieldReplaceableUnitCode     uint8
	SenseKeySpecific             [3]uint8
	AdditionalSenseBytes         *uint8
}

type SCSI_PASS_THROUGH_DIRECT_STRUCT struct {
	Length             uint16
	ScsiStatus         uint8
	PathId             uint8
	TargetId           uint8
	Lun                uint8
	CdbLength          uint8
	SenseInfoLength    uint8
	DataIn             uint8
	DataTransferLength uint32
	TimeOutValue       uint32
	DataBuffer         *byte
	SenseInfoOffset    uint32
	Cdb                [16]uint8
}

type SPTD_SD struct {
	SPTD SCSI_PASS_THROUGH_DIRECT_STRUCT
	SD   SENSE_DATA
}
