// Module: Data Representation
// Tutor: Ian McLoughlin
// Author: Yongjin Kim
// Title: Guessing game
// Objective: Create a programme that the user must guess a randomly generated number.
// 1. After every guess the program tells the user whether their number was too high or too low.
// 2. At the end, the number of tries should be printed.
// 3. It counts only as one try if they input the same number multiple times consecutively.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generate a random number
func genRandInt(minNum int, maxNum int) int {
    return rand.Intn(maxNum) + minNum
}

func main() {

	var maxNum int
	var guessNum int
	var randNum int
	var tryNum int

	// initial setup
	tryNum = 1
	rand.Seed(time.Now().UnixNano())
	// setup the range of a random number starts from 1 to maximum number
	fmt.Print("Enter a maximum number to generate a random number, range of 1 to n: ")
	fmt.Scanln(&maxNum)
	// create slices(=dynamic array)
	numContainer := make([]int, 0)
	// generate a random number
	randNum = genRandInt(1, maxNum)
	fmt.Println("Random number generated", randNum) // debuging purpose
	fmt.Println("Contents of the container", numContainer)
	// prompt user's guessing number
	fmt.Print("Enter a guessing number: ")
	fmt.Scanln(&guessNum)
	// store the guessing number in the number container
	numContainer = append(numContainer, guessNum)
	fmt.Println("Contents of the container", numContainer) // debuging purpose


	for (guessNum != randNum) {
		if (guessNum < randNum) {
			fmt.Println("Your number is lower than the random number")
		} else if (guessNum > randNum) {
			fmt.Println("Your number is higher than the random number")
		}
		tryNum++
		fmt.Println("Number of tries:", tryNum)
		fmt.Print("Enter a guessing number: ")
		fmt.Scanln(&guessNum)
		
		find := true
		for find {
			for i := 0; i < len(numContainer); i++  {
				if (guessNum == numContainer[i]) {
					fmt.Print("Enter a different guessing number: ")
					fmt.Scanln(&guessNum)
				} else {
					find = false
				}
				fmt.Println("A number in the container", numContainer[i]) // debuging purpose
			}
		}
			
		numContainer = append(numContainer, guessNum)
		fmt.Println("Contents of the container", numContainer) // debuging purpose
	}

	if (guessNum == randNum) {
		fmt.Println("You have tried:", tryNum)
		fmt.Println("Numbers matched")
	}

}