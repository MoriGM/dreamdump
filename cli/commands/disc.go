package commands

import (
	"strconv"

	"dreamdump/cd/sections"
	"dreamdump/drive/setup"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/scsi/scsi_commands"
)

func DreamDumpDisc(opt *option.Option) {
	setup.InitializeDrive(opt)
	log.Println()
	if opt.Speed > 0 {
		scsi_commands.SetCDSpeed(opt)
		log.Println("Set Read Speed to:" + strconv.FormatInt(int64(opt.Speed), 10) + " kbs")
		log.Println()
	}

	sectionMap := sections.GetSectionMap(opt)
	sections.ReadSections(opt, sectionMap)

	dense, qtoc := sections.ExtractSections(opt, sectionMap)

	trackMetas, toc, qtoc := split(opt, dense, qtoc)
	info(opt, trackMetas, toc, qtoc)
}
