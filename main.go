package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello Gophers!")
	myTime := time.Now()
	fmt.Println(myTime.Format("Jan _2 15:04:05.000"))
	fmt.Println(myTime.Format("01_02_15:04:05.000"))
	fmt.Println(myTime.Format("01_02_2006"))
	fmt.Print("Enter text: ")
	//var text []byte
	//fmt.Scanln(&text)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		break
	}

	myText := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(myText)

}
