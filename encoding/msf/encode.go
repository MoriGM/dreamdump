package msf

import (
	"fmt"
)

const (
	MSF_FRAME  = 1
	MSF_SECOND = MSF_FRAME * 75
	MSF_MINUTE = MSF_SECOND * 60
)

func Encode(msf int32) string {
	minute := msf / MSF_MINUTE
	second := (msf % MSF_MINUTE) / MSF_SECOND
	frame := msf % MSF_SECOND
	return fmt.Sprintf("%02d:%02d:%02d", minute, second, frame)
}
