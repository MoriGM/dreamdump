package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"dreamdump/cd"
	"dreamdump/cd/sections"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/scsi/scsi_commands"
)

func DreamDumpDisc(opt *option.Option) {
	_, err := os.Stat(opt.PathName)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(opt.PathName, 0o744)
		if err != nil {
			panic(err)
		}
	}
	sectionMap := sections.GetSectionMap(opt)
	if opt.Speed > 0 {
		scsi_commands.SetCDSpeed(opt)
		log.Println("Set Read Speed to:" + strconv.FormatInt(int64(opt.Speed), 10) + " kbs")
		log.Println()
	}
	sections.ReadSections(opt, &sectionMap)
	dense := sections.ExtractSectionsToDense(opt, &sectionMap)
	qtoc := sections.ExtractSectionsToQtoc(&sectionMap)
	offsetManager := dense.NewOffsetManager(option.DC_START)
	specialSector := dense.GetLBA(offsetManager, option.DC_LBA_START)
	specialSector.Descramble()
	toc := cd.ParseToc(specialSector)
	log.Println()
	fmt.Printf("Write Offset: %d\n", offsetManager.SampleOffset)
	log.Println()
	cd.PrintToc(toc)
	log.Println()
	qtoc.Print()
	log.Println()
	trackMetas := dense.QTocSplit(opt, qtoc)
	cd.WriteCue(opt, qtoc, trackMetas)
	cd.WriteGdi(opt, qtoc, trackMetas)
	for _, trackNumber := range qtoc.TrackNumbers {
		trackMeta := trackMetas[trackNumber]
		romVaultLine := fmt.Sprintf("<rom name=\"%s\" size=\"%d\" crc=\"%x\" md5=\"%x\" sha1=\"%x\" />", filepath.Base(trackMeta.FileName), trackMeta.Size, trackMeta.CRC32, trackMeta.MD5, trackMeta.SHA1)
		log.Println(romVaultLine)
	}
}
