package sections

import (
	"slices"
	"strconv"

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
