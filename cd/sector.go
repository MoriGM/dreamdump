package cd

import (
	"dreamdump/scsi"

	bcd "github.com/johnsonjh/gobcd"
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

func (sec Sector) GetSubcodeLocationLBA() int32 {
	return (int32(bcd.ToUint8(sec.Sub[3])) * MSF_MINUTE) + (int32(bcd.ToUint8(sec.Sub[4])) * MSF_SECOND) + (int32(bcd.ToUint8(sec.Sub[5])))
}
