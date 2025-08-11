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
