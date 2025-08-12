package cd

import "math"

type QToc struct {
	Tracks map[uint8]Track
}

func (qtoc *QToc) AddSector(sector Sector) {
	if qtoc.Tracks == nil {
		qtoc.Tracks = make(map[uint8]Track)
	}

	if track, ok := qtoc.Tracks[sector.SubcodeTrackNumber()]; ok {
		if index, ok := track.Indexs[sector.SubcodeIndexNumber()]; ok {
			index.LBA = min(sector.SubcodeLBA(), index.LBA)
			track.Indexs[sector.SubcodeIndexNumber()] = index
		} else {
			track.Indexs[sector.SubcodeIndexNumber()] = Index{
				LBA: sector.SubcodeLBA(),
			}
		}

		if sector.SubcodeIndexNumber() == 1 {
			track.LBA = min(sector.SubcodeLBA(), track.LBA)
		}

		qtoc.Tracks[sector.SubcodeTrackNumber()] = track
	} else {
		lba := sector.SubcodeLBA()
		if sector.SubcodeIndexNumber() == 0 {
			lba = math.MaxInt32
		}
		qtoc.Tracks[sector.SubcodeTrackNumber()] = Track{
			LBA:  lba,
			Type: sector.SubcodeTrackType(),
			Indexs: map[uint8]Index{sector.SubcodeIndexNumber(): {
				LBA: sector.SubcodeLBA(),
			}},
		}
	}
}
