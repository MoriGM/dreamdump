package sections

import (
	"dreamdump/cd"
	"dreamdump/option"
)

func ExtractSectionsToSectors(opt *option.Option, sections *[]Section) *cd.Dense {
	dense := make(cd.Dense, int(DC_END-DC_START)+1)

	skip := (opt.ReadOffset * 4)
	if skip < 0 {
		panic("Drive read offset cannot be minus")
	}
	for sectionNumber := range len(*sections) {
		for sectorNumber := range len((*sections)[sectionNumber].Sectors) {
			copy(dense, (*sections)[sectionNumber].Sectors[sectorNumber].Data[skip:])
			skip = 0
		}
	}

	return &dense
}

func ExtractSectionsToQtoc(sections *[]Section) *cd.QToc {
	qtoc := new(cd.QToc)
	for sectionNumber := range len(*sections) {
		for sectorNumber := range len((*sections)[sectionNumber].Sectors) {
			qtoc.AddSector(&(*sections)[sectionNumber].Sectors[sectorNumber].Sub.Qchannel)
		}
	}
	return qtoc
}
