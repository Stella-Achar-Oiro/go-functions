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

	for _, char := range args {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) > 1 {

// 		fmt.Println(os.Args[1])
// 	}
// }
