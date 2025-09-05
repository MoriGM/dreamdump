package driver

import (
	"errors"

	"dreamdump/scsi/parse"
)

func CheckSense(status *Status) error {
	if status.Key == 0 && status.Asc == 0 && status.AscQ == 0 {
		return nil
	}
	return errors.New(parse.GetErrString(status.Asc, status.AscQ))
}
