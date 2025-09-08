package sections

import (
	"errors"
	"strconv"

	"dreamdump/cd"
	"dreamdump/log"
	"dreamdump/option"
)

func ReadFileSections(opt *option.Option, sectionMap []*Section) {
	for sectionNumber := range len(sectionMap) {
		sectionMap[sectionNumber].ReadSection(opt)
	}
}

func ReadSections(opt *option.Option, sectionMap []*Section) {
	ReadFileSections(opt, sectionMap)
	for {
		allMatching := true
		for sectionNumber := range len(sectionMap) {
			section := sectionMap[sectionNumber]
			if section.Matched {
				continue
			}
			section.Sectors = make([]*cd.Sector, section.EndSector-section.StartSector)
			err := ReadSection(opt, section)
			if err != nil {
				allMatching = false
				section.Sectors = nil
				log.Println("Error while reading section " + strconv.Itoa(sectionNumber) + "\nError text: " + err.Error())
				continue
			}
			hash := section.Hash()
			if section.IsMatching(hash) {
				section.Matched = true
				log.Println("Section hash is matching " + strconv.Itoa(sectionNumber) + " Hash:" + hash)
				section.WriteSection(opt)
				continue
			}
			allMatching = false
			section.AddHash(hash)
			section.Sectors = nil
			if len(section.Hashes) > 1 {
				log.Println("Section " + strconv.Itoa(sectionNumber) + " not matching and read from " + strconv.FormatInt(int64(section.StartSector), 10) + " to " + strconv.FormatInt(int64(section.EndSector), 10))
			} else {
				log.Println("Inital section " + strconv.Itoa(sectionNumber) + " read from " + strconv.FormatInt(int64(section.StartSector), 10) + " to " + strconv.FormatInt(int64(section.EndSector), 10))
			}
		}

		if allMatching {
			return
		}
	}
}

func ReadSection(opt *option.Option, section *Section) error {
	for lba := section.StartSector; lba < section.EndSector; lba++ {
		sector, err := cd.ReadSector(opt, lba)
		if err != nil {
			return errors.Join(errors.New("scsi error while reading sector "+strconv.FormatInt(int64(lba), 10)), err)
		}
		if sector.C2.Amount() > 0 {
			return errors.New("error reading sector " + strconv.FormatInt(int64(lba), 10) + " as it contained a c2 error")
		}
		section.Sectors[lba-section.StartSector] = sector
		log.PrintClean("Sector read " + strconv.FormatInt(int64(lba), 10))
	}
	return nil
}
