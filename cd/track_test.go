package cd_test

import (
	"testing"

	"dreamdump/cd"

	"gotest.tools/v3/assert"
)

func TestGetTrackTypeModeAudio(t *testing.T) {
	track := cd.Track{
		Type: cd.TRACK_TYPE_AUDIO,
	}

	meta := cd.TrackMeta{}

	assert.Equal(t, track.GetTrackType(&meta), "AUDIO")
}

func TestGetTrackTypeMode1(t *testing.T) {
	track := cd.Track{
		Type: cd.TRACK_TYPE_DATA,
	}

	meta := cd.TrackMeta{
		DataMode: cd.TRACK_TYPE_DATA_MODE1,
	}

	assert.Equal(t, track.GetTrackType(&meta), "MODE1/2352")
}

func TestGetTrackTypeMode2(t *testing.T) {
	track := cd.Track{
		Type: cd.TRACK_TYPE_DATA,
	}

	meta := cd.TrackMeta{
		DataMode: cd.TRACK_TYPE_DATA_MODE2,
	}

	assert.Equal(t, track.GetTrackType(&meta), "MODE2/2352")
}

func TestGetTrackTypeMode0(t *testing.T) {
	track := cd.Track{
		Type: cd.TRACK_TYPE_DATA,
	}

	meta := cd.TrackMeta{}

	assert.Equal(t, track.GetTrackType(&meta), "MODE0/2352")
}
