package cd

import (
	"slices"

	"dreamdump/log"
	"dreamdump/option"
)

type QToc struct {
	Tracks          map[uint8]*Track
	TrackNumbers    []uint8
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
			index.Lba = min(qchannel.LBA(), index.Lba)
		} else {
			track.Indexs[qchannel.IndexNumber()] = &Index{
				Lba: qchannel.LBA(),
			}
		}

		track.Lba = max(min(qchannel.LBA(), track.Lba), option.DC_LBA_START)
		track.Type = qchannel.TrackType()

		return
	}
	lba := qchannel.LBA()
	qtoc.Tracks[qchannel.TrackNumber()] = &Track{
		Lba:    lba,
		LbaEnd: option.DC_END,
		Type:   qchannel.TrackType(),
		Indexs: map[uint8]*Index{qchannel.IndexNumber(): {
			Lba: qchannel.LBA(),
		}},
		TrackNumber: qchannel.TrackNumber(),
	}
	if qchannel.TrackNumber() != 110 {
		qtoc.LastTrackNumber = max(qchannel.TrackNumber(), qtoc.LastTrackNumber)
	}
}

func (qtoc *QToc) Sort() {
	trackKeys := []uint8{}
	for _, track := range qtoc.Tracks {
		if track.TrackNumber == 110 {
			continue
		}
		trackKeys = append(trackKeys, track.TrackNumber)
		indexKeys := []uint8{}
		for indexNumber := range track.Indexs {
			indexKeys = append(indexKeys, indexNumber)
		}
		slices.Sort(indexKeys[:])
		track.IndexNumbers = indexKeys
	}
	slices.Sort(trackKeys[:])
	qtoc.TrackNumbers = trackKeys
}

func (qtoc *QToc) Print() {
	log.Println("final QTOC:")
	for _, trackKey := range qtoc.TrackNumbers {
		track := qtoc.Tracks[trackKey]
		trackType := "data"
		if track.Type == TRACK_TYPE_AUDIO {
			trackType = "audio"
		}
		log.Printf("  track %d {  %s }\n", track.TrackNumber, trackType)
		indexKeys := []uint8{}
		for key := range track.Indexs {
			indexKeys = append(indexKeys, key)
		}
		slices.Sort(indexKeys)
		for _, indexKey := range indexKeys {
			index := track.Indexs[indexKey]
			startLBA := index.Lba
			endLBA := track.LbaEnd - 1
			if nextIndex, ok := track.Indexs[indexKey+1]; ok {
				endLBA = nextIndex.Lba - 1
			}

			log.Printf("    index %02d { LBA: [% 7d ..% 6d]}\n", indexKey, startLBA, min(endLBA, option.DC_LBA_END))
		}
	}
}
