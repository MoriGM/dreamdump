package cd

import (
	"bytes"
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
	trackMetas := make(map[uint8]TrackMeta, len(qtoc.TrackNumbers))
	for _, trackNumber := range qtoc.TrackNumbers {
		track := qtoc.Tracks[trackNumber]

		trackMetas[trackNumber] = dense.SplitTrack(opt, track, offsetManager)
	}
	return trackMetas
}

func (dense *Dense) TocSplit(opt *option.Option, tracks []*Track) map[uint8]TrackMeta {
	offsetManager := dense.NewOffsetManager(option.DC_START)
	trackMetas := make(map[uint8]TrackMeta, len(tracks)-1)
	for _, track := range tracks {
		if track.TrackNumber == 110 {
			break
		}
		trackMetas[track.TrackNumber] = dense.SplitTrack(opt, track, offsetManager)
	}
	return trackMetas
}

func (dense *Dense) SplitTrack(opt *option.Option, track *Track, offsetManager *OffsetManager) TrackMeta {
	if track.Type == TRACK_TYPE_DATA {
		return dense.splitData(opt, track, offsetManager)
	}
	return dense.splitAudio(opt, track, offsetManager)
}

func (dense *Dense) splitData(opt *option.Option, track *Track, offsetManager *OffsetManager) TrackMeta {
	trackFileName := opt.PathName + "/" + opt.ImageName + " (Track " + strconv.Itoa(int(track.TrackNumber)) + ").bin"
	trackFile, err := os.OpenFile(trackFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	dataType := uint8(0)

	crc32Sum := crc32.NewIEEE()
	md5Sum := md5.New()
	sha1Sum := sha1.New()

	trackStartSize := (track.GetStartLBA()-option.DC_START)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
	trackEndSize := (min(track.LbaEnd, option.DC_LBA_END)-option.DC_START)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset

	var cdSectorData CdSectorData
	var descrambledData bytes.Buffer
	for lba := (track.GetStartLBA()) - option.DC_START; lba < min(track.LbaEnd, option.DC_LBA_END)-option.DC_START; lba++ {
		descrambledData.Reset()

		lbaStartSize := lba*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
		lbaEndSize := (lba+1)*scsi.SECTOR_DATA_SIZE + offsetManager.ByteOffset
		copy(cdSectorData[0:scsi.SECTOR_DATA_SIZE], (*dense)[lbaStartSize:lbaEndSize])
		if cdSectorData.HasSyncHeader() {
			cdSectorData.Descramble()
			dataType |= cdSectorData.GetDataMode()
			descrambledData.Write(cdSectorData[:])
		} else {
			descrambledData.Write(make([]byte, scsi.SECTOR_DATA_SIZE))
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

	trackMeta := TrackMeta{
		TrackNumber: track.TrackNumber,
		FileName:    trackFileName,
		Size:        uint32(trackEndSize - trackStartSize),
		CRC32:       crc32Sum.Sum32(),
		MD5:         [16]byte(md5Sum.Sum(nil)),
		SHA1:        [20]byte(sha1Sum.Sum(nil)),
		DataType:    dataType,
	}

	return trackMeta
}

func (dense *Dense) splitAudio(opt *option.Option, track *Track, offsetManager *OffsetManager) TrackMeta {
	trackFileName := opt.PathName + "/" + opt.ImageName + " (Track " + strconv.Itoa(int(track.TrackNumber)) + ").bin"
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
		DataType:    0,
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	return trackMeta
}
