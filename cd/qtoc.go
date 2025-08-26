package cd

import (
	"math"
)

type QToc struct {
	Tracks map[uint8]*Track
}

func QTocNew() QToc {
	qtoc := QToc{}
	qtoc.Tracks = make(map[uint8]*Track)
	return qtoc
}

func (qtoc *QToc) AddSector(qchannel *QChannel) {
	if track, ok := qtoc.Tracks[qchannel.TrackNumber()]; ok {
		if index, ok := track.Indexs[qchannel.IndexNumber()]; ok {
			index.LBA = min(qchannel.LBA(), index.LBA)
		} else {
			track.Indexs[qchannel.IndexNumber()] = &Index{
				LBA: qchannel.LBA(),
			}
		}

		if qchannel.IndexNumber() == 1 {
			track.LBA = min(qchannel.LBA(), track.LBA)
			track.Type = qchannel.TrackType()
		}

		return
	}
	lba := qchannel.LBA()
	if qchannel.IndexNumber() == 0 {
		lba = math.MaxInt32
	}
	qtoc.Tracks[qchannel.TrackNumber()] = &Track{
		LBA:  lba,
		Type: qchannel.TrackType(),
		Indexs: map[uint8]*Index{qchannel.IndexNumber(): {
			LBA: qchannel.LBA(),
		}},
	}
}
