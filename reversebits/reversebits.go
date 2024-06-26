package main

import "fmt"

func ReverseBits(oct byte) byte {
	var result byte
	for i := 0; i < 8; i++ {
		result = (result << 1) | (oct & 1)
		oct >>= 1
	}
	return result
}

func main() {
	var b byte = 0b00100110
	//fmt.Printf("Original byte: %08b\n", b)
	fmt.Printf("Reversed byte: %08b\n", ReverseBits(b))
}
