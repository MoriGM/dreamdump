package sections

import (
	"slices"
	"strconv"

	"dreamdump/option"

	"github.com/cespare/xxhash/v2"
)

func (sect *Section) Hash() string {
	digest := xxhash.New()
	for _, sector := range sect.Sectors {
		n, err := digest.Write(sector.Data[:])
		if err != nil {
			panic(err)
		}
		if n != len(sector.Data) {
			panic("Incorrect write length for sector")
		}
	}

	return strconv.FormatInt(int64(digest.Sum64()), 36)
}

func (sect *Section) IsMatching(currentHash string) bool {
	return slices.Contains(sect.Hashes, currentHash)
}

func (sect *Section) AddHash(currentHash string) {
	sect.Hashes = append(sect.Hashes, currentHash)
}

func (sect *Section) FileName(opt *option.Option) string {
	sectionRange := strconv.FormatInt(int64(sect.StartSector), 10) + "-" + strconv.FormatInt(int64(sect.EndSector), 10)
	sectionName := opt.PathName + "/" + opt.ImageName + "-" + sectionRange
	return sectionName
}
