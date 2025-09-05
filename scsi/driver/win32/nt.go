//go:build windows

package win32

import (
	"reflect"
	"unsafe"

	"golang.org/x/sys/windows"
)

func ToByPo(data any) *byte {
	return (*byte)(unsafe.Pointer(reflect.ValueOf(data).Pointer()))
}

func OpenScsiDevice(fname string) (windows.Handle, error) {
	deviceName := "//./" + fname + ":"
	utf16DeviceName, err := windows.UTF16FromString(deviceName)
	if err != nil {
		panic(err)
	}
	handle, err := windows.CreateFile(&utf16DeviceName[0], windows.GENERIC_READ|windows.GENERIC_WRITE, windows.FILE_SHARE_READ, nil, windows.OPEN_EXISTING, 0, 0)
	return handle, err
}

func ScsiCall(handle windows.Handle, sptd_sd *SPTD_SD) error {
	size := uint32(unsafe.Sizeof(*sptd_sd))
	var returnLen uint32
	err := windows.DeviceIoControl(handle, SCSI_PASS_THROUGH_DIRECT, ToByPo(sptd_sd), size, ToByPo(sptd_sd), size, &returnLen, nil)
	return err
}
