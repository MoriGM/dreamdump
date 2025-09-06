package cd

import (
	"dreamdump/log"
	"dreamdump/option"
)

const (
	TOC_OFFSET            = 0x110
	TOC_SIZE              = 512
	TOC_OFFSET_END        = TOC_OFFSET + TOC_SIZE
	TOC_MAGIC_NUMBER      = 0
	TOC_MAGIC_NUMBER_SIZE = 4
	TOC_TRACK_START       = TOC_MAGIC_NUMBER_SIZE
	TOC_TRACK_END         = TOC_TRACK_SIZE
	TOC_TRACK_SIZE        = 392
	TOC_TRACK_COUNT       = 398
	TOC_FIRST_TRACK       = 3
)

func ParseToc(sector *CdSectorData) []*Track {
	gdToc := sector[TOC_OFFSET:TOC_OFFSET_END]
	if gdToc[0] != 'T' || gdToc[1] != 'O' || gdToc[2] != 'C' || gdToc[3] != '1' {
		panic("Header TOC IN LBA 45000 doesn't have a magic number TOC1")
	}

	trackAmount := gdToc[TOC_TRACK_COUNT]

	tracks := make([]*Track, ((trackAmount+1)-TOC_FIRST_TRACK)+1)

	trackNumber := TOC_FIRST_TRACK
	for i := TOC_TRACK_START; i < TOC_TRACK_SIZE; i += TOC_MAGIC_NUMBER_SIZE {
		if gdToc[i+3] == 0xff {
			break
		}
		lba := (int(gdToc[i])) | (int(gdToc[i+1]) << 8) | (int(gdToc[i+2]) << 16)
		trackType := (gdToc[i+3] >> 4) & TRACK_TYPE_DATA
		lba -= 300
		if trackNumber == TOC_FIRST_TRACK {
			lba += 150
		}
		if trackType == TRACK_TYPE_DATA {
			lba -= 75
		}
		track := new(Track)
		track.Lba = int32(lba)
		track.LbaEnd = option.DC_LBA_END
		track.Type = trackType
		track.TrackNumber = uint8(trackNumber)

		tracks[trackNumber-3] = track

		if trackNumber != TOC_FIRST_TRACK {
			tracks[trackNumber-3-1].LbaEnd = int32(lba)
		}

		trackNumber++
	}

	track := new(Track)
	track.Lba = int32(option.DC_LBA_END)
	track.LbaEnd = option.DC_END
	track.Type = TRACK_TYPE_AUDIO
	track.TrackNumber = 110
	tracks[trackNumber-3] = track

	return tracks
}

func PrintToc(tracks []*Track) {
	log.Println("final TOC:")
	for _, track := range tracks {
		trackType := " data"
		if track.Type == TRACK_TYPE_AUDIO {
			trackType = "audio"
		}
		log.Printf("  track %d { %s }\n", track.TrackNumber, trackType)
		log.Printf("    { LBA: [% 7d ..% 6d]}\n", track.Lba, track.LbaEnd-1)
	}
}
