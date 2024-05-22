package main

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Wdmatch(s1, s2 string) {
	index := 0

	for _, char := range s2 {
		if index < len(s1) && char == rune(s1[index]) {
			index++
		}
	}

	if index == len(s1) {
		for _, char := range s1 {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	Wdmatch(os.Args[1], os.Args[2])
}
