package sections

import (
	"os"

	"dreamdump/option"
)

func (sect *Section) WriteSection(opt *option.Option) {
	sect.WriteBinSection(opt)
	sect.WriteSubSection(opt)
}

func (sect *Section) WriteBinSection(opt *option.Option) {
	binFile, err := os.OpenFile(sect.FileName(opt)+".bin", os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer binFile.Close()
	for _, sector := range sect.Sectors {
		_, err = binFile.Write(sector.Data[:])
		if err != nil {
			panic(err)
		}
	}
	err = binFile.Sync()
	if err != nil {
		panic(err)
	}
}

func (sect *Section) WriteSubSection(opt *option.Option) {
	subFile, err := os.OpenFile(sect.FileName(opt)+".subq", os.O_CREATE|os.O_WRONLY, 0o644)
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
