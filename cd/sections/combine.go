package sections

import (
	"os"

	"dreamdump/option"
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
	defer scramFile.Close()
	for sectionNumber := range len(sections) {
		section := sections[sectionNumber]
		for _, sector := range section.Sectors {
			_, err := scramFile.Write(sector.Data[:])
			if err != nil {
				panic(err)
			}
		}
	}
}

func CombineToSub(opt *option.Option, sections []*Section) {
	scramFileName := opt.PathName + "/" + opt.ImageName + ".subq"
	scramFile, err := os.OpenFile(scramFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer scramFile.Close()
	for sectionNumber := range len(sections) {
		section := sections[sectionNumber]
		for _, sector := range section.Sectors {
			_, err := scramFile.Write(sector.Sub.Qchannel[:])
			if err != nil {
				panic(err)
			}
		}
	}
}
