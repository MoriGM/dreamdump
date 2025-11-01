//go:build linux

package sgio

import (
	"errors"
	"os"
	"syscall"
	"unsafe"
)

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
	f, err := os.OpenFile(fname, syscall.O_RDWR|syscall.O_NONBLOCK, 0)
	if err != nil {
		return nil, err
	}
	var version uint32
	if (ioctl(f.Fd(), SG_GET_VERSION_NUM, uintptr(unsafe.Pointer(&version))) != nil) || (version < 30000) {
		return nil, errors.New("device does not appear to be an sg device")
	}
	return f, nil
}
