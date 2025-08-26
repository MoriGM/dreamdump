package sections

import (
	"dreamdump/cd"
	"dreamdump/option"
)

const (
	DC_START          int32 = 44990
	DC_END            int32 = 549151
	DC_INTERVAL       int32 = 10289
	DC_DEFAULT_CUTOFF int32 = 446261
)

type Section struct {
	StartSector int32
	EndSector   int32
	Hashes      []string
	Sectors     []cd.Sector
	Matched     bool
}

func GetSectionMap(opt *option.Option) []Section {
	count := DC_START
	var sections []Section
	for {
		sections = append(sections, Section{
			StartSector: count,
			EndSector:   min(count+DC_INTERVAL, opt.CutOff),
			Hashes:      []string{},
			Matched:     false,
		})
		count += DC_INTERVAL
		if count >= opt.CutOff {
			break
		}
	}
	sections = append(sections, Section{
		StartSector: opt.CutOff,
		EndSector:   DC_END,
		Hashes:      []string{},
		Matched:     false,
	})

	return sections
}
