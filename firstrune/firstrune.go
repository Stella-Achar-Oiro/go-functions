package main

import "github.com/01-edu/z01"

func FirstRune(s string, n int) rune {
	if n <= 0 || n >= len(s) {
		return 0
	}
	return []rune(s)[n-1]
}



func main() {
	z01.PrintRune(FirstRune("Hello!", 2))
	z01.PrintRune(FirstRune("Salut!", 4))
	z01.PrintRune(FirstRune("Ola!", 3))
	z01.PrintRune('\n')
}
