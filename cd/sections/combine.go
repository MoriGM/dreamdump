package sections

import (
	"os"

	"dreamdump/option"
	"dreamdump/scsi"
)

func CombineSections(opt *option.Option, sections []*Section) {
	CombineToScram(opt, sections)
	CombineToSub(opt, sections)
}

func CombineToScram(opt *option.Option, sections []*Section) {
	scramFileName := opt.PathName + "/" + opt.ImageName + ".scram"
	scramFile, err := os.OpenFile(scramFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	skip := opt.ReadOffset * scsi.SAMPLE_SIZE
	if skip < 0 {
		skip = scsi.SECTOR_DATA_SIZE - skip
		_, err := scramFile.Write(make([]byte, scsi.SECTOR_DATA_SIZE))
		if err != nil {
			panic(err)
		}
	}
	defer scramFile.Close()
	for sectionNumber := range len(sections) {
		section := sections[sectionNumber]
		for _, sector := range section.Sectors {
			_, err := scramFile.Write(sector.Data[skip:])
			if err != nil {
				panic(err)
			}
		}
		skip = 0
	}
}

func CombineToSub(opt *option.Option, sections []*Section) {
	subcodeFileName := opt.PathName + "/" + opt.ImageName + ".subq"
	subcodeFile, err := os.OpenFile(subcodeFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer subcodeFile.Close()
	for sectionNumber := range len(sections) {
		section := sections[sectionNumber]
		for _, sector := range section.Sectors {
			_, err := subcodeFile.Write(sector.Sub.Qchannel[:])
			if err != nil {
				panic(err)
			}
		}
	}
}
