package bigendian_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	bigendian "dreamdump/encoding/big_endian"

	"gotest.tools/v3/assert"
)

func TestInt32(t *testing.T) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, bigendian.Int32(0x77AA5511))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, buf.Len(), 4)
	assert.Equal(t, buf.Bytes()[3], byte(0x77))
	assert.Equal(t, buf.Bytes()[2], byte(0xAA))
	assert.Equal(t, buf.Bytes()[1], byte(0x55))
	assert.Equal(t, buf.Bytes()[0], byte(0x11))
}

func TestUint16(t *testing.T) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, bigendian.Uint16(0x77AA))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, buf.Len(), 2)
	assert.Equal(t, buf.Bytes()[1], byte(0x77))
	assert.Equal(t, buf.Bytes()[0], byte(0xAA))
}
