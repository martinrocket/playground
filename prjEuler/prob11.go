package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

	Problem 11


	In the 20×20 grid below, four numbers along a diagonal line have been marked in red.

	08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
	49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
	81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
	52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
	22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
	24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
	32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
	67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
	24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
	21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
	78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
	16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
	86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
	19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
	04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
	88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
	04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
	20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
	20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
	01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48

	The product of these numbers is 26 × 63 × 78 × 14 = 1788696.

	What is the greatest product of four adjacent numbers in the same direction (up, down, left, right, or diagonally) in the 20×20 grid?

*/

const adjacent int = 4

func main() {

	grid := createGrid()
	grW := len(grid)
	grH := len(grid[0])
	fmt.Printf("Grid is type %T with width %d and grid length is %d.\n", grid, grW, grH)
	myRArray, myRCount := right(grid)
	myLArray, myLCount := left(grid)
	myDArray, myDCount := down(grid)
	myARArray, myARCount := acrossR(grid)
	myALArray, myALCount := acrossL(grid)

	fmt.Println(myRArray, myRCount)
	fmt.Println(myLArray, myLCount)
	fmt.Println(myDArray, myDCount)
	fmt.Println(myARArray, myARCount)
	fmt.Println(myALArray, myALCount)

	largest := 0
	if myRCount > largest {
		largest = myRCount
	}

	if myLCount > largest {
		largest = myLCount
	}

	if myDCount > largest {
		largest = myDCount
	}

	if myARCount > largest {
		largest = myARCount
	}
	if myALCount > largest {
		largest = myALCount
	}

	fmt.Printf("\nThe largest product is: %d\n", largest)
}

func createGrid() [][]int {
	// the data for our grid was copied to a file called data_11.dat.
	// Open the file
	file, err := os.Open("data_11.dat")

	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}

	defer file.Close() //Close the file when we've read the data

	scanner := bufio.NewScanner(file) // a new scanner to read the data from the file
	scanner.Split(bufio.ScanLines)    // set the scanner to read a line, we'll seperate it later

	lines := make([]string, 0, 40) // make a slice of 0 length, we'll add lines in the for loop below

	i := 0
	for scanner.Scan() {
		lines = lines[:i+1] // add a line to our []slice
		lines[i] = scanner.Text()

		i++

	}

	// Build a 2D Slice (grid!)
	// Make a grid of integers for our calculations. The grid is a two dimensional
	// slice. This is built by first build a 2D slice, setting the number of rows first
	// with the len(lines) variable. Then for every row, you build the number of
	// columns using the len(elements) var (where I find the number of elements in
	// each row.

	// first the 2d Slice called Grid with row number len(lines)
	grid := make([][]int, len(lines))

	// second the number of columns
	elements := strings.Split(lines[0], " ")

	// third the number of columns in each row
	for i := range lines {
		grid[i] = make([]int, len(elements))
	}

	// now fill the slice grid with values
	for j := range lines {
		lineStr := strings.Split(lines[j], " ")
		for k := range lineStr {
			grid[j][k], err = strconv.Atoi(lineStr[k])
			if err != nil {
				fmt.Printf("Error in building grid at %v : %v", grid[j][k], err)
			}
		}

	}

	return grid
}

func right(grid [][]int) ([4]int, int) {
	count := 0
	var countArray [4]int
	calc := 1
	var calcArray [4]int
	for i := range grid {
		for j := 0; j < len(grid[i])-4; j++ {
			calc = 1
			for r := range calcArray {
				calcArray[r] = 0
			}
			for k := 0; k < 4; k++ {
				calc = calc * grid[i][j+k]
				calcArray[k] = grid[i][j+k]
				if calc > count {
					count = calc
					countArray = calcArray
				}
			}
		}

	}
	return countArray, count
}

func left(grid [][]int) ([4]int, int) {
	count := 0
	var countArray [4]int
	calc := 1
	var calcArray [4]int
	for i := range grid {
		for j := len(grid[i]) - 1; j > 4; j-- {
			calc = 1
			for r := range calcArray {
				calcArray[r] = 0
			}
			for k := 0; k < 4; k++ {
				calc = calc * grid[i][j-k]
				calcArray[k] = grid[i][j-k]
				if calc > count {
					count = calc
					countArray = calcArray
				}
			}
		}

	}
	return countArray, count
}

func down(grid [][]int) ([4]int, int) {
	count := 0
	var countArray [4]int
	calc := 1
	var calcArray [4]int
	for i := 0; i < len(grid)-4; i++ {
		for j := 0; j < len(grid[i]); j++ {
			calc = 1
			for r := range calcArray {
				calcArray[r] = 0
			}
			for k := 0; k < 4; k++ {
				calc = calc * grid[i+k][j]
				calcArray[k] = grid[i+k][j]
				if calc > count {
					count = calc
					countArray = calcArray
				}
			}
		}

	}
	return countArray, count
}

func acrossR(grid [][]int) ([4]int, int) {
	count := 0
	var countArray [4]int
	calc := 1
	var calcArray [4]int
	for i := 0; i < len(grid)-4; i++ {
		for j := 0; j < len(grid[i])-4; j++ {
			calc = 1
			for r := range calcArray {
				calcArray[r] = 0
			}
			for k := 0; k < 4; k++ {
				calc = calc * grid[i+k][j+k]
				calcArray[k] = grid[i+k][j+k]
				if calc > count {
					count = calc
					countArray = calcArray
				}
			}
		}

	}
	return countArray, count
}

func acrossL(grid [][]int) ([4]int, int) {
	count := 0
	var countArray [4]int
	calc := 1
	var calcArray [4]int
	for i := 0; i < len(grid)-4; i++ { // i := o - i = len(grid) = 4 or 0 - i++
		for j := 4; j < len(grid[i]); j++ { // j := 4 - len(grid[i] = 4 or 0 - j++ or j--
			calc = 1
			for r := range calcArray {
				calcArray[r] = 0
			}
			for k := 0; k < 4; k++ { // k:=0 - k<4 - k++
				calc = calc * grid[i+k][j-k] // i+k j-k - i+k j+k - i+k j - i j-k -
				calcArray[k] = grid[i+k][j-k]
				if calc > count {
					count = calc
					countArray = calcArray
				}
			}
		}

	}
	return countArray, count
}

func findHighest(grid [][]int, iVal int, jVal int, kval int, iGrid int, jGrid int) ([4]int, int) {
	// local variables for counting
	var count, calc int
	var countArray, calcArray [4]int
	for i := iVal; i < len(grid)-iGrid; i++ {
		for j := jVal; j < leng(grid[i]-jGrid); j++ {
				count, countArray = getCalcAL()
			
		}

}

func getCalcAL() int int {  // get calc and calcArray for Across and to Left
	for k := 0; k < 4; k++ { 
			calc := 1
			calcArray := [4]int(0,0,0,0)
			calc = calc * grid[i+k][j-k] // i+k j-k - i+k j+k - i+k j - i j-k -
			calcArray[k] = grid[i+k][j-k]
			if calc > count {
				count = calc
				countArray = calcArray
				
			}

	}
	return count, countArray
}
