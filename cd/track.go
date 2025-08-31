package cd

import "dreamdump/option"

func (track *Track) GetStartLBA() int32 {
	if index, ok := track.Indexs[0]; ok {
		return max(index.Lba, option.DC_LBA_START)
	}
	return max(track.Lba, option.DC_LBA_START)
}
