package driver

const (
	OK    = 0
	ERROR = 1
)

type Status struct {
	Status byte
	Key    byte
	Asc    byte
	AscQ   byte
	Block  []byte
}
