package cd

func (track *Track) GetStartLBA() int32 {
	if index, ok := track.Indexs[0]; ok {
		return max(index.Lba, DENSE_LBA_START)
	}
	return max(track.Lba, DENSE_LBA_START)
}
