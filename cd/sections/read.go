package sections

import (
	"strconv"

	"dreamdump/cd"
	"dreamdump/log"
	"dreamdump/option"
)

func ReadSection(opt *option.Option, section *Section) error {
	for i := section.StartSector; i <= section.EndSector; i++ {
		sector, err := cd.ReadSector(opt, i)
		if err != nil {
			return err
		}
		section.Sectors = append(section.Sectors, sector)
		log.WriteLine("Sector read " + strconv.FormatInt(int64(i), 10))
	}
	return nil
}
