package sections

import (
	"os"

	"dreamdump/option"
)

func (sect *Section) WriteSection(opt *option.Option) {
	sect.WriteScramSection(opt)
	sect.WriteSubSection(opt)
}

func (sect *Section) WriteScramSection(opt *option.Option) {
	scramFileName := sect.FileName(opt) + ".scram"
	scramFile, err := os.OpenFile(scramFileName, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer scramFile.Close()
	for _, sector := range sect.Sectors {
		_, err = scramFile.Write(sector.Data[:])
		if err != nil {
			panic(err)
		}
	}
	err = scramFile.Sync()
	if err != nil {
		panic(err)
	}
}

func (sect *Section) WriteSubSection(opt *option.Option) {
	subFileName := sect.FileName(opt) + ".subq"
	subFile, err := os.OpenFile(subFileName, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer subFile.Close()
	for _, sector := range sect.Sectors {
		_, err = subFile.Write(sector.Sub.Qchannel[:])
		if err != nil {
			panic(err)
		}
	}
	err = subFile.Sync()
	if err != nil {
		panic(err)
	}
}

func (sect *Section) WriteHash(opt *option.Option) {
	hashFileName := sect.FileName(opt) + ".hash"
	hashFile, err := os.OpenFile(hashFileName, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer hashFile.Close()
	for _, hash := range sect.Hashes {
		_, err := hashFile.WriteString(hash + "\n")
		if err != nil {
			panic(err)
		}
	}
	err = hashFile.Sync()
	if err != nil {
		panic(err)
	}
}
