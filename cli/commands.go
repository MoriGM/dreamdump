package cli

import (
	"dreamdump/cd"
	"dreamdump/cd/sections"
	"dreamdump/option"
)

func DreamDumpDisc(opt *option.Option) {
	sectionMap := sections.GetSectionMap(opt)
	for {
		allMatching := true
		for _, section := range sectionMap {
			section.Sectors = []cd.Sector{}
			err := sections.ReadSection(&section)
			if err != nil {
				continue
			}
			hash := section.Hash()
			if section.IsMatching(hash) {
				continue
			}
			allMatching = false
			section.AddHash(hash)
		}

		if allMatching {
			break
		}
	}
}
