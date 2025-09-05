//go:build linux

package sgio

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"dreamdump/scsi/parse"
)

func TestUnitReady(f *os.File) error {
	senseBuf := make([]byte, SENSE_BUF_LEN)
	inqCmdBlk := []uint8{0, 0, 0, 0, 0, 0}
	ioHdr := &SgIoHdr{
		InterfaceID:    int32('S'),
		CmdLen:         uint8(len(inqCmdBlk)),
		MxSbLen:        SENSE_BUF_LEN,
		DxferDirection: SG_DXFER_FROM_DEV,
		Cmdp:           &inqCmdBlk[0],
		Sbp:            &senseBuf[0],
		Timeout:        TIMEOUT_20_SECS,
	}

	err := SgioSyscall(f, ioHdr)
	if err != nil {
		return err
	}

	err = CheckSense(ioHdr, &senseBuf)
	if err != nil {
		return err
	}

	return nil
}

func CheckSense(i *SgIoHdr, s *[]byte) error {
	var b bytes.Buffer
	if (i.Info & SG_INFO_OK_MASK) != SG_INFO_OK {
		_, err := b.WriteString(
			fmt.Sprintf("\nSCSI response not ok\n"+
				"SCSI status: %v host status: %v driver status: %v",
				i.Status, i.HostStatus, i.DriverStatus))
		if err != nil {
			return err
		}
		if i.SbLenWr > 0 {
			_, err := b.WriteString(
				fmt.Sprintf("\nSENSE:\n%s",
					parse.DumpHex(*s)))
			if err != nil {
				return err
			}
		}
		return errors.New(b.String())
	}
	return nil
}

func SgioSyscall(f *os.File, i *SgIoHdr) error {
	return ioctl(f.Fd(), SG_IO, uintptr(unsafe.Pointer(i)))
}

func ioctl(fd, cmd, ptr uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptr)
	if err != 0 {
		return err
	}
	return nil
}

func OpenScsiDevice(fname string) (*os.File, error) {
	f, err := os.OpenFile(fname, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	var version uint32
	if (ioctl(f.Fd(), SG_GET_VERSION_NUM, uintptr(unsafe.Pointer(&version))) != nil) || (version < 30000) {
		return nil, errors.New("device does not appear to be an sg device")
	}
	return f, nil
}
