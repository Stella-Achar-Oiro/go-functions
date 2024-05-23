// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Rot13(s string) string {
// 	return strings.Map(rot13, s)
// }

// func rot13(r rune) rune {
// 	switch {
// 	case 'a' <= r && r <= 'z':
// 		return 'a' + (r-'a'+13)
// 	case 'A' <= r && r <= 'Z':
// 		return 'A' + (r-'A'+13)
// 	default:
// 		return r
// 	}
// }

// func main() {
// 	fmt.Println(Rot13("abc"))
// }

package main

import "fmt"

func main() {
	fmt.Println(Rot13("abc"))
}

func Rot13(s string) string {
	q := []rune(s)
	for i, char := range s {
		if char >= 'a' && char <= 'z' {
			q[i] = 'a' + (char - 'a' + 13)%26
		} else if char >= 'A' && char <= 'Z' {
			q[i] = 'A' + (char - 'A' + 13)%26
		}
	}
	return string(q)
}
