package cd

import (
	"fmt"
	"os"
	"path/filepath"

	"dreamdump/encoding/msf"
	"dreamdump/option"
)

func WriteCue(opt *option.Option, qtoc *QToc, metas map[uint8]TrackMeta) {
	cueFileName := opt.PathName + "/" + opt.ImageName + ".cue"
	cueFile, err := os.OpenFile(cueFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o744)
	if err != nil {
		panic(err)
	}
	defer cueFile.Close()

	for _, trackNumber := range qtoc.TrackNumbers {
		meta := metas[trackNumber]
		track := qtoc.Tracks[trackNumber]
		fileLine := fmt.Sprintf("FILE \"%s\" BINARY\n", filepath.Base(meta.FileName))
		_, err = cueFile.WriteString(fileLine)
		if err != nil {
			panic(err)
		}
		trackLine := fmt.Sprintf("  TRACK %02d %s\n", trackNumber, getTrackType(track, &meta))
		_, err = cueFile.WriteString(trackLine)
		if err != nil {
			panic(err)
		}
		lba := track.Lba
		for _, indexNumber := range track.IndexNumbers {
			index := track.Indexs[indexNumber]
			indexOffset := index.Lba - lba
			if indexOffset < 0 {
				continue
			}
			indexLine := fmt.Sprintf("    INDEX %02d %s\n", indexNumber, msf.Encode(indexOffset))
			_, err = cueFile.WriteString(indexLine)
			if err != nil {
				panic(err)
			}
			lba = index.Lba
		}
	}
}

func getTrackType(track *Track, meta *TrackMeta) string {
	if track.Type == TRACK_TYPE_AUDIO {
		return "AUDIO"
	} else if meta.DataType == TRACK_TYPE_DATA_MODE1 {
		return "MODE1/2352"
	} else if meta.DataType == TRACK_TYPE_DATA_MODE2 {
		return "MODE2/2352"
	} else {
		return "MODE0/2352"
	}
}
