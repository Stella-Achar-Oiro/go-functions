package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		for _, char := range os.Args[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}

}

// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	args := os.Args[1:]
// 	for _, arg := range args {
// 		for _, q := range arg {
// 			z01.PrintRune(q)
// 		}
// 		z01.PrintRune('\n')
// 	}

// }
