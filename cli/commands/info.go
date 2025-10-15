package commands

import (
	"dreamdump/cd"
	"dreamdump/log"
	"dreamdump/option"
)

func info(opt *option.Option, trackMetas map[uint8]cd.TrackMeta, toc []*cd.Track, qtoc *cd.QToc, headerSector *cd.CdSectorData) {
	cd.WriteCue(opt, qtoc, trackMetas)
	cd.WriteGdi(opt, qtoc, trackMetas)
	cd.PrintXMLHashes(toc, trackMetas)
	log.Println()
	cd.PrintTrackMeta(toc, trackMetas)
	log.Println()
	cd.PrintHeader(headerSector)
}
