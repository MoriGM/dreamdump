package sections

import (
	"dreamdump/cd"
	"dreamdump/option"
	"dreamdump/scsi"
)

func ExtractSections(opt *option.Option, sections []*Section) (*cd.Dense, *cd.QToc) {
	qtoc := ExtractSectionsToQtoc(sections)
	dense := ExtractSectionsToDense(opt, sections)
	return dense, qtoc
}

func ExtractSectionsToDense(opt *option.Option, sections []*Section) *cd.Dense {
	dense := make(cd.Dense, (int(option.DC_END-option.DC_START)+1)*scsi.SECTOR_DATA_SIZE)

	skip := (opt.ReadOffset * 4)
	if skip < 0 {
		panic("Drive read offset cannot be minus")
	}
	pos := 0
	endPos := scsi.SECTOR_DATA_SIZE - int(skip)
	for sectionNumber := range len(sections) {
		for sectorNumber := range len(sections[sectionNumber].Sectors) {
			copy(dense[pos:endPos], sections[sectionNumber].Sectors[sectorNumber].Data[skip:scsi.SECTOR_DATA_SIZE])
			sections[sectionNumber].Sectors[sectorNumber] = nil
			pos += scsi.SECTOR_DATA_SIZE - int(skip)
			endPos += scsi.SECTOR_DATA_SIZE
			endPos = min(endPos, int(option.DC_END-option.DC_START)*scsi.SECTOR_DATA_SIZE)
			skip = 0
		}
	}

	return &dense
}

func ExtractSectionsToQtoc(sections []*Section) *cd.QToc {
	qtoc := cd.QTocNew()
	for sectionNumber := range len(sections) {
		for sectorNumber := range len(sections[sectionNumber].Sectors) {
			qtoc.AddSector(&sections[sectionNumber].Sectors[sectorNumber].Sub.Qchannel)
		}
	}
	qtoc.Sort()
	return qtoc
}
