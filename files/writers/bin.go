package writers

import (
	"os"

	"dreamdump/cd/sections"
	"dreamdump/option"
)

func WriteBinSection(opt *option.Option, section *sections.Section) {
	binFile, err := os.OpenFile(opt.ImageName+".bin", os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer binFile.Close()
	for _, sector := range section.Sectors {
		_, err = binFile.Write(sector.Data[:])
		if err != nil {
			panic(err)
		}
	}
	err = binFile.Sync()
	if err != nil {
		panic(err)
	}
}
