package log

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"dreamdump/option"
)

var (
	cleanLen = 0
	logFile  *os.File
)

func cleanLine() {
	cleanText := "\r" + strings.Repeat(" ", cleanLen) + "\r"
	fmt.Print(cleanText)
}

func Print(a ...any) {
	text := fmt.Sprint(a...)
	fmt.Print(text)
	cleanLen = len(text)
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
	cleanLen = 0
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
	cleanLen = len(text)
	fmt.Print(text)
	if logFile != nil {
		_, err := logFile.WriteString(text)
		if err != nil {
			panic(err)
		}
	}
}

func Setup(opt *option.Option) {
	_, err := os.Stat(opt.PathName)
	if errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(opt.PathName, 0o744)
		if err != nil {
			panic(err)
		}
	}
	logFileName := opt.PathName + "/" + opt.ImageName + ".log"
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	logFile = file
}
