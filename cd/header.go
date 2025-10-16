package cd

import (
	"strings"

	"dreamdump/log"
)

const (
	HEADER_START                  = 0x10
	HEADER_END                    = 0x100
	HEADER_SIZE                   = HEADER_END - HEADER_START
	HEADER_BUILD_DATE_YEAR_START  = 0x50
	HEADER_BUILD_DATE_YEAR_END    = 0x54
	HEADER_BUILD_DATE_MONTH_START = 0x54
	HEADER_BUILD_DATE_MONTH_END   = 0x56
	HEADER_BUILD_DATE_DAY_START   = 0x56
	HEADER_BUILD_DATE_DAY_END     = 0x58
	HEADER_VERSION_START          = 0x4B
	HEADER_VERSION_END            = 0x50
	HEADER_SERIAL_START           = 0x40
	HEADER_SERIAL_END             = 0x4a
	HEADER_REGION_START           = 0x30
	HEADER_REGION_END             = 0x3a
	HEADER_REGION_JAPAN           = 0x00
	HEADER_REGION_USA             = 0x01
	HEADER_REGION_EUROPE          = 0x02
)

func PrintHeader(sector *CdSectorData) {
	gdToc := sector[TOC_OFFSET:TOC_OFFSET_END]
	if gdToc[0] != 'T' || gdToc[1] != 'O' || gdToc[2] != 'C' || gdToc[3] != '1' {
		panic("Header TOC IN LBA 45000 doesn't have a magic number TOC1")
	}
	log.Println("DC Header:")
	header := sector[HEADER_START:HEADER_END]
	log.Printf("  build date: %s-%s-%s\n", header[HEADER_BUILD_DATE_YEAR_START:HEADER_BUILD_DATE_YEAR_END], header[HEADER_BUILD_DATE_MONTH_START:HEADER_BUILD_DATE_MONTH_END], header[HEADER_BUILD_DATE_DAY_START:HEADER_BUILD_DATE_DAY_END])
	log.Printf("  version: %s\n", header[HEADER_VERSION_START:HEADER_VERSION_END])
	log.Printf("  serial: %s\n", strings.ReplaceAll(string(header[HEADER_SERIAL_START:HEADER_SERIAL_END]), " ", ""))
	log.Printf("  region: %s\n", headerRegion(header[HEADER_REGION_START:HEADER_REGION_END]))
	log.Printf("  header: \n")
	for row := 0; row < HEADER_END; row += 0x10 {
		log.Printf("%04x : % 2x   %s\n", row, header[row:(row+0x10)], header[row:(row+0x10)])
	}
}

func headerRegion(region []byte) string {
	var regionString strings.Builder
	if region[HEADER_REGION_JAPAN] == 'J' {
		regionString.WriteString("Japan ")
	}
	if region[HEADER_REGION_USA] == 'U' {
		regionString.WriteString("USA ")
	}
	if region[HEADER_REGION_EUROPE] == 'E' {
		regionString.WriteString("Europe")
	}
	return regionString.String()
}
