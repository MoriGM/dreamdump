package writers

import (
	"dreamdump/cd/sections"
	"dreamdump/option"
)

func WriteSection(opt *option.Option, section *sections.Section) {
	WriteBinSection(opt, section)
	WriteSubSection(opt, section)
}
