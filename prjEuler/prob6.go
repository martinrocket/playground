package main

import (
	"fmt"
)

/*
Sum square difference
Problem 6
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.1``1`
*/

func main() {
	x := 0
	y := 0
	for i := 1; i < 101; i++ {
		x += i * i
	}

	for i := 1; i < 101; i++ {
		y = y + i

	}

	fmt.Println(y)
	y = y * y
	fmt.Println(x)
	fmt.Println(y)
	z := y - x
	fmt.Println(z)

}
