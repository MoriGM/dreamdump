package cli

import (
	"fmt"

	"dreamdump/cd"
	"dreamdump/cd/sections"
	"dreamdump/option"
)

func DreamDumpDisc(opt *option.Option) {
	sectionMap := sections.GetSectionMap(opt)
	sections.ReadSections(opt, &sectionMap)
	sectors := sections.ExtractSectionsToSectors(&sectionMap)
	qtoc := cd.QTocNew()
	qtoc.AddSectors(&sectors)
	for i, track := range qtoc.Tracks {
		fmt.Println(i, *track)
		for ii, index := range track.Indexs {
			fmt.Println(ii, *index)
		}
	}
}
