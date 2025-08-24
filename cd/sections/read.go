package sections

import (
	"fmt"
	"strconv"

	"dreamdump/cd"
	"dreamdump/log"
	"dreamdump/option"
)

func ReadSections(opt *option.Option, sectionMap *[]Section) {
	for {
		allMatching := true
		for sectionNumber := range len(*sectionMap) {
			section := &(*sectionMap)[sectionNumber]
			if section.Matched {
				continue
			}
			fmt.Println(section.EndSector - section.StartSector)
			section.Sectors = make([]cd.Sector, section.EndSector-section.StartSector)
			err := ReadSection(opt, section)
			if err != nil {
				allMatching = false
				section.Sectors = nil
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
			section.Sectors = nil
			if len(section.Hashes) > 1 {
				log.WriteLn("Section " + strconv.Itoa(sectionNumber) + " not matching and read from " + strconv.FormatInt(int64(section.StartSector), 10) + " to " + strconv.FormatInt(int64(section.EndSector), 10))
			} else {
				log.WriteLn("Inital section " + strconv.Itoa(sectionNumber) + " read from " + strconv.FormatInt(int64(section.StartSector), 10) + " to " + strconv.FormatInt(int64(section.EndSector), 10))
			}
		}

		if allMatching {
			return
		}
	}
}

func ReadSection(opt *option.Option, section *Section) error {
	for i := section.StartSector; i < section.EndSector; i++ {
		sector, err := cd.ReadSector(opt, i)
		if err != nil {
			return err
		}
		if sector.C2.Amount() > 0 {
			return fmt.Errorf("error reading sector as it contained a c2 error")
		}
		section.Sectors[i-section.StartSector] = sector
		log.WriteCleanLine("Sector read " + strconv.FormatInt(int64(i), 10))
	}
	return nil
}
