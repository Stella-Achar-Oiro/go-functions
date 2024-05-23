package main

import "fmt"

// SwapBits swaps the higher 4 bits with the lower 4 bits of the given byte
func SwapBits(octet byte) byte {
	return (octet >> 4) | (octet << 4)
}

func main() {
	octet := byte(0b01000001) // Example byte: 0100 0001
	swapped := SwapBits(octet)
	fmt.Printf("Original byte: %08b\n", octet)  // Output: 0100 0001
	fmt.Printf("Swapped byte: %08b\n", swapped) // Output: 0001 0100
}
