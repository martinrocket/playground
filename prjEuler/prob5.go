package main

import (
	"fmt"
)

/*

Smallest multiple
Problem 5
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

*/

var x int
var y int = 0

func main() {
	myBool := false
	for {
		x = getNext()
		myBool = isItIt(x)
		if myBool == true {
			fmt.Println("Your number is ", x)
			break
		}
	}

}

func getNext() int {
	y += 20
	return y

}

func isItIt(x int) bool {
	for i := 1; i < 20; i++ {
		if x%i != 0 {
			return false
		}
	}
	return true

}
