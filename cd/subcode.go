package cd

import (
	"dreamdump/encoding/bcd"
	"dreamdump/encoding/msf"
	"dreamdump/scsi"

	"github.com/pasztorpisti/go-crc"
)

const (
	SUBCHANNEL_Q = 6
)

const (
	Q_CHANNEL_TRACK_NUMBER    = 1
	Q_CHANNEL_INDEX_NUMBER    = 2
	Q_CHANNEL_ABSOLUTE_MINUTE = 7
	Q_CHANNEL_ABSOLUTE_SECOND = 8
	Q_CHANNEL_ABSOLUTE_FRAME  = 9
	Q_CHANNEL_CHECKSUM_MSB    = 10
	Q_CHANNEL_CHECKSUM_LSB    = 11
)

func (sub *Subchannel) Parse(subcodes [scsi.SECTOR_SUB_SIZE]uint8) {
	for i, subcode := range subcodes {
		sub.Qchannel[(i-(i%8))/8] |= ((subcode >> SUBCHANNEL_Q) & 0b00000001) << (7 - (i % 8))
	}
}

func (sub *QChannel) LBA() int32 {
	return sub.AbsolutLBA() - (msf.MSF_SECOND * 2)
}

func (sub *QChannel) AbsolutLBA() int32 {
	minute := (int32(bcd.ToUint8(sub[Q_CHANNEL_ABSOLUTE_MINUTE])) * msf.MSF_MINUTE)
	second := (int32(bcd.ToUint8(sub[Q_CHANNEL_ABSOLUTE_SECOND])) * msf.MSF_SECOND)
	frame := (int32(bcd.ToUint8(sub[Q_CHANNEL_ABSOLUTE_FRAME])))
	return minute + second + frame
}

func (sub *QChannel) TrackNumber() uint8 {
	return bcd.ToUint8(sub[Q_CHANNEL_TRACK_NUMBER])
}

func (sub *QChannel) IndexNumber() uint8 {
	return bcd.ToUint8(sub[Q_CHANNEL_INDEX_NUMBER])
}

func (sub *QChannel) TrackType() uint8 {
	if (sub[0] & 0b01000000) == 0 {
		return TRACK_TYPE_AUDIO
	}

	return TRACK_TYPE_DATA
}

func (sub *QChannel) CheckParity() bool {
	return crc.CRC16GSM.Calc(sub[0:Q_CHANNEL_CHECKSUM_MSB]) == ((uint16(sub[Q_CHANNEL_CHECKSUM_MSB]) << 8) | uint16(sub[Q_CHANNEL_CHECKSUM_LSB]))
}
