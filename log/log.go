package log

import (
	"fmt"

	"dreamdump/drive"
)

var clean_len = 0

func cleanLine() {
	fmt.Print("\r")
	for i := 0; i < clean_len+1; i++ {
		fmt.Print(" ")
	}
	fmt.Print("\r")
}

func Print(a ...any) {
	text := fmt.Sprint(a...)
	fmt.Print(text)
	clean_len = len(text)
}

func PrintClean(a ...any) {
	cleanLine()
	Print(a...)
}

func Println(a ...any) {
	cleanLine()
	fmt.Println(a...)
	clean_len = 0
}

func Printf(msg string, a ...any) {
	cleanLine()
	clean_len, _ = fmt.Printf(msg, a...)
}

func PrintDriveInfo(drive *drive.Drive) {
	Printf("Drive: %s %s %s\n", drive.VendorName, drive.ProductInquiryData, drive.RevisionNumber)
}
