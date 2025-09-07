package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"dreamdump/cd"
	"dreamdump/cd/sections"
	"dreamdump/drive/setup"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/scsi/scsi_commands"
)

func DreamDumpDisc(opt *option.Option) {
	setup.InitializeDrive(opt)
	_, err := os.Stat(opt.PathName)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(opt.PathName, 0o744)
		if err != nil {
			panic(err)
		}
	}

	log.Println()
	if opt.Speed > 0 {
		scsi_commands.SetCDSpeed(opt)
		log.Println("Set Read Speed to:" + strconv.FormatInt(int64(opt.Speed), 10) + " kbs")
		log.Println()
	}

	sectionMap := sections.GetSectionMap(opt)
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

	var trackMetas map[uint8]cd.TrackMeta
	if opt.QTocSplit {
		trackMetas = dense.QTocSplit(opt, qtoc)
	} else {
		trackMetas = dense.TocSplit(opt, toc)
	}

	cd.WriteCue(opt, qtoc, trackMetas)
	cd.WriteGdi(opt, qtoc, trackMetas)
	cd.PrintXMLHashes(toc, trackMetas)
}
