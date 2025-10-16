package cd

const (
	SECTOR_DATA_MODE                    = 0x0F
	SECTOR_DATA_MODE2_SUBMODE           = 0x12
	SECTOR_DATA_MODE2_SUBMODE_MASK_FORM = 0b00100000
)

func (data *CdSectorData) GetDataMode() uint8 {
	return data[SECTOR_DATA_MODE]
}

func (data *CdSectorData) GetDataModeForm() uint8 {
	if data.GetDataMode() == TRACK_TYPE_DATA_MODE1 {
		return 0
	}

	if (data[SECTOR_DATA_MODE2_SUBMODE] & SECTOR_DATA_MODE2_SUBMODE_MASK_FORM) == 0 {
		return TRACK_TYPE_DATA_MODE2_FORM1
	}

	return TRACK_TYPE_DATA_MODE2_FORM2
}
