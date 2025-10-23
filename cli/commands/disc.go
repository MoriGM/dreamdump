package commands

import (
	"os"
	"strconv"

	"dreamdump/cd/sections"
	"dreamdump/drive/setup"
	"dreamdump/exit_codes"
	"dreamdump/log"
	"dreamdump/option"
	"dreamdump/scsi/scsi_commands"
)

func DreamDumpDisc(opt *option.Option) {
	setup.InitializeDrive(opt)
	log.Println()
	if opt.Speed > 4 {
		scsi_commands.SetCDSpeed(opt)
		log.Println("Set Read Speed to:" + strconv.FormatInt(int64(opt.Speed), 10) + " kbs")
		log.Println()
	}

	sectionMap := sections.GetSectionMap(opt)
	matching := sections.ReadSections(opt, sectionMap)

	if !matching {
		log.Println()
		log.Println("Retry count exceeded. Please try again.")
		if !sectionMap[0].Matched {
			log.Println("Section 1 isn't matching. Please try with --train")
		}
		os.Exit(exit_codes.NO_MATCHING_READS)
	}

	sections.CombineSections(opt, sectionMap)

	dense, qtoc := sections.ExtractSections(opt, sectionMap)

	trackMetas, toc, qtoc, headerSector := split(opt, dense, qtoc)
	info(opt, trackMetas, toc, qtoc, headerSector)
}
