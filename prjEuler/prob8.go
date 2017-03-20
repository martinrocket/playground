package main

import (
	"fmt"
	"os"
	"io/ioutil"
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
const length int = 13

func main() {
	fmt.Println("sugar")

	file, err := os.Open("data.dat")
	if err != nil {
		fmt.Println("error opening file")
	}

	defer file.Close()

	myData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error reading file into memory")
	}

	//fmt.Printf("%s %T \n", myData, myData)
	fmt.Println(len(myData))
	myDataInt := make([]int64, len(myData))

	for i := 0; i < len(myData); i++ {
		myDataInt [i], err = strconv.ParseInt(string(myData[i]), 10, 64)
		if err != nil {
			fmt.Println(err)
		}
	}
	//fmt.Println(myDataInt)
	var count int64 = 0
	var calc int64 = 1
	var calcArray [length]int64
	var countArray [length]int64
	for i :=0; i < len(myDataInt)-length; i++ {
		for j := 0; j < length; j++ {
			calc = calc * myDataInt[i + j]
			calcArray[j] = myDataInt[i + j]
		}
		//fmt.Println(calc)
		//fmt.Println(calcArray)
		if calc > 0 {
			fmt.Printf("%v %v ", calc, calcArray)
			var m int64 = 1
			for i := range calcArray {
				m = m * calcArray[i]

			}
			fmt.Printf("%v \n", m)
		}
		if calc > count {
			count = calc
			countArray = calcArray
		}
		calc = 1
	}
	//fmt.Println(count)
	//fmt.Print(countArray)
	var n int64 = 1
	for i := 0; i < len(countArray); i++ {
		n = n * countArray[i]
	}
	//fmt.Println(n)
	fmt.Println()



}
