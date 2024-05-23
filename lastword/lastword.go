package main

import (
	"os"

	//"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	lastWord := ""
	wordStart := false
	for i := len(args) - 1; i >= 0; i-- {
		if args[i] != ' ' {
			wordStart = true
			lastWord = string(args[i]) + lastWord
		} else if wordStart {
			break
		}
		//z01.PrintRune(lastWord)
	}
	//z01.PrintRune('\n')
	os.Stdout.WriteString(lastWord + "\n")
}
