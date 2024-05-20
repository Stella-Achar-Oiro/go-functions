package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	str, oldLetter, newLetter := os.Args[1], os.Args[2], os.Args[3]

	if !strings.Contains(str, oldLetter) {
		fmt.Println(str)
		return
	}

	result := strings.ReplaceAll(str, oldLetter, newLetter)
	fmt.Println(result)
}
