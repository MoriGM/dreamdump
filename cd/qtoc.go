package cd

import (
	"math"
	"slices"

	"dreamdump/log"
)

type QToc struct {
	Tracks          map[uint8]*Track
	TrackNames      []uint8
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

		track.Lba = max(min(qchannel.LBA(), track.Lba), DENSE_LBA_START)
		track.Type = qchannel.TrackType()

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
	for _, trackNumber := range qtoc.Tracks {
		trackKeys = append(trackKeys, trackNumber.TrackNumber)
	}
	slices.Sort(trackKeys)
	tracks := make(map[uint8]*Track)
	for _, key := range trackKeys {
		tracks[key] = qtoc.Tracks[key]
	}
	qtoc.Tracks = tracks
	qtoc.TrackNames = trackKeys
}

func (qtoc *QToc) Print() {
	log.Println("final QTOC:")
	for trackKeyPos, trackKey := range qtoc.TrackNames {
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
			endLBA := DENSE_LBA_END
			if nextIndex, ok := track.Indexs[indexKey+1]; ok {
				endLBA = int(nextIndex.Lba) - 1
			} else {
				if len(qtoc.TrackNames) > int(trackKeyPos)+1 {
					if nextTrack, ok := qtoc.Tracks[qtoc.TrackNames[trackKeyPos+1]]; ok {
						endLBA = int(nextTrack.Lba) - 1
					}
				}
			}

			log.Printf("    index %02d { LBA: [% 7d ..% 6d]}\n", indexKey, startLBA, endLBA)
		}
	}
}
