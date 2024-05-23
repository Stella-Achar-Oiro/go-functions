package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1]
	var result string
	for _, ch := range args {
		if ch >= 'a' && ch <= 'z' {
			result += string('z' - (ch - 'a'))
		} else if ch >= 'A' && ch <= 'Z' {
			result += string('Z' - (ch - 'A'))
		} else {
			result += string(ch)
		}
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
