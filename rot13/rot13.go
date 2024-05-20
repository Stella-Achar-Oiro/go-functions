package main

import (
	"fmt"
	"strings"
)

func Rot13(s string) string {
	return strings.Map(rot13, s)
}

func rot13(r rune) rune {
	switch {
	case 'a' <= r && r <= 'z':
		return 'a' + (r-'a'+13)%26
	case 'A' <= r && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	default:
		return r
	}
}

func main() {
	fmt.Println(Rot13("abc"))
}
