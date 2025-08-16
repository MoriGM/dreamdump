package bcd_test

import (
	"dreamdump/encoding/bcd"
	"testing"

	"gotest.tools/v3/assert"
)

func TestZero(t *testing.T) {
	assert.Equal(t, bcd.ToUint8(0x00), uint8(0x00))
}

func Test99(t *testing.T) {
	assert.Equal(t, bcd.ToUint8(0x99), uint8(99))
}

func TestRandom(t *testing.T) {
	assert.Equal(t, bcd.ToUint8(0x01), uint8(1))
	assert.Equal(t, bcd.ToUint8(0x12), uint8(12))
	assert.Equal(t, bcd.ToUint8(0x46), uint8(46))
	assert.Equal(t, bcd.ToUint8(0x55), uint8(55))
	assert.Equal(t, bcd.ToUint8(0x69), uint8(69))
	assert.Equal(t, bcd.ToUint8(0x77), uint8(77))
	assert.Equal(t, bcd.ToUint8(0x89), uint8(89))
}

func TestA0(t *testing.T) {
	assert.Equal(t, bcd.ToUint8(0xA0), uint8(100))
}
