package main

// import "fmt"

func StrLen(s string) int {
	// length := len(s)
	// return fmt.Sprintf("%d", length)
	count := 0
	for range s {
		count++
	}
	return count
}

// func main() {
// 	s := "Hello"
// 	fmt.Println(StrLen(s))
// }

// func StrLen(s string) int {
// 	return len(s)
// }

// func main() {
// 	s := "mayonnaise"
// 	fmt.Println(StrLen(s))
// }
