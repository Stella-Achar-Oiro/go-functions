package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := len(os.Args[1:])
	z01.PrintRune(rune(count + '0'))
}

