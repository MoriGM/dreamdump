package sections

import (
	"dreamdump/cd"
	"dreamdump/option"
)

type Section struct {
	StartSector int32
	EndSector   int32
	Hashes      []string
	Sectors     []*cd.Sector
	Matched     bool
}

func GetSectionMap(opt *option.Option) []*Section {
	count := option.DC_START
	sections := make([]*Section, 0)
	for {
		section := new(Section)
		section.StartSector = count
		section.EndSector = min(count+option.DC_INTERVAL, opt.CutOff)
		section.Hashes = []string{}
		section.Matched = false
		sections = append(sections, section)
		count += option.DC_INTERVAL
		if count >= opt.CutOff {
			break
		}
	}

	section := new(Section)
	section.StartSector = count
	section.EndSector = option.DC_END
	section.Hashes = []string{}
	section.Matched = false
	sections = append(sections, section)

	return sections
}
