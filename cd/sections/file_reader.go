package sections

import (
	"errors"
	"os"

	"dreamdump/cd"
	"dreamdump/option"
	"dreamdump/scsi"
)

func (sect *Section) ReadSection(opt *option.Option) {
	binFileName := sect.FileName(opt) + ".bin"
	_, err := os.Stat(binFileName)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	binFile, err := os.OpenFile(binFileName, os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer binFile.Close()

	subFileName := sect.FileName(opt) + ".subq"
	_, err = os.Stat(subFileName)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	subFile, err := os.OpenFile(subFileName, os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer subFile.Close()

	sect.Sectors = make([]cd.Sector, sect.EndSector-sect.StartSector)
	for i := range sect.EndSector - sect.StartSector {
		data := make([]byte, scsi.SECTOR_DATA_SIZE)
		subq := make([]byte, scsi.SECTOR_SUBQ_SIZE)
		_, err = binFile.ReadAt(data, int64(i)*int64(scsi.SECTOR_DATA_SIZE))
		if err != nil {
			panic(err)
		}
		_, err = subFile.ReadAt(subq, int64(i)*int64(scsi.SECTOR_SUBQ_SIZE))
		if err != nil {
			panic(err)
		}
		copy(sect.Sectors[i].Data[:], data[:])
		copy(sect.Sectors[i].Sub.Qchannel[:], subq[:])
	}
	sect.Matched = true
}
