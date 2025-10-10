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
		sectionMap[sectionNumber].ReadHash(opt)
		sectionMap[sectionNumber].ReadSection(opt)
	}
}

func ReadSections(opt *option.Option, sectionMap []*Section) {
	ReadFileSections(opt, sectionMap)
	for {
		if opt.Train && !sectionMap[0].Matched {
			TrainStart(opt)
		}
		allMatching := true
		var sectionNumber int
		for sectionPosition := range len(sectionMap) {
			section := sectionMap[sectionPosition]
			sectionNumber = sectionPosition + 1
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
				section.WriteSection(opt)
				section.Matched = true
				log.Println("Section hash is matching " + strconv.Itoa(sectionNumber) + " Hash:" + hash)
				continue
			}

			section.Sectors = nil

			allMatching = false
			section.AddHash(hash)
			section.WriteHash(opt)

			if len(section.Hashes) == 1 {
				log.Println("Inital section " + strconv.Itoa(sectionNumber) + " read from " + strconv.FormatInt(int64(section.StartSector), 10) + " to " + strconv.FormatInt(int64(section.EndSector), 10))
				continue
			}
			log.Println("Section " + strconv.Itoa(sectionNumber) + " not matching and read from " + strconv.FormatInt(int64(section.StartSector), 10) + " to " + strconv.FormatInt(int64(section.EndSector), 10))
		}

		if allMatching {
			return
		}
	}
}

func ReadSection(opt *option.Option, section *Section) error {
	for lba := section.StartSector; lba < section.EndSector; lba += int32(opt.ReadAtOnce) {
		readAmount := opt.ReadAtOnce
		if (lba + int32(opt.ReadAtOnce)) > section.EndSector {
			readAmount = opt.ReadAtOnce - uint8(lba+int32(opt.ReadAtOnce)-(section.EndSector))
		}

		sectors, err := cd.ReadSectors(opt, lba, readAmount)
		if err != nil {
			return errors.Join(errors.New("scsi error while reading sector "+strconv.FormatInt(int64(lba), 10)), err)
		}

		for i, sector := range sectors {
			if sector.C2.Amount() > 0 {
				return errors.New("error reading sector " + strconv.FormatInt(int64(lba), 10) + " as it contained a c2 error")
			}
			section.Sectors[(lba+int32(i))-section.StartSector] = sector
		}

		log.PrintClean("Sector read " + strconv.FormatInt(int64(lba), 10))
	}
	return nil
}
