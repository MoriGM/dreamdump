package cd

const SECTOR_DATA_MODE = 15

func (data *CdSectorData) GetDataMode() uint8 {
	return data[SECTOR_DATA_MODE]
}
