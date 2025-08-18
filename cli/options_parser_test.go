package cli_test

import (
	"os"
	"testing"

	"dreamdump/cli"

	"gotest.tools/v3/assert"
)

func TestFindArgumentStringDrive(t *testing.T) {
	os.Args = []string{"--drive=/dev/sr0"}
	drive := cli.FindArgumentString("drive")
	assert.Equal(t, *drive, "/dev/sr0")
}

func TestFindArgumentStringNotFound(t *testing.T) {
	os.Args = []string{"--drive=/dev/sr0"}
	drive := cli.FindArgumentString("device")
	if drive != nil {
		t.Error()
	}
}
