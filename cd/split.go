package cd

import (
	"dreamdump/scsi"
	"os"
)

func (dense *Dense) Split(qtoc *QToc) {
	testfile, err := os.OpenFile("test.bin", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	start := (297968 - DENSE_LBA_OFFSET) * scsi.SECTOR_DATA_SIZE
	end := (303918 - DENSE_LBA_OFFSET) * scsi.SECTOR_DATA_SIZE
	start += 52
	end += 52
	testfile.Write((*dense)[start:end])
	testfile.Close()
}
