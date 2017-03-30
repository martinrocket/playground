package main

import "fmt"
import "os"
import "bufio"

func main() {

	myFile, err := os.Open("data.dat")
	if err != nil {
		fmt.Println("cannot open file")
	}
	defer myFile.Close()

	var myData string
	myScanner := bufio.NewScanner(myFile)
	myScanner.Split(ScanLines)
	for myScanner.Scan() {
		myData = myScanner.Text()
	}

	fmt.Println(myData)
}
