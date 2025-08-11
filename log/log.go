package log

import (
	"fmt"
	"log"
)

var clean_len = 0

func cleanLine() {
	log.Print("\r")
	for i := 0; i < clean_len+1; i++ {
		log.Print(" ")
	}
	log.Print("\r")
}

func writeLine(a ...any) {
	text := fmt.Sprint(a...)
	log.Print(text)
	clean_len = len(text)
}

func writeLineF(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	log.Print(text)
	clean_len = len(text)
}

func WriteCleanLine(a ...any) {
	cleanLine()
	writeLine(a...)
}

func WriteLineNew(a ...any) {
	cleanLine()
	log.Println(a...)
	clean_len = 0
}
