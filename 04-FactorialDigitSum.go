// Module: Data Representation
// Tutor: Ian McLoughlin
// Author: Yongjin Kim
// Title: Factorial digit sum
// Objective: Create a programme that calculates the sum of the digits of the factorial of a number.
// 1. the factorial of a number, means n! = n(n-1)!
// 2. the sum of the digits from the factorial, example: 12345 => 1+2+3+4+5 = 15

package main

import "fmt"

func calFactorial(x float64) float64 {
	if (x == 0) {
		return 1
	}
	return x * calFactorial(x-1)
}

func calSumDigit(x float64) uint {
	var sum, remainder uint

	for (x != 0) {
	   remainder = uint(x) % 10
	   sum = sum + remainder
	   x = x / 10
	}
	return sum
}

func main() {	
	x := float64(100)
	fmt.Println(calFactorial(x))
	fmt.Println(calSumDigit(calFactorial(x)))
}