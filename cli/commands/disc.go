package commands

import (
	"errors"
	"os"
	"strconv"

	"dreamdump/cd/sections"
	"dreamdump/drive/setup"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/scsi/scsi_commands"
)

func DreamDumpDisc(opt *option.Option) {
	setup.InitializeDrive(opt)
	_, err := os.Stat(opt.PathName)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(opt.PathName, 0o744)
		if err != nil {
			panic(err)
		}
	}

	log.Println()
	if opt.Speed > 0 {
		scsi_commands.SetCDSpeed(opt)
		log.Println("Set Read Speed to:" + strconv.FormatInt(int64(opt.Speed), 10) + " kbs")
		log.Println()
	}

	sectionMap := sections.GetSectionMap(opt)
	sections.ReadSections(opt, &sectionMap)

	trackMetas, toc, qtoc := split(opt, &sectionMap)
	info(opt, trackMetas, toc, qtoc)
}
