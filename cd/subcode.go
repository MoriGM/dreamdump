package cd

import (
	"dreamdump/encoding/bcd"
)

func (sub *CdSectorSubchannel) LBA() int32 {
	return sub.AbsolutLBA() - (MSF_SECOND * 2)
}

func (sub *CdSectorSubchannel) AbsolutLBA() int32 {
	return (int32(bcd.ToUint8(sub[7])) * MSF_MINUTE) + (int32(bcd.ToUint8(sub[8])) * MSF_SECOND) + (int32(bcd.ToUint8(sub[9])))
}

func (sub *CdSectorSubchannel) TrackNumber() uint8 {
	return bcd.ToUint8(sub[1])
}

func (sub *CdSectorSubchannel) IndexNumber() uint8 {
	return bcd.ToUint8(sub[2])
}

func (sub *CdSectorSubchannel) TrackType() uint8 {
	if (sub[0] & 0b01000000) == 0 {
		return TRACK_TYPE_AUDIO
	}

	return TRACK_TYPE_DATA
}
