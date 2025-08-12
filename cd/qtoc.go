package cd

type QToc struct {
	tracks []Track
}

func (qtoc *QToc) AddSector(sector Sector) {
	if len(qtoc.tracks) == 0 && sector.SubcodeTrackNumber() > 0 {
	}
}
