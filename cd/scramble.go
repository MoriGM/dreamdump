package cd

import (
	"dreamdump/scsi"
)

var DescrambleTable [scsi.SECTOR_DATA_SIZE]uint8

func init() {
	shiftRegister := uint16(1)

	for i := scsi.SECTOR_SYNC_SIZE; i < scsi.SECTOR_DATA_SIZE; i++ {
		DescrambleTable[i] = uint8(shiftRegister)
		for bit := 0; bit < 8; bit++ {
			carry := uint16((shiftRegister & 1) ^ (shiftRegister >> 1 & 1))
			shiftRegister = (carry<<15 | shiftRegister) >> 1
		}
	}
}

func (sec *Sector) Descramble() bool {
	if len(sec.Data) != scsi.SECTOR_DATA_SIZE {
		return false
	}

	for i := scsi.SECTOR_SYNC_SIZE; i < scsi.SECTOR_DATA_SIZE; i++ {
		sec.Data[i] ^= DescrambleTable[i]
	}

	return true
}
