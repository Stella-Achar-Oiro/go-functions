package main

import (
	"fmt"
	"strconv"
)

func HexToInt(hex string) string {
	number, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(number)
}

func main() {
	fmt.Println(HexToInt("1E"))
}
