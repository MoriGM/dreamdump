package drive

type Drive struct {
	VendorName         [8]byte
	ProductInquiryData [16]byte
	RevisionNumber     [4]byte
	RevionDate         [10]byte
}
