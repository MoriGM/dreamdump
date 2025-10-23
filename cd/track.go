package cd

import "dreamdump/option"

func (track *Track) GetStartLBA() int32 {
	if index, ok := track.Indexs[0]; ok {
		return max(index.Lba, option.DC_LBA_START)
	}
	return max(track.Lba, option.DC_LBA_START)
}

func (track *Track) GetTrackType(meta *TrackMeta) string {
	if track.Type == TRACK_TYPE_AUDIO {
		return "AUDIO"
	} else if meta.DataMode == TRACK_TYPE_DATA_MODE1 {
		return "MODE1/2352"
	} else if meta.DataMode == TRACK_TYPE_DATA_MODE2 {
		return "MODE2/2352"
	} else {
		return "MODE0/2352"
	}
}
