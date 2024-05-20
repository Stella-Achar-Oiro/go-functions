package main

import (
	"fmt"
	"os"
)

func Wdmatch(s1, s2 string) {
	if len(os.Args) != 3 {
		return
	}

	i, j := os.Args[1], os.Args[2]
	index := 0

	for _, char := range j {
		if index < len(i) && string(char) == string(i[index]) {
			index++
		}
	}

	if index == len(i) {
		fmt.Println(i)
	}
}

func main() {
	if len(os.Args) == 3 {
		Wdmatch(os.Args[1], os.Args[2])
	}
}
