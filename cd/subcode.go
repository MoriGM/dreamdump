package cd

import (
	bcd "github.com/johnsonjh/gobcd"
)

func (sec Sector) SubcodeLBA() int32 {
	return sec.SubcodeAbsolutLBA() - (MSF_SECOND * 2)
}

func (sec Sector) SubcodeAbsolutLBA() int32 {
	return (int32(bcd.ToUint8(sec.Sub[7])) * MSF_MINUTE) + (int32(bcd.ToUint8(sec.Sub[8])) * MSF_SECOND) + (int32(bcd.ToUint8(sec.Sub[9])))
}

func (sec Sector) SubcodeTrackNumber() uint8 {
	return bcd.ToUint8(sec.Sub[1])
}

func (sec Sector) SubcodeIndexNumber() uint8 {
	return bcd.ToUint8(sec.Sub[2])
}

func (sec Sector) SubcodeTrackType() int {
	if (sec.Sub[0] & 0b01000000) == 0 {
		return TRACK_TYPE_AUDIO
	}

	return TRACK_TYPE_DATA
}
