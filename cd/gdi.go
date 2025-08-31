package cd

import (
	"dreamdump/option"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

func WriteGdi(opt *option.Option, qtoc *QToc, metas map[uint8]TrackMeta) {
	cueFileName := opt.PathName + "/" + opt.ImageName + ".gdi"
	cueFile, err := os.OpenFile(cueFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o744)
	if err != nil {
		panic(err)
	}
	defer cueFile.Close()

	_, err = cueFile.WriteString(strconv.FormatInt(int64(slices.Max(qtoc.TrackNumbers)), 10) + "\n")
	if err != nil {
		panic(err)
	}
	_, err = cueFile.WriteString("1      0 4 2352 [fix] 0\n")
	if err != nil {
		panic(err)
	}
	_, err = cueFile.WriteString("2  [fix] 4 2352 [fix] 0\n")
	if err != nil {
		panic(err)
	}

	for _, trackNumber := range qtoc.TrackNumbers {
		meta := metas[trackNumber]
		track := qtoc.Tracks[trackNumber]
		trackLine := fmt.Sprintf("%d% 7d ", trackNumber, track.Lba)
		if track.Type == TRACK_TYPE_AUDIO {
			trackLine += "0"
		} else {
			trackLine += "4"
		}
		trackLine += fmt.Sprintf(" 2352 \"%s\" 0\n", filepath.Base(meta.FileName))
		_, err = cueFile.WriteString(trackLine)
		if err != nil {
			panic(err)
		}
	}
}
