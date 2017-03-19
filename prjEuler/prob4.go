package main

import (
	"fmt"
	"strconv"
)



/*

Largest palindrome product
Problem 4
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
 */



func main() {

	pal4()
	pal6()

}

func pal4() {

	for i := 9999; i > 999; i-- {
		var a, b string
		a = strconv.Itoa(i)[:2]
		b = strconv.Itoa(i)[3:] + strconv.Itoa(i)[2:3]
		if a == b {
			//fmt.Printf("number %d is palindrome \n", i)
			for j := 2; j < 100; j++ {
				if i%j == 0 && i/j < 99 {
					k := i / j
					fmt.Printf("number %d is palindrom and %d and %d are factors.\n", i, j, k)
					return
				}
			}
		}
	}
}

func pal6() {

	for i := 999999; i > 9999; i-- {
		var a, b string
		a = strconv.Itoa(i)[:3]
		b = strconv.Itoa(i)[5:] + strconv.Itoa(i)[4:5] + strconv.Itoa(i)[3:4]
		if a == b {
			//fmt.Printf("number %d is palindrome \n", i)
			for j := 2; j < 1000; j++ {
				if i%j == 0 && i/j < 999 {
					k := i / j
					fmt.Printf("number %d is palindrom and %d and %d are factors.\n", i, j, k)
					return
				}
			}
		}
	}

}
