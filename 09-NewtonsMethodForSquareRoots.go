// Module: Data Representation
// Tutor: Ian McLoughlin
// Author: Yongjin Kim
// Title: Newton’s method for square roots
// Source: https://gist.github.com/abesto/3476594
// Objective: Write a function to calculate the square root of a number using Newton’s method.

package main

import (
	"fmt"
	"math"
)

// setup digit accuracy (7 decimal point) 
const DELTA = 0.0000001

func Sqrt(x float64, y float64) float64 {
	// setup an initial guess
	z := y
    
    step := func() float64 {
    	return z - (z*z - x) / (2 * z)
    }
    // math.Abs returns absoulte value of (x), repeat to get the closest value
    for nextZ := step(); math.Abs(nextZ - z) > DELTA
    {
    	z = nextZ
		nextZ = step()
    }
    return z
}

func main() {
	var userNumber float64
	var userInitGuess float64
	
	fmt.Print("Enter a number to be processed: ")
	fmt.Scanf("%f\n", &userNumber)
	fmt.Print("Enter an initial guessing number: ")
	fmt.Scanf("%f\n", &userInitGuess)

	fmt.Println(Sqrt(userNumber, userInitGuess))
	fmt.Println(math.Sqrt(userNumber))
}