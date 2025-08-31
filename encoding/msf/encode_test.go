package msf_test

import (
	"testing"

	"dreamdump/encoding/msf"

	"gotest.tools/v3/assert"
)

func TestEncodeZero(t *testing.T) {
	assert.Equal(t, "00:00:00", msf.Encode(0))
}

func TestEncodeFrame(t *testing.T) {
	assert.Equal(t, "00:00:01", msf.Encode(msf.MSF_FRAME))
}

func TestEncodeSecond(t *testing.T) {
	assert.Equal(t, "00:01:00", msf.Encode(msf.MSF_SECOND))
}

func TestEncodeMinute(t *testing.T) {
	assert.Equal(t, "01:00:00", msf.Encode(msf.MSF_MINUTE))
}

func TestEncodeTwoSeconds(t *testing.T) {
	assert.Equal(t, "00:02:00", msf.Encode(msf.MSF_SECOND*2))
}
