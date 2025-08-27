package cd

const (
	DENSE_LBA_OFFSET = 44990
	DENSE_LBA_START  = 45000
	DENSE_LBA_END    = 549151
)

// LBA 44990 is at the first byte
type Dense []byte
