// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}

// 	n, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		return
// 	}

// 	isPowerOfTwo := (n > 0) && (n&(n-1) == 0)

// 	fmt.Println(isPowerOfTwo)
// }

package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n := BasicAtoi(os.Args[1])

	isPowerOfTwo := (n > 0) && (n&(n-1) == 0)

	if isPowerOfTwo {
		os.Stdout.WriteString("true\n")
	} else {
		os.Stdout.WriteString("false\n")
	}
}

func BasicAtoi(s string) int {
	var result int
	for _, c := range s {
		result = result*10 + int(c-'0')
	}
	return result
}
