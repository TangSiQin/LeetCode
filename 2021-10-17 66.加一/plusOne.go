package main

import "fmt"

func main() {
	digits := []int{1,2,4} // 1,9,9
	fmt.Print(plusOne(digits))
}

func plusOne(digits []int) []int{
	l := len(digits)

	for i := l-1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		}

		digits[i] = 0
	}


	digits = make([]int, l+1)
	digits[0] = 1


	return digits
}
