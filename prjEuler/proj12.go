package main

import (
	"fmt"
)

var n1 = 2
var n2 = 2

func main() {
	fmt.Println("Project 12")
	for {
		counter := 0
		for i := 1; i <= n2; i++ {
			if n2%i == 0 {
				counter++
			}
		}

		if counter > 500 {
			fmt.Printf("%v = %v \n", n2, counter)
			break
		}

		if counter%100 == 0 {
			fmt.Printf("*")
		}
		if counter > 500 {
			fmt.Printf("%v = %v \n", n2, counter)
			break
		}

		n1 = n1 + 2
		n2 = n1 + n2
	}

}

// 1. Start by inputting a number n
// 2. Let an int variable limit be sqrt(n)
// 3. Run a loop from i=1 to  i=limit
//     3.1 if n is divisible by i
//           3.1.1 Print i
//           3.1.2 if i and n/i are unequal, print n/i also.
// 4. End.
