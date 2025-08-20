package sections

import (
	"dreamdump/cd"
	"dreamdump/option"
)

const (
	DC_START          = 44990
	DC_END            = 549151
	DC_INTERVAL       = 10289
	DC_DEFAULT_CUTOFF = 446261
)

type Section struct {
	StartSector uint32
	EndSector   uint32
	Hashes      []string
	Sectors     []cd.Sector
}

func GetSectionMap(opt *option.Option) []Section {
	count := (opt.CutOff - DC_START) / DC_INTERVAL
	var sections []Section
	for i := 0; i < int(count); i++ {
		sections = append(sections, Section{
			StartSector: uint32(DC_START + (i * DC_INTERVAL)),
			EndSector:   uint32(DC_START + ((i + 1) * DC_INTERVAL)),
			Hashes:      []string{},
			Sectors:     []cd.Sector{},
		})
	}
	sections = append(sections, Section{
		StartSector: uint32(DC_START + (count * DC_INTERVAL)),
		EndSector:   uint32(DC_END),
		Hashes:      []string{},
		Sectors:     []cd.Sector{},
	})

	return sections
}
