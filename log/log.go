package log

import (
	"fmt"
	"os"
	"strings"

	"dreamdump/option"
)

var (
	clean_len = 0
	logFile   *os.File
)

func cleanLine() {
	cleanText := "\r" + strings.Repeat(" ", clean_len) + "\r"
	fmt.Print(cleanText)
}

func Print(a ...any) {
	text := fmt.Sprint(a...)
	fmt.Print(text)
	clean_len = len(text)
	if logFile != nil {
		_, err := logFile.WriteString(text)
		if err != nil {
			panic(err)
		}
	}
}

func PrintClean(a ...any) {
	cleanLine()
	text := fmt.Sprint(a...)
	fmt.Print(text)
}

func Println(a ...any) {
	text := fmt.Sprint(a...)
	text += "\n"
	cleanLine()
	fmt.Print(text)
	clean_len = 0
	if logFile != nil {
		_, err := logFile.WriteString(text)
		if err != nil {
			panic(err)
		}
	}
}

func Printf(msg string, a ...any) {
	text := fmt.Sprintf(msg, a...)
	cleanLine()
	clean_len = len(text)
	fmt.Print(text)
	if logFile != nil {
		_, err := logFile.WriteString(text)
		if err != nil {
			panic(err)
		}
	}
}

func Setup(opt *option.Option) {
	logFileName := opt.PathName + "/" + opt.ImageName + ".log"
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	logFile = file
}
