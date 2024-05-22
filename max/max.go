package main

import "fmt"

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	max := a[0]
	for _, q := range a {
		if q > max {
			max = q
		}
	}
	return max
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
