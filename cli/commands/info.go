package commands

import (
	"dreamdump/cd"
	"dreamdump/log"
	"dreamdump/option"
)

func info(opt *option.Option, trackMetas map[uint8]cd.TrackMeta, toc []*cd.Track, qtoc *cd.QToc, headerSector *cd.CdSectorData) {
	cd.GenerateCue(opt, qtoc, trackMetas)
	log.Println()
	cd.GenerateGdi(opt, qtoc, trackMetas)
	log.Println()
	cd.PrintXMLHashes(toc, trackMetas)
	log.Println()
	cd.PrintTrackMeta(toc, trackMetas)
	log.Println()
	cd.PrintHeader(headerSector)
}
