package commands

import (
	"dreamdump/cd"
	"dreamdump/log"
	"dreamdump/option"
)

func info(opt *option.Option, trackMetas map[uint8]cd.TrackMeta, toc []*cd.Track, qtoc *cd.QToc, headerSector *cd.CdSectorData) {
	if opt.QTocSplit {
		cd.GenerateCueByQToc(opt, qtoc, trackMetas)
	} else {
		cd.GenerateCueByToc(opt, toc, trackMetas)
	}
	log.Println()
	cd.GenerateGdi(opt, qtoc, trackMetas)
	log.Println()
	cd.PrintXMLHashes(toc, trackMetas)
	log.Println()
	cd.PrintTrackMeta(toc, trackMetas)
	log.Println()
	cd.PrintHeader(headerSector)
}
