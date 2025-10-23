package cd

import (
	"fmt"
	"os"
	"path/filepath"

	"dreamdump/encoding/msf"
	"dreamdump/log"
	"dreamdump/option"
)

func GenerateCue(opt *option.Option, qtoc *QToc, metas map[uint8]TrackMeta) {
	cueFileName := opt.ImageName + ".cue"
	cueFileNamePath := opt.PathName + "/" + cueFileName
	log.Printf("CUE [%s]\n", cueFileName)
	cueFile, err := os.OpenFile(cueFileNamePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer cueFile.Close()

	for _, trackNumber := range qtoc.TrackNumbers {
		meta := metas[trackNumber]
		track := qtoc.Tracks[trackNumber]
		fileLine := fmt.Sprintf("FILE \"%s\" BINARY\n", filepath.Base(meta.FileName))
		log.Print(fileLine)
		_, err = cueFile.WriteString(fileLine)
		if err != nil {
			panic(err)
		}
		trackLine := fmt.Sprintf("  TRACK %02d %s\n", trackNumber, getTrackType(track, &meta))
		log.Print(trackLine)
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
			log.Print(indexLine)
			_, err = cueFile.WriteString(indexLine)
			if err != nil {
				panic(err)
			}
			lba = index.Lba
		}
	}
}

func getTrackType(track *Track, meta *TrackMeta) string {
	switch track.Type {
	case TRACK_TYPE_AUDIO:
		return "AUDIO"
	case TRACK_TYPE_DATA_MODE1:
		return "MODE1/2352"
	case TRACK_TYPE_DATA_MODE2:
		return "MODE2/2352"
	}
	return "MODE0/2352"
}
