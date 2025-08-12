package bcd

func ToUint8(value byte) uint8 {
	upper := value & 0xF0
	lower := value & 0x0F
	return ((upper / 16) * 10) + lower
}
