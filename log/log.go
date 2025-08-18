package log

import (
	"fmt"
)

var clean_len = 0

func cleanLine() {
	fmt.Print("\r")
	for i := 0; i < clean_len+1; i++ {
		fmt.Print(" ")
	}
	fmt.Print("\r")
}

func WriteLine(a ...any) {
	text := fmt.Sprint(a...)
	fmt.Print(text)
	clean_len = len(text)
}

func WriteCleanLine(a ...any) {
	cleanLine()
	WriteLine(a...)
}

func WriteLn(a ...any) {
	cleanLine()
	fmt.Println(a...)
	clean_len = 0
}
