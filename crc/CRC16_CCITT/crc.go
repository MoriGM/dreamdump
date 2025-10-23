package crc16_ccitt

const (
	POLYNOMIAL uint16 = 0x1021
)

// https://problemkaputt.de/psxspx-cdrom-subchannels.htm
func Calculate(numbers []uint8) uint16 {
	var lsb, msb uint8

	for _, number := range numbers {
		x := number ^ msb
		x ^= (x >> 4)
		msb = lsb ^ (x >> 3) ^ (x << 4)
		lsb = x ^ (x << 5)
	}
	return uint16(msb^0xFF)<<8 | uint16(lsb^0xFF)
}
