package main

import "fmt"

/*

10001st prime
Problem 7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

*/

func main() {
	x := 2
	count := 0
	for {
		if isItPrime(x) == true {
			count++

			if count == 10001 {

				break
			}

		}
		x++

	}
	fmt.Println(x)
}

func isItPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
