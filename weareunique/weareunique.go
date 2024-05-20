package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	charCount := make(map[rune]int)
	for _, char := range str1 {
		charCount[char]++
	}
	for _, char := range str2 {
		charCount[char]++
	}
	uniqueCount := 0
	for _, v := range charCount {
		if v == 1 {
			uniqueCount++
		}
	}
	return uniqueCount
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo")) // Output: 2
	fmt.Println(WeAreUnique("", ""))       // Output: -1
	fmt.Println(WeAreUnique("abc", "def")) // Output: 6
}
