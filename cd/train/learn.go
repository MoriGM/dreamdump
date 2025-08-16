package train

import (
	"dreamdump/cd"
	"dreamdump/option"
)

const (
	TRAIN_DIRECTION_START = 0x01
	TRAIN_DIRECTION_END   = 0x02
	TRAIN_MAX_JUMP        = 0x10000
	TRAIN_MIN_JUMP        = 0x10
)

type Training struct {
	Direction int
	LBA       []int32
}

func Train(opt *option.Option, direction int) (Training, error) {
	training := Training{
		Direction: direction,
	}

	return training, nil
}

func (training *Training) Play(opt *option.Option, untilLBA int32) {
	for _, lba := range training.LBA {
		if training.Direction == TRAIN_DIRECTION_START && lba > untilLBA {
			break
		}
		if training.Direction == TRAIN_DIRECTION_END && lba < untilLBA {
			break
		}
		cd.ReadSector(opt, lba)
	}
}
