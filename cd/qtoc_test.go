package cd_test

import (
	"testing"

	"dreamdump/cd"

	"gotest.tools/v3/assert"
)

func TestAddSector(t *testing.T) {
	qtoc := cd.QToc{}

	sector := cd.Sector{}
	sector.Sub = [16]uint8{
		0x41, 0x05, 0x00, 0x36, 0x51, 0x15, 0x00, 0xA4,
		0x28, 0x32, 0xD0, 0x0D, 0xBE, 0x7B, 0xD9, 0x7D,
	}
	qtoc.AddSector(sector)

	sector = cd.Sector{}
	sector.Sub = [16]uint8{
		0x41, 0x05, 0x01, 0x36, 0x51, 0x15, 0x00, 0xA4,
		0x28, 0x33, 0xD0, 0x0D, 0xBE, 0x7B, 0xD9, 0x7D,
	}
	qtoc.AddSector(sector)

	sector = cd.Sector{}
	sector.Sub = [16]uint8{
		0x41, 0x05, 0x01, 0x36, 0x51, 0x15, 0x00, 0xA4,
		0x28, 0x34, 0xD0, 0x0D, 0xBE, 0x7B, 0xD9, 0x7D,
	}
	qtoc.AddSector(sector)

	sector = cd.Sector{}
	sector.Sub = [16]uint8{
		0x41, 0x06, 0x00, 0x36, 0x51, 0x15, 0x00, 0xA4,
		0x28, 0x35, 0xD0, 0x0D, 0xBE, 0x7B, 0xD9, 0x7D,
	}
	qtoc.AddSector(sector)

	sector = cd.Sector{}
	sector.Sub = [16]uint8{
		0x41, 0x06, 0x01, 0x36, 0x51, 0x15, 0x00, 0xA4,
		0x28, 0x36, 0xD0, 0x0D, 0xBE, 0x7B, 0xD9, 0x7D,
	}
	qtoc.AddSector(sector)

	if track, ok := qtoc.Tracks[5]; ok {
		assert.Equal(t, track.LBA, int32(469983))

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
		assert.Equal(t, track.LBA, int32(469986))

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
