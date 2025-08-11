package cd

import bcd "github.com/johnsonjh/gobcd"

func (sec Sector) GetSubcodeLBA() int32 {
	return (int32(bcd.ToUint8(sec.Sub[3])) * MSF_MINUTE) + (int32(bcd.ToUint8(sec.Sub[4])) * MSF_SECOND) + (int32(bcd.ToUint8(sec.Sub[5])))
}

func (sec Sector) GetSubcodeTrack() uint8 {
	return bcd.ToUint8(sec.Sub[1])
}

func (sec Sector) GetSubcodeIndex() uint8 {
	return bcd.ToUint8(sec.Sub[2])
}
