package sections

import (
	"os"
	"strconv"

	"dreamdump/cd"
	"dreamdump/exit_codes"
	"dreamdump/log"
	"dreamdump/option"
)

func TrainStart(opt *option.Option) {
	log.Println("Train drive start:")
	for lba := 55000; lba > int(option.DC_START); lba -= int(opt.ReadAtOnce) {
		_, err := cd.ReadSectors(opt, int32(lba), opt.ReadAtOnce)
		if err != nil {
			log.Println("Error when training start")
			os.Exit(exit_codes.ERROR_TRAINING_START)
		}
		log.PrintClean("Sector read " + strconv.FormatInt(int64(lba), 10))
	}
	log.Println("Training drive start finished")
	log.Println()
}
