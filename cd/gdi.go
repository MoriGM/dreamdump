package cd

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	"dreamdump/log"
	"dreamdump/option"
)

func GenerateGdi(opt *option.Option, qtoc *QToc, metas map[uint8]TrackMeta) {
	cueFileName := opt.ImageName + ".gdi"
	cueFileNamePath := opt.PathName + "/" + cueFileName
	cueFile, err := os.OpenFile(cueFileNamePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
	log.Printf("GDI [%s]\n", cueFileName)
	if err != nil {
		panic(err)
	}
	defer cueFile.Close()

	trackAmountLine := strconv.FormatInt(int64(slices.Max(qtoc.TrackNumbers)), 10) + "\n"
	_, err = cueFile.WriteString(trackAmountLine)
	log.Print(trackAmountLine)
	if err != nil {
		panic(err)
	}

	zeroPadding := ""
	if qtoc.LastTrackNumber > 9 {
		zeroPadding = "0"
	}

	firstTrackLine := fmt.Sprintf("%s1       0 4 2352 [fix] 0\n", zeroPadding)
	_, err = cueFile.WriteString(firstTrackLine)
	log.Print(firstTrackLine)
	if err != nil {
		panic(err)
	}
	secondTrackLine := fmt.Sprintf("%s2   [fix] 0 2352 [fix] 0\n", zeroPadding)
	_, err = cueFile.WriteString(secondTrackLine)
	log.Print(secondTrackLine)
	if err != nil {
		panic(err)
	}

	for _, trackNumber := range qtoc.TrackNumbers {
		meta := metas[trackNumber]
		track := qtoc.Tracks[trackNumber]
		var trackLine string
		if qtoc.LastTrackNumber > 9 {
			trackLine = fmt.Sprintf("%02d % 7d ", trackNumber, track.Lba)
		} else {
			trackLine = fmt.Sprintf("%d % 7d ", trackNumber, track.Lba)
		}
		if track.Type == TRACK_TYPE_AUDIO {
			trackLine += "0"
		} else {
			trackLine += "4"
		}
		trackLine += fmt.Sprintf(" 2352 \"%s\" 0\n", filepath.Base(meta.FileName))
		_, err = cueFile.WriteString(trackLine)
		log.Print(trackLine)
		if err != nil {
			panic(err)
		}
	}
}
