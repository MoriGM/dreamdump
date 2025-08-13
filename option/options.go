package option

import "os"

type Option struct {
	Device      string
	Drive       *os.File
	SectorOrder int
	CutOff      int32
	ReadOffset  int32
}
