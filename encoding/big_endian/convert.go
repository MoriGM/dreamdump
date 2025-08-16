package bigendian

import "encoding/binary"

func Int32(num int32) int32 {
	var bytes [4]byte
	binary.NativeEndian.PutUint32(bytes[:], uint32(num))
	return int32(binary.BigEndian.Uint32(bytes[:]))
}

func Uint32(num uint32) uint32 {
	var bytes [4]byte
	binary.NativeEndian.PutUint32(bytes[:], num)
	return binary.BigEndian.Uint32(bytes[:])
}
