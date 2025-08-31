package cbd

type Inquiry struct {
	OperationCode  uint8
	Reserved       uint8
	PageCode       uint8
	TransferLength uint16
	Control        uint8
}
