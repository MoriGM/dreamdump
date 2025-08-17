package cd

import "math"

type QToc struct {
	Tracks map[uint8]Track
}

func (qtoc *QToc) AddSector(sector Sector) {
	if qtoc.Tracks == nil {
		qtoc.Tracks = make(map[uint8]Track)
	}

	if track, ok := qtoc.Tracks[sector.Sub.TrackNumber()]; ok {
		if index, ok := track.Indexs[sector.Sub.IndexNumber()]; ok {
			index.LBA = min(sector.Sub.LBA(), index.LBA)
			track.Indexs[sector.Sub.IndexNumber()] = index
		} else {
			track.Indexs[sector.Sub.IndexNumber()] = Index{
				LBA: sector.Sub.LBA(),
			}
		}

		if sector.Sub.IndexNumber() == 1 {
			track.LBA = min(sector.Sub.LBA(), track.LBA)
		}

		qtoc.Tracks[sector.Sub.TrackNumber()] = track
	} else {
		lba := sector.Sub.LBA()
		if sector.Sub.IndexNumber() == 0 {
			lba = math.MaxInt32
		}
		qtoc.Tracks[sector.Sub.TrackNumber()] = Track{
			LBA:  lba,
			Type: sector.Sub.TrackType(),
			Indexs: map[uint8]Index{sector.Sub.IndexNumber(): {
				LBA: sector.Sub.LBA(),
			}},
		}
	}
}
