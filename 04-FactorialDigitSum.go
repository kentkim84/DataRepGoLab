// Module: Data Representation
// Tutor: Ian McLoughlin
// Author: Yongjin Kim
// Title: Factorial digit sum
// Source: https://play.golang.org/p/YfpxfVa5pw
// Objective: Create a programme that calculates the sum of the digits of the factorial of a number.
// 1. the factorial of a number, means n! = n(n-1)!
// 2. the sum of the digits from the factorial, example: 12345 => 1+2+3+4+5 = 15

package main

import (
	"fmt"
	"math/big"
)

const INITIAL_NUMBER = 100
const INITIAL_MODULUS = 10
const INITIAL_DIVISION = 10

func calFactorial(x *big.Int) (result *big.Int) {
	result = new(big.Int)
	tempX := new(big.Int)
	tempX.Set(x)

	// func (x *Int) Cmp(y *Int) (r int)
	// -1 if x <  y
	// 0 if x == y
	// +1 if x >  y
	switch tempX.Cmp(&big.Int{}) {
		case -1, 0:
			result.SetInt64(1)
		default:
			result.Set(tempX)
			var one big.Int
			one.SetInt64(1)
			// same as: x * calFactorial(x-1)
			result.Mul(result, calFactorial(tempX.Sub(tempX, &one)))
	}
	return
}

func calDigitSum(x *big.Int) (sum *big.Int) {
	sum = new(big.Int)
	tempX := new(big.Int)
	tempX.Set(x)
	
	for (tempX.Cmp(&big.Int{}) != 0) {
		// 
		remainder := new(big.Int)
		modTen := new(big.Int)
		divTen := new(big.Int)

		// convert initial numbers to big integer
		modTen.SetInt64(INITIAL_MODULUS)
		divTen.SetInt64(INITIAL_DIVISION)

		// get a remainder of x (remainder = x % 10)
		remainder.Mod(tempX, modTen)
		// add up a remainder (sum = sum + remainder)
		sum.Add(sum, remainder)
		// update next x value (x = x / 10)
		tempX.Div(tempX, divTen)
	}
	return
}

func main() {
	x := big.NewInt(INITIAL_NUMBER)
	facto := new(big.Int)
	digitSum := new(big.Int)

	facto = calFactorial(x)
	digitSum = calDigitSum(facto)

	fmt.Println("Factorial is:", facto)
	fmt.Println("Digit Sum is:", digitSum)
}