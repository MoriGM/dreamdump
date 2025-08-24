package crc16_ccitt_test

import (
	"testing"

	crc16_ccitt "dreamdump/crc/CRC16_CCITT"

	"gotest.tools/v3/assert"
)

func TestCRCGood(t *testing.T) {
	assert.Equal(t, crc16_ccitt.Calculate([]uint8{}), uint16(0xFFFF))
	assert.Equal(t, crc16_ccitt.Calculate([]uint8{'A'}), uint16(0xA71A))
}
