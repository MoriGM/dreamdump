package cd

func (track *Track) GetStartLBA() int32 {
	if index, ok := track.Indexs[0]; track.Type == TRACK_TYPE_AUDIO && ok {
		return index.LBA
	}
	return track.Lba
}
