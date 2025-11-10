package cd

import (
	"encoding/binary"
	"fmt"

	"github.com/pasztorpisti/go-crc"
)

const (
	USER_DATA_MODE1_END       = 2064
	USER_DATA_MODE2_FORM1_END = 2072
	USER_DATA_MODE2_FORM2_END = 2348
	EDC_MODE1_START           = 2064
	EDC_MODE1_END             = 2068
	EDC_MODE2_FORM1_START     = 2072
	EDC_MODE2_FORM1_END       = 2076
	EDC_MODE2_FORM2_START     = 2348
	EDC_MODE2_FORM2_END       = 2352
)

func (data *CdSectorData) GetEDC() [4]byte {
	var edc [4]byte

	if data.GetDataMode() == TRACK_TYPE_DATA_MODE1 {
		copy(edc[:], data[EDC_MODE1_START:EDC_MODE1_END])
	} else if data.GetDataMode() == TRACK_TYPE_DATA_MODE2 && data.GetDataModeForm() == TRACK_TYPE_DATA_MODE2_FORM1 {
		copy(edc[:], data[EDC_MODE2_FORM1_START:EDC_MODE2_FORM1_END])
	} else if data.GetDataMode() == TRACK_TYPE_DATA_MODE2 && data.GetDataModeForm() == TRACK_TYPE_DATA_MODE2_FORM2 {
		copy(edc[:], data[EDC_MODE2_FORM2_START:EDC_MODE2_FORM2_END])
	}

	return edc
}

func (data *CdSectorData) HasEDC() bool {
	edc := data.GetEDC()

	if data.GetDataMode() == TRACK_TYPE_DATA_MODE2 && data.GetDataModeForm() == TRACK_TYPE_DATA_MODE2_FORM2 && edc[0] == 0 && edc[1] == 0 && edc[2] == 0 && edc[3] == 0 {
		return false
	}
	return true
}

func (data *CdSectorData) CheckEDC() bool {
	edc := data.GetEDC()

	if data.HasEDC() {
		return true
	}

	crc32CdRom := crc.CRC32CDROMEDC
	checksum := uint32(0)
	if data.GetDataMode() == TRACK_TYPE_DATA_MODE1 {
		checksum = crc32CdRom.Calc(data[0:USER_DATA_MODE1_END])
	} else if data.GetDataMode() == TRACK_TYPE_DATA_MODE2 && data.GetDataModeForm() == TRACK_TYPE_DATA_MODE2_FORM1 {
		checksum = crc32CdRom.Calc(data[16:USER_DATA_MODE2_FORM1_END])
	} else if data.GetDataMode() == TRACK_TYPE_DATA_MODE2 && data.GetDataModeForm() == TRACK_TYPE_DATA_MODE2_FORM2 {
		checksum = crc32CdRom.Calc(data[16:USER_DATA_MODE2_FORM2_END])
	}

	fmt.Println(checksum)
	fmt.Println(binary.LittleEndian.Uint32(edc[:]))
	return checksum == binary.LittleEndian.Uint32(edc[:])
}
