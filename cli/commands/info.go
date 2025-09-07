package commands

import (
	"dreamdump/cd"
	"dreamdump/option"
)

func info(opt *option.Option, trackMetas map[uint8]cd.TrackMeta, toc []*cd.Track, qtoc *cd.QToc) {
	cd.WriteCue(opt, qtoc, trackMetas)
	cd.WriteGdi(opt, qtoc, trackMetas)
	cd.PrintXMLHashes(toc, trackMetas)
}
