package sections_test

import (
	"testing"

	"dreamdump/cd"
	"dreamdump/cd/sections"
	"dreamdump/scsi"

	"gotest.tools/v3/assert"
)

func TestGoodHashes(t *testing.T) {
	section := sections.Section{}
	section.Sectors = []cd.Sector{}

	for range sections.DC_INTERVAL {
		data := cd.CdSectorData{}
		for i := range scsi.SECTOR_DATA_SIZE {
			data[i] = 0
		}
		section.Sectors = append(section.Sectors, cd.Sector{
			Data: data,
		})
	}
	firstHash := section.Hash()
	section.Sectors = []cd.Sector{}

	for range sections.DC_INTERVAL {
		data := cd.CdSectorData{}
		for i := range scsi.SECTOR_DATA_SIZE {
			data[i] = 0
		}
		section.Sectors = append(section.Sectors, cd.Sector{
			Data: data,
		})
	}
	secondHash := section.Hash()

	assert.Equal(t, firstHash, secondHash)
}

func TestSecondTryHashes(t *testing.T) {
	section := sections.Section{}
	section.Sectors = []cd.Sector{}

	for range sections.DC_INTERVAL {
		data := cd.CdSectorData{}
		for i := range scsi.SECTOR_DATA_SIZE {
			data[i] = 0
		}
		section.Sectors = append(section.Sectors, cd.Sector{
			Data: data,
		})
	}
	hash := section.Hash()
	assert.Equal(t, section.IsMatching(hash), false)
	section.AddHash(hash)
	assert.Equal(t, len(section.Hashes), 1)

	section.Sectors = []cd.Sector{}
	for range sections.DC_INTERVAL {
		data := cd.CdSectorData{}
		for i := range scsi.SECTOR_DATA_SIZE {
			data[i] = 1
		}
		section.Sectors = append(section.Sectors, cd.Sector{
			Data: data,
		})
	}
	hash = section.Hash()
	assert.Equal(t, section.IsMatching(hash), false)
	section.AddHash(hash)
	assert.Equal(t, len(section.Hashes), 2)

	section.Sectors = []cd.Sector{}
	for range sections.DC_INTERVAL {
		data := cd.CdSectorData{}
		for i := range scsi.SECTOR_DATA_SIZE {
			data[i] = 0
		}
		section.Sectors = append(section.Sectors, cd.Sector{
			Data: data,
		})
	}
	hash = section.Hash()
	assert.Equal(t, section.IsMatching(hash), true)
	section.AddHash(hash)
	assert.Equal(t, len(section.Hashes), 3)
}

func BenchmarkRandomHashes(b *testing.B) {
	section := sections.Section{}
	for i := range 256 {
		section.Sectors = []cd.Sector{}

		for range sections.DC_INTERVAL {
			data := cd.CdSectorData{}
			for pos := range scsi.SECTOR_DATA_SIZE {
				data[pos] = uint8(i)
			}
			section.Sectors = append(section.Sectors, cd.Sector{
				Data: data,
			})
		}
		hash := section.Hash()
		assert.Equal(b, section.IsMatching(hash), (uint8(0) == uint8(i) && i != 0))
		section.AddHash(hash)
		assert.Equal(b, len(section.Hashes), i+1)
	}
}
