package cli

import (
	"errors"
	"fmt"
	"os"

	"dreamdump/cd/sections"
	"dreamdump/log"
	"dreamdump/option"
)

func DreamDumpDisc(opt *option.Option) {
	_, err := os.Stat(opt.PathName)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(opt.PathName, 0o744)
		if err != nil {
			panic(err)
		}
	}
	sectionMap := sections.GetSectionMap(opt)
	sections.ReadSections(opt, &sectionMap)
	qtoc := sections.ExtractSectionsToQtoc(&sectionMap)
	dense := sections.ExtractSectionsToDense(opt, &sectionMap)
	fmt.Printf("Write Offset: %d\n", dense.NewOffsetManager(sections.DC_START).SampleOffset)
	qtoc.Print()
	trackMetas := dense.Split(opt, qtoc)
	for _, trackMeta := range trackMetas {
		romVaultLine := fmt.Sprintf("<rom name=\"%s\" size=\"%d\" crc=\"%x\" md5=\"%x\" sha1=\"%x\" />", trackMeta.FileName, trackMeta.Size, trackMeta.CRC32, trackMeta.MD5, trackMeta.SHA1)
		log.Println(romVaultLine)
	}
}
