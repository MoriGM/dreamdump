package train

import (
	"dreamdump/cd"
	"dreamdump/cd/sections"
	"dreamdump/log"
	"dreamdump/option"
	"os"
)

const (
	TRAIN_DIRECTION_START           = 0x01
	TRAIN_DIRECTION_END             = 0x02
	TRAIN_DIRECTION_START_LBA int32 = 0x100000
	TRAIN_DIRECTION_END_LBA   int32 = 0x300000
	TRAIN_MAX_JUMP            int32 = 0x10000
	TRAIN_MIN_JUMP            int32 = 0x10
)

type Training struct {
	Direction int
	LBA       []int32
}

func offsetDirection(direction int, offset int32) int32 {
	if direction == TRAIN_DIRECTION_START {
		return -offset
	}
	return offset
}

func Train(opt *option.Option, direction int) (Training, error) {
	training := Training{
		Direction: direction,
		LBA:       []int32{},
	}

	last_sector := TRAIN_DIRECTION_END_LBA
	if direction == TRAIN_DIRECTION_START {
		last_sector = TRAIN_DIRECTION_START_LBA
	}

	_, err := cd.ReadSector(opt, int32(last_sector))
	if err != nil {
		log.WriteLN("Cannot read inital train sector")
		os.Exit(2)
	}
	training.LBA = append(training.LBA, last_sector)

	offsetTimer := TRAIN_MAX_JUMP

	for {
		offset := offsetTimer
		next_sector := last_sector + offsetDirection(direction, offset)
		if next_sector > sections.DC_END || next_sector < sections.DC_START {
			break
		}
		_, err := cd.ReadSector(opt, int32(next_sector))
		if err != nil {
			offsetTimer = offset >> 8
			continue
		}
		last_sector = next_sector
		training.LBA = append(training.LBA, last_sector)

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
		_, err := cd.ReadSector(opt, lba)
		if err != nil {
			log.WriteLN("Error while playing the trained list of lba's")
			os.Exit(3)
		}
	}
}
