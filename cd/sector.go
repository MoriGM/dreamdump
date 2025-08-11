package cd

import "dreamdump/scsi"

type Sector struct {
	Data [scsi.SECTOR_DATA_SIZE]uint8
	C2   [scsi.SECTOR_C2_SIZE]uint8
	Sub  [scsi.SECTOR_SUB_SIZE]uint8
}
