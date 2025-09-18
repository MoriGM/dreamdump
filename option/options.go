package option

const (
	DATA        = 0
	DATA_C2     = 2
	DATA_SUB    = 4
	DATA_C2_SUB = 6
	DATA_SUB_C2 = 7
)

const (
	DC_START          int32 = 44990
	DC_LBA_START      int32 = 45000
	DC_END            int32 = 549152
	DC_LBA_END        int32 = 549150
	DC_INTERVAL       int32 = 10289
	DC_DEFAULT_CUTOFF int32 = 446261
)

type Option struct {
	Device      string
	Drive       any
	SectorOrder int
	CutOff      int32
	ReadOffset  int16
	Speed       uint16
	ImageName   string
	PathName    string
	QTocSplit   bool
	Train       bool
	ReadAtOnce  uint8
}
