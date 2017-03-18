package main

import (
	"fmt"

)

/*
Multiples of 3 and 5
Problem 1
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
 */

const x int = 1000

func main() {
	ctr := 0
	for i := 1; i < x; i ++ {
		if i % 3 == 0 || i % 5 == 0 {
			ctr += i
			fmt.Println(i)
		}
	}
	fmt.Println(ctr)
}