package writers

import (
	"os"
	"strconv"

	"dreamdump/cd/sections"
	"dreamdump/option"
)

func WriteSubSection(opt *option.Option, section *sections.Section) {
	sectionRange := strconv.FormatInt(int64(section.StartSector), 10) + "-" + strconv.FormatInt(int64(section.EndSector), 10)
	binFile, err := os.OpenFile(opt.ImageName+"-"+sectionRange+".subq", os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer binFile.Close()
	for _, sector := range section.Sectors {
		_, err = binFile.Write(sector.Sub.Qchannel[:])
		if err != nil {
			panic(err)
		}
	}
	err = binFile.Sync()
	if err != nil {
		panic(err)
	}
}
