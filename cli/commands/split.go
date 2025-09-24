package commands

import (
	"errors"
	"os"

	"dreamdump/cd"
	"dreamdump/cd/sections"
	"dreamdump/log"
	"dreamdump/option"
)

func split(opt *option.Option, dense *cd.Dense, qtoc *cd.QToc) (map[uint8]cd.TrackMeta, []*cd.Track, *cd.QToc, *cd.CdSectorData) {
	offsetManager := dense.NewOffsetManager(option.DC_START)

	headerSector := dense.GetLBA(offsetManager, option.DC_LBA_START)
	headerSector.Descramble()
	toc := cd.ParseToc(headerSector)

	log.Println()
	log.Printf("Write Offset: %d\n", offsetManager.SampleOffset)

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

	dense = nil

	return trackMetas, toc, qtoc, headerSector
}

func DreamDumpSplit(opt *option.Option) {
	_, err := os.Stat(opt.PathName)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(opt.PathName, 0o744)
		if err != nil {
			panic(err)
		}
	}
	sectionMap := sections.GetSectionMap(opt)
	sections.ReadFileSections(opt, sectionMap)
	for _, section := range sectionMap {
		if !section.Matched {
			log.Println("Not all sections are matching")
			return
		}
	}

	dense, qtoc := sections.ExtractSections(opt, sectionMap)
	trackMetas, toc, qtoc, headerSector := split(opt, dense, qtoc)
	info(opt, trackMetas, toc, qtoc, headerSector)
}
