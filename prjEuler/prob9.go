package main

import (
	"fmt"
	"math"
)

/*
	Special Pythagorean triplet
	Problem 9
	A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

	a2 + b2 = c2
	For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

	There exists exactly one Pythagorean triplet for which a + b + c = 1000.
	Find the product abc.
*/
func main() {

	var i float64
	var j float64
	var k float64

	for i = 1; i <= 1000; i++ {
		for j = 1; j <= 1000; j++ {
			for k = 1; k <= 1000; k++ {
				if i+j+k == 1000 {
					if i < j && j < k {
						if math.Pow(i, 2)+math.Pow(j, 2) == math.Pow(k, 2) {
							fmt.Printf(" The Product of %v %v %v is %f.\n", i, j, k, i*j*k)

						}
					}

				}
			}

		}
	}

}
