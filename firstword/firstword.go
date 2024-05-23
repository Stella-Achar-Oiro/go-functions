package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	for _, v := range arg {
		if v == ' ' {
			break
		}
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
