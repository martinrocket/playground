package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	dataFeed := strings.NewReader("drone data comming at you....")

	myBuffer := make([]byte, 4)
	counter := 1
	for {

		n, err := dataFeed.Read(myBuffer)
		if err != nil {
			fmt.Printf("error: %v\n\n", err)
			if err == io.EOF {
				fmt.Println(counter, " end of file reached:\n", string(myBuffer[:n]))
				counter++
			}
			break
		}
		fmt.Println(counter, " ...", string(myBuffer[:n]))
		counter++
	}
}
