package cd

import (
	"crypto/md5"
	"crypto/sha1"
	"hash/crc32"
	"os"
	"strconv"

	"dreamdump/option"
	"dreamdump/scsi"
)

func (dense *Dense) Split(opt *option.Option, qtoc *QToc) map[int]TrackMeta {
	offsetManager := dense.NewOffsetManager(DENSE_LBA_OFFSET)
	trackMetas := make(map[int]TrackMeta)
	for trackNumber, track := range qtoc.Tracks {
		if trackNumber == 110 {
			continue
		}
		trackFileName := opt.PathName + "/" + opt.ImageName + " (Track " + strconv.Itoa(int(trackNumber)) + ").bin"
		trackStartSize := (track.GetStartLBA()-DENSE_LBA_OFFSET)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
		trackEndSize := (min(track.LbaEnd, DENSE_LBA_END)-DENSE_LBA_OFFSET)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
		trackSize := trackEndSize - trackStartSize
		trackFile, err := os.OpenFile(trackFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			panic(err)
		}
		var data []byte
		if track.Type == TRACK_TYPE_AUDIO {
			data = (*dense)[trackStartSize:trackEndSize]
		} else {
			data = make([]byte, 0)
			counter := 0
			for lba := track.GetStartLBA() - DENSE_LBA_OFFSET; lba < min(track.LbaEnd, DENSE_LBA_END)-DENSE_LBA_OFFSET; lba++ {
				lbaStartSize := lba*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
				lbaEndSize := (lba+1)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
				var cdSectorData CdSectorData
				copy(cdSectorData[0:scsi.SECTOR_DATA_SIZE], (*dense)[lbaStartSize:lbaEndSize])
				cdSectorData.Descramble()
				descrambledData := make([]byte, scsi.SECTOR_DATA_SIZE)
				copy(descrambledData[:], cdSectorData[:])
				data = append(data, descrambledData...)
				counter++
			}
		}

		_, err = trackFile.Write(data)
		if err != nil {
			panic(err)
		}

		trackMetas[int(trackNumber)] = TrackMeta{
			FileName: trackFileName,
			Size:     uint32(trackSize),
			CRC32:    crc32.ChecksumIEEE(data),
			MD5:      md5.Sum(data),
			SHA1:     sha1.Sum(data),
		}
	}
	return trackMetas
}
