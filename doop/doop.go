package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return // Invalid number of arguments
	}

	value1 := os.Args[1]
	value2 := os.Args[3]
	operator := os.Args[2]

	var result string

	result = performOperation(value1, operator, value2)

	if result != "" {
		os.Stdout.WriteString(result + "\n")
	}
}

// Atoi function (converts string to int64)
func Atoi(s string) (bool, int64) {
	var num int64
	var sign int64 = 1
	for i, c := range s {
		if c == '-' && i == 0 {
			sign = -1
		} else if c >= '0' && c <= '9' {
			num = num*10 + int64(c-'0')
			if num < 0 {
				return false, 0 // Overflow occurred
			}
		} else {
			return false, 0
		}
	}
	return true, sign * num
}

// Itoa function (converts int64 to string)
func Itoa(num int64) string {
	if num == 0 {
		return "0"
	}
	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}
	var result string
	for num > 0 {
		digit := num % 10
		result = string(rune('0'+digit)) + result
		num /= 10
	}
	return sign + result
}

// performOperation performs arithmetic based on the operator and handles overflow.
func performOperation(a, operator, b string) string {
	validA, numA := Atoi(a)
	validB, numB := Atoi(b)
	if !validA || !validB {
		return ""
	}

	var result int64

	switch operator {
	case "+":
		// Check for overflow
		if (numB > 0 && numA > 9223372036854775807-numB) || (numB < 0 && numA < -9223372036854775808-numB) {
			return ""
		}
		result = numA + numB
	case "-":
		// Check for overflow
		if (numB < 0 && numA > 9223372036854775807+numB) || (numB > 0 && numA < -9223372036854775808+numB) {
			return ""
		}
		result = numA - numB
	case "*":
		// Check for overflow
		if numB != 0 && ((numA > 9223372036854775807/numB && numB > 0) || (numA < -9223372036854775808/numB && numB < 0)) {
			return ""
		}

		result = numA * numB
	case "/":
		if numB == 0 {
			return "No division by 0"
		}
		result = numA / numB
	case "%":
		if numB == 0 {
			return "No modulo by 0"
		}
		result = numA % numB
	default:
		return "" // Invalid operator
	}

	return Itoa(result)
}


