package main

func ReverseWords(str string) string {
	rev := ""
	word := ""

	for _, i := range str {
		if i == ' ' {
			rev = rev + word + " " // Adds word and space to result
			word = ""              // Empties word variable
		} else {
			word = string(i) + word // Adds letter to temporary word variable
		}
	}

	return rev + word // reverse those words
	//return rev
}

// func ReverseWords(str string) string {
// 	words := strings.Fields(str) // Split the input string into words
// 	for i := 0; i < len(words)/2; i++ {
// 		// Swap words in the slice to reverse them
// 		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
// 	}
// 	return strings.Join(words, " ") // Join the words back into a string with spaces
// }

// func main() {
// 	fmt.Println(ReverseWords("I am having a good day"))
// }
