package option

import "os"

const (
	DATA        = 0
	DATA_C2     = 2
	DATA_SUB    = 4
	DATA_C2_SUB = 6
	DATA_SUB_C2 = 7
)

type Option struct {
	Device      string
	Drive       *os.File
	SectorOrder int
	CutOff      int32
	ReadOffset  int16
	ImageName   string
}
