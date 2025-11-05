package cd

import (
	"encoding/binary"

	"github.com/pasztorpisti/go-crc"
)

const (
	USER_DATA_MODE1_START = 16
	USER_DATA_MODE1_END   = 2064
	EDC_MODE1_START       = 2064
	EDC_MODE1_END         = 2068
)

func (data *CdSectorData) GetEDC() [4]byte {
	var edc [4]byte

	if data.GetDataMode() == TRACK_TYPE_DATA_MODE1 {
		copy(edc[:], data[EDC_MODE1_START:EDC_MODE1_END])
	}

	return edc
}

func (data *CdSectorData) GetUserData() []byte {
	if data.GetDataMode() == TRACK_TYPE_DATA_MODE1 {
		return data[USER_DATA_MODE1_START:USER_DATA_MODE1_END]
	}
	return []byte{}
}

func (data *CdSectorData) CheckEDC() bool {
	edc := data.GetEDC()

	crc32CdRom := crc.CRC32CDROMEDC
	checksum := uint32(0)
	if data.GetDataMode() == TRACK_TYPE_DATA_MODE1 {
		checksum = crc32CdRom.Calc(data[0:USER_DATA_MODE1_END])
	}

	return checksum == binary.LittleEndian.Uint32(edc[:])
}
