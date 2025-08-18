package cd

func (c2 *CdSectorC2) Amount() uint16 {
	c2BitCount := uint16(0)
	for _, b := range c2 {
		upperSample := b & 0xF0
		lowerSample := b & 0x0F

		if upperSample != 0 {
			c2BitCount += 4
		}

		if lowerSample != 0 {
			c2BitCount += 4
		}
	}

	return c2BitCount
}
