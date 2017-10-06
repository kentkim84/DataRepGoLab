GMIT Software Development 3rd year
Module: Data Representation
GoLang Problem sheet

### Go Examples

This repository contains some example code written in the programming language Go. The author is Yongjin Kim, a Student at GMIT.

### 1. Kon’nichiwa, Sekai!
Program that prints Hello, world! in Japanese (using Japanese characters) to the screen.
> A basic practice in any programming languages, different programme syntax, but worth to try to understand how it's structured.
### 2. Current time
Program that prints the current time and date to the console.
> Use a package called 'time', understand of usage of functions in a package 
### 3. FizzBuzz
Program that prints the numbers from 1 to 100, and follows conditions.
- Multiples of three print 'Fizz' instead of the number.
- Multiples of five print Buzz.
- Numbers that are multiples of both three and five print FizzBuzz.
> Understanding of nested if statements and use of modulus to get multiples of a number (multiples of three can be figured by using **'number' % 3**)
### 4. Factorial digit sum
Program that calculates the sum of the digits of the factorial of a number.
- n! means n x (n − 1) ... x 3 x 2 x 1. For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800
- The sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
- Then find the sum of the digits in the number 100!
### 5. Guessing game
Write a guessing game where the user must guess a randomly generated number. After every guess the program tells the user whether their number was too high or too low. At the end, the number of tries should be printed. It counts only as one try if they input the same number multiple times consecutively.
> Compose of two packages, *math/rand* and *time*, these two will generate a non-repeated number in the program run time, *rand.Seed(time.Now().UnixNano())*.
### 6. Largest and smallest in list
Write a function that returns the largest and smallest elements in a list.
> Comparing two strings which will be used in the for-loop and if statements to control the flow of program.
### 7. Palindrome test
Write a function that tests whether a string is a palindrome. A palindrome is a string that reads the same in reverse, e.g. radar.
> Understanding of reversing algorithm. In this program, user's input text, *string* type, can be reversed and compared to the original string text.
### 8. Merge list and sort
Write a function that merges two sorted lists into a new sorted list, e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].
> One of many sorting algorithms, practiced to understand **Recursion** logic that *sort* function runs till the conditions reached, then *merge* sorted data.
### 9. Newton’s method for square roots
Write a function to calculate the square root of a number using Newton’s method. Newton’s method is to approximate the square root of a number x by picking a starting point z and then repeating the following operation.
- z_next = z - ((z*z - x) / (2 * z))
> Actual implementation of the Newton's square root method, and compare to the function *math.Sqrt* from the math package. [How square root is implemented in GO](https://golang.org/src/math/sqrt.go?s=3702:3730#L82)
### 10. Reverse string
Write a function to reverse a string in Go.
> Same reverse logic like *Palindrome test*


## Coding Standards
// Version 0.1 using C standards

### 1. Naming Conventions and Style
1.1. Use Pascal casing for class and structs
    
    class PlayerManager;
    struct AnimationInfo;

1.2. Use camel casing for local variable names and function parameters
    
    void SomeMethod(const int someParameter);
    {
        int someNumber;
    
    }

1.3. Use verb-object pairs for method names
a.	Use pascal casing for public methods
        
    public:
    void DoSomething();

b.	Use camel casing for other methods
        
    private:
    void doSomething();

### References
This examples are from
* [Data Representation and Querying - Ian McLoughlin](https://data-representation.github.io/problems/go-fundamentals.html)

This coding standard is followed by
* [popeKim's c/c++ coding standards](https://docs.google.com/document/d/1cT8EPgMXe0eopeHvwuFmbHG4TJr5kUmcovkr5irQZmo/edit#heading=h.r2n9mhxbh2gg)