// Module: Data Representation
// Tutor: Ian McLoughlin
// Author: Yongjin Kim
// Title: FizzBuzz
// Objective: Create a programme that prints the numbers from 1 to 100
// 1. multiples of three print "Fizz"
// 2. multiples of five print "Buzz"
// 3. multiples of both three and five print "FizzBuzz"

package main

import "fmt"

func main() {
	// prints the numbers 1 to 100
	for i := 1; i <= 100; i++ {
		if (i%3 == 0 && i%5 == 0) { // multiples of both three and five
			fmt.Printf("FizzBuzz\n")
		} else if (i%3 == 0) { // multiples of three
			fmt.Printf("Fizz\n")
		} else if (i%5 == 0) { // multiples of five
			fmt.Printf("Buzz\n")
		} else {
			fmt.Printf("value of i: %d\n", i)
		}	
	}
}