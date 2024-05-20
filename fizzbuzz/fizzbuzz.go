package main

import "fmt"

func FizzBuzz(n int) string {
	if n < 0 {
		return "error number is negative\n"
	}
	switch {
	case n%2 == 0 && n%3 == 0:
		return "fizzbuzz"
	case n%2 == 0:
		return "fizz"
	case n%3 == 0:
		return "buzz"
	default:
		return "error: non divisible"
	}
}

func main() {
	// n := 12
	fmt.Println(FizzBuzz(6))
}