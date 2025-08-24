package cd

import (
	crc16_ccitt "dreamdump/crc/CRC16_CCITT"
	"dreamdump/encoding/bcd"
	"dreamdump/scsi"
)

const (
	SUBCHANNEL_P = 7
	SUBCHANNEL_Q = 6
)

func (sub *Subchannel) Parse(subcodes [scsi.SECTOR_SUB_SIZE]uint8) {
	for i, subcode := range subcodes {
		sub.Pchannel[(i-(i%8))/8] |= ((subcode >> SUBCHANNEL_P) & 0x01) << (7 - (i % 8))
		sub.Qchannel[(i-(i%8))/8] |= ((subcode >> SUBCHANNEL_Q) & 0x01) << (7 - (i % 8))
	}
}

func (sub *QChannel) LBA() int32 {
	return sub.AbsolutLBA() - (MSF_SECOND * 2)
}

func (sub *QChannel) AbsolutLBA() int32 {
	return (int32(bcd.ToUint8(sub[7])) * MSF_MINUTE) + (int32(bcd.ToUint8(sub[8])) * MSF_SECOND) + (int32(bcd.ToUint8(sub[9])))
}

func (sub *QChannel) TrackNumber() uint8 {
	return bcd.ToUint8(sub[1])
}

func (sub *QChannel) IndexNumber() uint8 {
	return bcd.ToUint8(sub[2])
}

func (sub *QChannel) TrackType() uint8 {
	if (sub[0] & 0b01000000) == 0 {
		return TRACK_TYPE_AUDIO
	}

	return TRACK_TYPE_DATA
}

func (sub *QChannel) CheckParity() bool {
	return crc16_ccitt.Calculate(sub[0:10]) == ((uint16(sub[10]) << 8) | uint16(sub[11]))
}
