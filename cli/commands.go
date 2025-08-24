package cli

import (
	"strconv"

	"dreamdump/cd"
	"dreamdump/cd/sections"
	"dreamdump/log"
	"dreamdump/option"
)

func DreamDumpDisc(opt *option.Option) {
	sectionMap := sections.GetSectionMap(opt)
	for {
		allMatching := true
		for sectionNumber, section := range sectionMap {
			if section.Matched {
				continue
			}
			section.Sectors = []cd.Sector{}
			err := sections.ReadSection(opt, &section)
			if err != nil {
				log.WriteLn("Error while reading section " + strconv.Itoa(sectionNumber))
				continue
			}
			hash := section.Hash()
			if section.IsMatching(hash) {
				section.Matched = true
				log.WriteLn("Section hash is matching " + strconv.Itoa(sectionNumber) + " Hash:" + hash)
				continue
			}
			allMatching = false
			section.AddHash(hash)
			if len(section.Hashes) > 1 {
				log.WriteLine("Section not matching read " + strconv.Itoa(sectionNumber))
			} else {
				log.WriteLine("Inital section read " + strconv.Itoa(sectionNumber))
			}
		}

		if allMatching {
			break
		}
	}
}
