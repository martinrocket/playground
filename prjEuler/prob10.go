package main

import "fmt"

/*

Summation of primes
Problem 10
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

*/

func main() {

	count := 2

	for i := 3; i < 2000000; i += 2 {
		if isItPr(i) == true {
			count += i
		}
	}

	fmt.Println(count)

}

func isItPr(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
