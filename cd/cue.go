package cd

import (
	"fmt"
	"os"
	"path/filepath"

	"dreamdump/encoding/msf"
	"dreamdump/log"
	"dreamdump/option"
)

func GenerateCueByQToc(opt *option.Option, qtoc *QToc, metas map[uint8]TrackMeta) {
	cueFileName := opt.ImageName + ".cue"
	cueFileNamePath := opt.PathName + "/" + cueFileName
	log.Printf("CUE [%s]:\n", cueFileName)
	cueFile, err := os.OpenFile(cueFileNamePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		panic(err)
	}
	defer cueFile.Close()

	remText := "REM HIGH-DENSITY AREA\n"
	log.Print(remText)
	_, err = cueFile.WriteString(remText)
	if err != nil {
		panic(err)
	}

	for _, trackNumber := range qtoc.TrackNumbers {
		meta := metas[trackNumber]
		track := qtoc.Tracks[trackNumber]
		fileLine := fmt.Sprintf("FILE \"%s\" BINARY\n", filepath.Base(meta.FileName))
		log.Print(fileLine)
		_, err = cueFile.WriteString(fileLine)
		if err != nil {
			panic(err)
		}
		trackLine := fmt.Sprintf("  TRACK %02d %s\n", trackNumber, track.GetTrackType(&meta))
		log.Print(trackLine)
		_, err = cueFile.WriteString(trackLine)
		if err != nil {
			panic(err)
		}
		lba := track.Lba
		for _, indexNumber := range track.IndexNumbers {
			index := track.Indexes[indexNumber]
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

func GenerateCueByToc(opt *option.Option, tracks []*Track, metas map[uint8]TrackMeta) {
	cueFileName := opt.ImageName + ".cue"
	cueFileNamePath := opt.PathName + "/" + cueFileName
	log.Printf("CUE [%s]:\n", cueFileName)
	cueFile, err := os.OpenFile(cueFileNamePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		panic(err)
	}
	defer cueFile.Close()

	remText := "REM HIGH-DENSITY AREA\n"
	log.Print(remText)
	_, err = cueFile.WriteString(remText)
	if err != nil {
		panic(err)
	}

	for _, track := range tracks {
		if track.TrackNumber == 110 {
			continue
		}
		meta := metas[track.TrackNumber]
		fileLine := fmt.Sprintf("FILE \"%s\" BINARY\n", filepath.Base(meta.FileName))
		log.Print(fileLine)
		_, err = cueFile.WriteString(fileLine)
		if err != nil {
			panic(err)
		}
		trackLine := fmt.Sprintf("  TRACK %02d %s\n", track.TrackNumber, track.GetTrackType(&meta))
		log.Print(trackLine)
		_, err = cueFile.WriteString(trackLine)
		if err != nil {
			panic(err)
		}
		indexNumber := 0
		if track.TrackNumber == 3 {
			indexNumber = 1
		}
		indexLine := fmt.Sprintf("    INDEX %02d %s\n", indexNumber, msf.Encode(0))
		log.Print(indexLine)
		_, err = cueFile.WriteString(indexLine)
		if err != nil {
			panic(err)
		}
		if track.TrackNumber > 3 {
			offsetIndex := int32(msf.MSF_SECOND * 2)
			if track.Type == TRACK_TYPE_DATA {
				offsetIndex += int32(meta.InvalidSyncSectors)
			}
			indexLine = fmt.Sprintf("    INDEX %02d %s\n", 1, msf.Encode(offsetIndex))
			log.Print(indexLine)
			_, err = cueFile.WriteString(indexLine)
			if err != nil {
				panic(err)
			}
		}
	}
}
