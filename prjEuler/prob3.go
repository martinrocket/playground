package main

import (
	"fmt"
	"math"
)

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

var n int = 600851475143

//var n int = 13195

func main() {
	for n%2 == 0 {
		fmt.Printf("%d ", 2)
		n = n / 2
	}

	x := int(math.Sqrt(float64(n)))

	for i := 3; i <= x; i = i + 2 {
		for n%i == 0 {
			fmt.Printf("%d ", i)
			n = n / i
		}
	}

	if n > 2 {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

}
