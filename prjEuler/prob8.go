package main

import (
	"fmt"
	"os"
	//"io/ioutil"
	"bufio"
	"strconv"
)

/*

Largest product in a series
Problem 8
The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.

73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450

Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?

*/

const length int = 13 // This constant allows us to adjust the adjacent digits to read - 4 or 13

func main() {

	// I created a file called data.dat to store the 1000 digit number.  First we open the file.
	file, err := os.Open("data.dat")
	if err != nil {
		fmt.Println("error opening file:", err)
	}

	defer file.Close() // Close the file when this function is finished

	// Scanner is used to ScanLines so that the "/n" is not imported
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var myData string
	for scanner.Scan() {
		myData = myData + scanner.Text()
	}

	// Convert the string into myDataInt slice of []int64 through this for statement
	myDataInt := make([]int64, len(myData))
	for i := range myData {
		myDataInt[i], err = strconv.ParseInt(string(myData[i]), 10, 64)
		if err != nil {
			fmt.Println(err)
		}
	}

	var count int64 = 0                          //use count variable to track the highest product.
	var calc int64 = 1                           // use calc while calculating the product (this is multiplication so start with 1.
	var calcArray [length]int64                  // as we calculate the array with the highest product hold the current array here.
	var countArray [length]int64                 // as we calcualate the array with the highest product hold the highest here.
	for i := 0; i < len(myDataInt)-length; i++ { // the for loop to build arrays to compare
		for j := 0; j < length; j++ {
			calc = calc * myDataInt[i+j] // calculate the array
			calcArray[j] = myDataInt[i+j]
		}

		if calc > 0 { // if the calculated array is > 0 (there we no 0's in the multiplication, do this...

			var m int64 = 1 // calculate the array
			for i := range calcArray {
				m = m * calcArray[i]

			}

		}
		if calc > count { // keep track of the highest product and array
			count = calc
			countArray = calcArray
		}
		calc = 1
	}
	fmt.Println(count)      // print the highest product
	fmt.Println(countArray) // print the coresponding array

}
