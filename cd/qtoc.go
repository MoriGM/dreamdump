package cd

import (
	"math"
)

type QToc struct {
	Tracks map[uint8]*Track
}

func QTocNew() QToc {
	qtoc := QToc{}
	qtoc.Tracks = make(map[uint8]*Track)
	return qtoc
}

func (qtoc *QToc) AddSectors(sectors *[]Sector) {
	for sectorNumber := range len(*sectors) {
		sector := &(*sectors)[sectorNumber]
		if !sector.Sub.Qchannel.CheckParity() {
			continue
		}
		qtoc.AddSector(sector)
	}
}

func (qtoc *QToc) AddSector(sector *Sector) {
	if track, ok := qtoc.Tracks[sector.Sub.Qchannel.TrackNumber()]; ok {
		if index, ok := track.Indexs[sector.Sub.Qchannel.IndexNumber()]; ok {
			index.LBA = min(sector.Sub.Qchannel.LBA(), index.LBA)
		} else {
			track.Indexs[sector.Sub.Qchannel.IndexNumber()] = &Index{
				LBA: sector.Sub.Qchannel.LBA(),
			}
		}

		if sector.Sub.Qchannel.IndexNumber() == 1 {
			track.LBA = min(sector.Sub.Qchannel.LBA(), track.LBA)
			track.Type = sector.Sub.Qchannel.TrackType()
		}

		return
	}
	lba := sector.Sub.Qchannel.LBA()
	if sector.Sub.Qchannel.IndexNumber() == 0 {
		lba = math.MaxInt32
	}
	qtoc.Tracks[sector.Sub.Qchannel.TrackNumber()] = &Track{
		LBA:  lba,
		Type: sector.Sub.Qchannel.TrackType(),
		Indexs: map[uint8]*Index{sector.Sub.Qchannel.IndexNumber(): {
			LBA: sector.Sub.Qchannel.LBA(),
		}},
	}
}
