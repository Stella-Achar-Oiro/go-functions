package main

import (
	"fmt"
)

func StrRev(s string) string {
	reversed := make([]rune, len(s))
	for i, char := range s {
		reversed[len(s)-1-i] = char
	}
	return string(reversed)
}

func main() {
	s := "mob"
	fmt.Println(StrRev(s))
}
