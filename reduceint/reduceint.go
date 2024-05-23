package main

import "github.com/01-edu/z01"

// func ReduceInt(a []int, f func(int, int) int) {
// 	if len(a) == 0 {
// 		return
// 	}
// 	result := a[0]
// 	for i := 1; i < len(a); i++ {
// 		result = f(result, a[i])
// 	}
// 	fmt.Println(result)
// }

func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}

	q := a[0]
	for _, char := range a[1:] {
		q = f(q, char)
	}
	r := Itoa(int64(q))
	for _, char := range r {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

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

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
