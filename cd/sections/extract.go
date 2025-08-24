package sections

import "dreamdump/cd"

func ExtractSectionsToSectors(sections *[]Section) []cd.Sector {
	sectors := make([]cd.Sector, int(DC_END-DC_START)+1)

	counter := 0
	for sectionNumber := range len(*sections) {
		for sectorNumber := range len((*sections)[sectionNumber].Sectors) {
			sectors[counter] = (*sections)[sectionNumber].Sectors[sectorNumber]
			counter++
		}
	}

	return sectors
}
