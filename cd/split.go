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

func (dense *Dense) QTocSplit(opt *option.Option, qtoc *QToc) map[uint8]TrackMeta {
	offsetManager := dense.NewOffsetManager(option.DC_START)
	trackMetas := make(map[uint8]TrackMeta, 0)
	for _, trackNumber := range qtoc.TrackNumbers {
		track := qtoc.Tracks[trackNumber]

		trackMetas[trackNumber] = dense.SplitTrack(opt, track, offsetManager)
	}
	return trackMetas
}

func (dense *Dense) TocSplit(opt *option.Option, tracks []*Track) map[uint8]TrackMeta {
	offsetManager := dense.NewOffsetManager(option.DC_START)
	trackMetas := make(map[uint8]TrackMeta, 0)
	for _, track := range tracks {
		if track.TrackNumber == 110 {
			break
		}
		trackMetas[track.TrackNumber] = dense.SplitTrack(opt, track, offsetManager)
	}
	return trackMetas
}

func (dense *Dense) SplitTrack(opt *option.Option, track *Track, offsetManager *OffsetManager) TrackMeta {
	trackFileName := opt.PathName + "/" + opt.ImageName + " (Track " + strconv.Itoa(int(track.TrackNumber)) + ").bin"
	trackFile, err := os.OpenFile(trackFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
	dataType := uint8(0)
	if err != nil {
		panic(err)
	}
	var data []byte
	if track.Type == TRACK_TYPE_AUDIO {
		trackStartSize := (track.GetStartLBA()-option.DC_START)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
		trackEndSize := (min(track.LbaEnd, option.DC_END)-option.DC_START)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
		data = (*dense)[trackStartSize:trackEndSize]
	} else {
		data = make([]byte, 0)
		for lba := (track.GetStartLBA()) - option.DC_START; lba < min(track.LbaEnd, option.DC_LBA_END)-option.DC_START; lba++ {
			lbaStartSize := lba*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
			lbaEndSize := (lba+1)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
			var cdSectorData CdSectorData
			copy(cdSectorData[0:scsi.SECTOR_DATA_SIZE], (*dense)[lbaStartSize:lbaEndSize])
			var descrambledData []byte
			if cdSectorData.HasSyncHeader() {
				cdSectorData.Descramble()
				dataType |= cdSectorData.GetDataMode()
				descrambledData = make([]byte, scsi.SECTOR_DATA_SIZE)
				copy(descrambledData[:], cdSectorData[:])
			} else {
				descrambledData = make([]byte, scsi.SECTOR_DATA_SIZE)
			}

			data = append(data, descrambledData...)
		}
	}

	_, err = trackFile.Write(data)
	if err != nil {
		panic(err)
	}

	return TrackMeta{
		TrackNumber: track.TrackNumber,
		FileName:    trackFileName,
		Size:        uint32(len(data)),
		CRC32:       crc32.ChecksumIEEE(data),
		MD5:         md5.Sum(data),
		SHA1:        sha1.Sum(data),
		DataType:    dataType,
	}
}

func ZeroSector(count int32) []byte {
	data := make([]byte, scsi.SECTOR_DATA_SIZE*count)
	for pos := range scsi.SECTOR_DATA_SIZE * count {
		data[pos] = 0
	}
	return data
}
