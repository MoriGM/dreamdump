package cd

import (
	"math"
)

type QToc struct {
	Tracks          map[uint8]*Track
	LastTrackNumber uint8
}

func QTocNew() *QToc {
	qtoc := new(QToc)
	qtoc.Tracks = make(map[uint8]*Track)
	return qtoc
}

func (qtoc *QToc) AddSector(qchannel *QChannel) {
	if !qchannel.CheckParity() {
		return
	}
	if track, ok := qtoc.Tracks[qchannel.TrackNumber()-1]; ok {
		track.LbaEnd = min(track.LbaEnd, qchannel.LBA())
	}
	if qchannel.TrackNumber() == 110 {
		if track, ok := qtoc.Tracks[qtoc.LastTrackNumber]; ok {
			track.LbaEnd = min(track.LbaEnd, qchannel.LBA())
		}
	}
	if track, ok := qtoc.Tracks[qchannel.TrackNumber()]; ok {
		if index, ok := track.Indexs[qchannel.IndexNumber()]; ok {
			index.LBA = min(qchannel.LBA(), index.LBA)
		} else {
			track.Indexs[qchannel.IndexNumber()] = &Index{
				LBA: qchannel.LBA(),
			}
		}

		if qchannel.IndexNumber() == 1 {
			track.Lba = min(qchannel.LBA(), track.Lba)
			track.Type = qchannel.TrackType()
		}

		return
	}
	lba := qchannel.LBA()
	if qchannel.IndexNumber() == 0 {
		lba = math.MaxInt32
	}
	qtoc.Tracks[qchannel.TrackNumber()] = &Track{
		Lba:    lba,
		LbaEnd: math.MaxInt32,
		Type:   qchannel.TrackType(),
		Indexs: map[uint8]*Index{qchannel.IndexNumber(): {
			LBA: qchannel.LBA(),
		}},
		TrackNumber: qchannel.TrackNumber(),
	}
	if qchannel.TrackNumber() != 110 {
		qtoc.LastTrackNumber = max(qchannel.TrackNumber(), qtoc.LastTrackNumber)
	}
}
