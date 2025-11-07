package cd

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"os"

	"dreamdump/option"
	"dreamdump/scsi"
)

func (dense *Dense) QTocSplit(opt *option.Option, qtoc *QToc) map[uint8]TrackMeta {
	offsetManager := dense.NewOffsetManager(option.DC_START)
	trackMetas := make(map[uint8]TrackMeta, len(qtoc.TrackNumbers))

	var trackFileName string
	for _, trackNumber := range qtoc.TrackNumbers {
		track := qtoc.Tracks[trackNumber]
		if len(qtoc.TrackNumbers) > 9 {
			trackFileName = fmt.Sprintf("%s/%s (Track %02d).bin", opt.PathName, opt.ImageName, track.TrackNumber)
		} else {
			trackFileName = fmt.Sprintf("%s/%s (Track %d).bin", opt.PathName, opt.ImageName, track.TrackNumber)
		}
		trackMetas[trackNumber] = dense.SplitTrack(trackFileName, track, offsetManager)
	}
	return trackMetas
}

func (dense *Dense) TocSplit(opt *option.Option, tracks []*Track) map[uint8]TrackMeta {
	offsetManager := dense.NewOffsetManager(option.DC_START)
	trackMetas := make(map[uint8]TrackMeta, len(tracks)-1)

	var trackFileName string
	for _, track := range tracks {
		if track.TrackNumber == 110 {
			break
		}
		if len(tracks) > 9 {
			trackFileName = fmt.Sprintf("%s/%s (Track %02d).bin", opt.PathName, opt.ImageName, track.TrackNumber)
		} else {
			trackFileName = fmt.Sprintf("%s/%s (Track %d).bin", opt.PathName, opt.ImageName, track.TrackNumber)
		}
		trackMetas[track.TrackNumber] = dense.SplitTrack(trackFileName, track, offsetManager)
	}
	return trackMetas
}

func (dense *Dense) SplitTrack(trackFileName string, track *Track, offsetManager *OffsetManager) TrackMeta {
	if track.Type == TRACK_TYPE_DATA {
		return dense.splitData(trackFileName, track, offsetManager)
	}
	return dense.splitAudio(trackFileName, track, offsetManager)
}

func (dense *Dense) splitData(trackFileName string, track *Track, offsetManager *OffsetManager) TrackMeta {
	trackFile, err := os.OpenFile(trackFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	dataMode := uint8(0)

	crc32Sum := crc32.NewIEEE()
	md5Sum := md5.New()
	sha1Sum := sha1.New()

	trackStartSize := (track.GetStartLBA()-option.DC_START)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
	trackEndSize := (min(track.LbaEnd, option.DC_LBA_END)-option.DC_START)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset

	var cdSectorData CdSectorData
	var descrambledData bytes.Buffer
	var invalidSyncSectors uint32
	var edc uint16
	for lba := (track.GetStartLBA()) - option.DC_START; lba < min(track.LbaEnd, option.DC_LBA_END)-option.DC_START; lba++ {
		descrambledData.Reset()

		lbaStartSize := lba*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
		lbaEndSize := (lba+1)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
		copy(cdSectorData[0:scsi.SECTOR_DATA_SIZE], (*dense)[lbaStartSize:lbaEndSize])
		if cdSectorData.HasSyncHeader() {
			cdSectorData.Descramble()
			dataMode |= cdSectorData.GetDataMode()
			descrambledData.Write(cdSectorData[:])
			if !cdSectorData.CheckEDC() {
				edc++
			}
		} else {
			descrambledData.Write(make([]byte, scsi.SECTOR_DATA_SIZE))
			invalidSyncSectors++
		}

		crc32Sum.Write(descrambledData.Bytes())
		md5Sum.Write(descrambledData.Bytes())
		sha1Sum.Write(descrambledData.Bytes())
		_, err = trackFile.Write(descrambledData.Bytes())
		if err != nil {
			panic(err)
		}
	}

	err = trackFile.Close()
	if err != nil {
		panic(err)
	}

	size := uint32(trackEndSize - trackStartSize)
	sectors := size / scsi.SECTOR_DATA_SIZE

	trackMeta := TrackMeta{
		TrackNumber:        track.TrackNumber,
		FileName:           trackFileName,
		Size:               size,
		Sectors:            sectors,
		CRC32:              crc32Sum.Sum32(),
		MD5:                [16]byte(md5Sum.Sum(nil)),
		SHA1:               [20]byte(sha1Sum.Sum(nil)),
		DataMode:           dataMode,
		InvalidSyncSectors: invalidSyncSectors,
		EDC:                edc,
	}

	return trackMeta
}

func (dense *Dense) splitAudio(trackFileName string, track *Track, offsetManager *OffsetManager) TrackMeta {
	file, err := os.OpenFile(trackFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	trackStartSize := (track.GetStartLBA()-option.DC_START)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
	trackEndSize := (min(track.LbaEnd, option.DC_LBA_END)-option.DC_START)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset

	_, err = file.Write((*dense)[trackStartSize:trackEndSize])
	if err != nil {
		panic(err)
	}
	trackMeta := TrackMeta{
		TrackNumber: track.TrackNumber,
		FileName:    trackFileName,
		Size:        uint32(trackEndSize - trackStartSize),
		CRC32:       crc32.ChecksumIEEE((*dense)[trackStartSize:trackEndSize]),
		MD5:         md5.Sum((*dense)[trackStartSize:trackEndSize]),
		SHA1:        sha1.Sum((*dense)[trackStartSize:trackEndSize]),
		DataMode:    0,
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	return trackMeta
}
