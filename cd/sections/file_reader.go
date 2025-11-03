package sections

import (
	"bufio"
	"errors"
	"os"

	"dreamdump/cd"
	"dreamdump/option"
	"dreamdump/scsi"
)

func (sect *Section) ReadHash(opt *option.Option) {
	hashFileName := sect.FileName(opt) + ".hash"
	_, err := os.Stat(hashFileName)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	hashFile, err := os.OpenFile(hashFileName, os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer hashFile.Close()
	hashFileScanner := bufio.NewScanner(hashFile)
	for hashFileScanner.Scan() {
		sect.Hashes = append(sect.Hashes, hashFileScanner.Text())
	}
}

func (sect *Section) ReadSection(opt *option.Option) {
	scramFileName := sect.FileName(opt) + ".scram"
	_, err := os.Stat(scramFileName)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	scramFile, err := os.OpenFile(scramFileName, os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer scramFile.Close()

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

	sect.Sectors = make([]*cd.Sector, sect.EndSector-sect.StartSector)
	for i := range sect.EndSector - sect.StartSector {
		data := make([]byte, scsi.SECTOR_DATA_SIZE)
		subq := make([]byte, scsi.CHANNEL_SIZE)
		_, err = scramFile.ReadAt(data, int64(i)*int64(scsi.SECTOR_DATA_SIZE))
		if err != nil {
			panic(err)
		}
		_, err = subFile.ReadAt(subq, int64(i)*int64(scsi.CHANNEL_SIZE))
		if err != nil {
			panic(err)
		}
		sect.Sectors[i] = new(cd.Sector)
		copy(sect.Sectors[i].Data[:], data)
		copy(sect.Sectors[i].Sub.Qchannel[:], subq)
	}
	sect.Matched = true
}
