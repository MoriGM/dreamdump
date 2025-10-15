package cd

import (
	"fmt"
	"path/filepath"

	"dreamdump/log"
)

func PrintXMLHashes(toc []*Track, trackMetas map[uint8]TrackMeta) {
	log.Println("*** HASH")
	log.Println("")
	log.Println("dat:")
	for _, tocTrack := range toc {
		if tocTrack.TrackNumber == 110 {
			break
		}
		trackMeta := trackMetas[tocTrack.TrackNumber]
		romVaultLine := fmt.Sprintf("<rom name=\"%s\" size=\"%d\" crc=\"%08x\" md5=\"%032x\" sha1=\"%040x\" />", filepath.Base(trackMeta.FileName), trackMeta.Size, trackMeta.CRC32, trackMeta.MD5, trackMeta.SHA1)
		log.Println(romVaultLine)
	}
}

func PrintTrackMeta(toc []*Track, trackMetas map[uint8]TrackMeta) {
	log.Println("")
	log.Println("*** INFO")
	for _, tocTrack := range toc {
		if tocTrack.TrackNumber == 110 || tocTrack.Type == TRACK_TYPE_AUDIO {
			break
		}
		trackMeta := trackMetas[tocTrack.TrackNumber]
		log.Printf("CD-ROM [%s]:\n", filepath.Base(trackMeta.FileName))
		log.Printf("  sector count: %d\n", trackMeta.Sectors)
		if trackMeta.InvalidSyncSectors > 0 {
			log.Printf("  invalid sync sectors: %d\n", trackMeta.InvalidSyncSectors)
		}
	}
}
