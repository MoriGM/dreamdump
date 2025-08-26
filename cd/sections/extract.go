package sections

import (
	"dreamdump/cd"
	"dreamdump/option"
	"dreamdump/scsi"
)

func ExtractSectionsToDense(opt *option.Option, sections *[]Section) *cd.Dense {
	dense := make(cd.Dense, (int(DC_END-DC_START)+1)*scsi.SECTOR_DATA_SIZE)

	skip := (opt.ReadOffset * 4)
	if skip < 0 {
		panic("Drive read offset cannot be minus")
	}
	pos := 0
	endPos := scsi.SECTOR_DATA_SIZE - int(skip)
	for sectionNumber := range len(*sections) {
		for sectorNumber := range len((*sections)[sectionNumber].Sectors) {
			copy(dense[pos:endPos], (*sections)[sectionNumber].Sectors[sectorNumber].Data[skip:scsi.SECTOR_DATA_SIZE])
			pos += scsi.SECTOR_DATA_SIZE - int(skip)
			endPos += scsi.SECTOR_DATA_SIZE
			endPos = min(endPos, int(DC_END-DC_START)*scsi.SECTOR_DATA_SIZE)
			skip = 0
		}
	}

	return &dense
}

func ExtractSectionsToQtoc(sections *[]Section) *cd.QToc {
	qtoc := cd.QTocNew()
	for sectionNumber := range len(*sections) {
		for sectorNumber := range len((*sections)[sectionNumber].Sectors) {
			qtoc.AddSector(&(*sections)[sectionNumber].Sectors[sectorNumber].Sub.Qchannel)
		}
	}
	return qtoc
}
