package sections

import (
	"dreamdump/cd"
	"dreamdump/option"
)

type Section struct {
	StartSector int32
	EndSector   int32
	Hashes      []string
	Sectors     []cd.Sector
	Matched     bool
}

func GetSectionMap(opt *option.Option) []Section {
	count := option.DC_START
	var sections []Section
	for {
		sections = append(sections, Section{
			StartSector: count,
			EndSector:   min(count+option.DC_INTERVAL, opt.CutOff),
			Hashes:      []string{},
			Matched:     false,
		})
		count += option.DC_INTERVAL
		if count >= opt.CutOff {
			break
		}
	}
	sections = append(sections, Section{
		StartSector: opt.CutOff,
		EndSector:   option.DC_END,
		Hashes:      []string{},
		Matched:     false,
	})

	return sections
}
