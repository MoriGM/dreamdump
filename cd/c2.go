package cd

func (sec Sector) HasC2() uint16 {
	c2BitCount := uint16(0)
	for _, c2 := range sec.C2 {
		upperSample := c2 & 0xF0
		lowerSample := c2 & 0x0F

		if upperSample != 0 {
			c2BitCount += 4
		}

		if lowerSample != 0 {
			c2BitCount += 4
		}
	}

	return c2BitCount
}
