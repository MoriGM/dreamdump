package sections

import (
	"os"

	"dreamdump/cd"
	"dreamdump/option"
	"dreamdump/scsi"
)

func ExtractSections(opt *option.Option, sections []*Section) (*cd.Dense, *cd.QToc) {
	qtoc := ExtractSectionsToQtoc(opt, sections)
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

func ExtractSectionsToQtoc(opt *option.Option, sections []*Section) *cd.QToc {
	qtoc := cd.QTocNew()
	subcodeFileName := opt.PathName + "/" + opt.ImageName + ".subq"
	subcodeFile, err := os.OpenFile(subcodeFileName, os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}

	qchannel := make([]byte, 12)
	for {
		size, err := subcodeFile.Read(qchannel)
		if size < scsi.CHANNEL_SIZE {
			break
		}
		if err != nil {
			panic(err)
		}
		qtoc.AddSector((*cd.QChannel)(qchannel))
	}
	qtoc.Sort()
	return qtoc
}
