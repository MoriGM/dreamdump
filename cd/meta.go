package cd

import (
	"fmt"
	"path/filepath"

	"dreamdump/log"
)

func PrintXMLHashes(toc []*Track, trackMetas map[uint8]TrackMeta) {
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
