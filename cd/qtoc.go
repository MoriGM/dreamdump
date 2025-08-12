package cd

type QToc struct {
	Tracks map[uint8]Track
}

func (qtoc *QToc) AddSector(sector Sector) {
	if qtoc.Tracks == nil {
		qtoc.Tracks = make(map[uint8]Track)
	}

	if track, ok := qtoc.Tracks[sector.SubcodeTrackNumber()]; ok {
		track.LBA = min(sector.SubcodeLBA(), track.LBA)
	} else {
		qtoc.Tracks[sector.SubcodeTrackNumber()] = Track{
			LBA:  sector.SubcodeLBA(),
			Type: sector.SubcodeTrackType(),
			Indexs: map[uint8]Index{sector.SubcodeIndexNumber(): {
				LBA: sector.SubcodeLBA(),
			}},
		}
	}
}
