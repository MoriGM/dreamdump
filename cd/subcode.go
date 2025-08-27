package cd

import (
	crc16_ccitt "dreamdump/crc/CRC16_CCITT"
	"dreamdump/encoding/bcd"
	"dreamdump/scsi"
)

const (
	SUBCHANNEL_Q = 6
)

const (
	Q_CHANNEL_TRACK_NUMBER = 1
	Q_CHANNEL_INDEX_NUMBER = 2
	Q_CHANNEL_MINUTE       = 7
	Q_CHANNEL_SECOND       = 8
	Q_CHANNEL_FRAME        = 9
	Q_CHANNEL_CHECKSUM_MSB = 10
	Q_CHANNEL_CHECKSUM_LSB = 11
)

func (sub *Subchannel) Parse(subcodes [scsi.SECTOR_SUB_SIZE]uint8) {
	for i, subcode := range subcodes {
		sub.Qchannel[(i-(i%8))/8] |= ((subcode >> SUBCHANNEL_Q) & 0b00000001) << (7 - (i % 8))
	}
}

func (sub *QChannel) LBA() int32 {
	return sub.AbsolutLBA() - (MSF_SECOND * 2)
}

func (sub *QChannel) AbsolutLBA() int32 {
	return (int32(bcd.ToUint8(sub[Q_CHANNEL_MINUTE])) * MSF_MINUTE) + (int32(bcd.ToUint8(sub[Q_CHANNEL_SECOND])) * MSF_SECOND) + (int32(bcd.ToUint8(sub[Q_CHANNEL_FRAME])))
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
	return crc16_ccitt.Calculate(sub[0:10]) == ((uint16(sub[Q_CHANNEL_CHECKSUM_MSB]) << 8) | uint16(sub[Q_CHANNEL_CHECKSUM_LSB]))
}
