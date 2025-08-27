package cd_test

import (
	"testing"

	"dreamdump/cd"

	"gotest.tools/v3/assert"
)

func TestAddSector(t *testing.T) {
	qtoc := cd.QTocNew()

	qchannel := cd.QChannel{
		0x41, 0x05, 0x00, 0x36, 0x51, 0x15, 0x00, 0xA4,
		0x28, 0x32, 0x87, 0xFF,
	}
	qtoc.AddSector(&qchannel)

	qchannel = cd.QChannel{
		0x41, 0x05, 0x01, 0x36, 0x51, 0x15, 0x00, 0xA4,
		0x28, 0x33, 0xD0, 0x0D,
	}
	qtoc.AddSector(&qchannel)

	qchannel = cd.QChannel{
		0x41, 0x05, 0x01, 0x36, 0x51, 0x15, 0x00, 0xA4,
		0x28, 0x34, 0xA0, 0xEA,
	}
	qtoc.AddSector(&qchannel)

	qchannel = cd.QChannel{
		0x41, 0x06, 0x00, 0x36, 0x51, 0x15, 0x00, 0xA4,
		0x28, 0x35, 0xDA, 0x5C,
	}
	qtoc.AddSector(&qchannel)

	qchannel = cd.QChannel{
		0x41, 0x06, 0x01, 0x36, 0x51, 0x15, 0x00, 0xA4,
		0x28, 0x36, 0xAD, 0xEC,
	}
	qtoc.AddSector(&qchannel)

	if track, ok := qtoc.Tracks[5]; ok {
		assert.Equal(t, track.Lba, int32(469983))
		assert.Equal(t, track.GetStartLBA(), int32(469982))
		assert.Equal(t, track.LbaEnd, int32(469985))

		if index, ok := track.Indexs[0]; ok {
			assert.Equal(t, index.LBA, int32(469982))
		} else {
			t.Error()
		}

		if index, ok := track.Indexs[1]; ok {
			assert.Equal(t, index.LBA, int32(469983))
		} else {
			t.Error()
		}
	} else {
		t.Error()
	}

	if track, ok := qtoc.Tracks[6]; ok {
		assert.Equal(t, track.Lba, int32(469986))
		assert.Equal(t, track.GetStartLBA(), int32(469985))

		if index, ok := track.Indexs[0]; ok {
			assert.Equal(t, index.LBA, int32(469985))
		} else {
			t.Error()
		}

		if index, ok := track.Indexs[1]; ok {
			assert.Equal(t, index.LBA, int32(469986))
		} else {
			t.Error()
		}
	} else {
		t.Error()
	}
}
