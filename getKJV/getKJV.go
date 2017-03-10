package main

import (
	"net/http"
	"fmt"

	"time"
	"bufio"
)

func main() {
	var count int = 0
	kjvText, err := http.Get("http://www.gutenberg.org/cache/epub/30/pg30.txt")
	if err != nil {
		fmt.Println("Could not get kjv")

	}
	xStart := time.Now()
	scanKJV := bufio.NewScanner(kjvText.Body)
	defer kjvText.Body.Close()

	scanKJV.Split(bufio.ScanWords)

	for scanKJV.Scan() {

		if scanKJV.Text() == "God" {
			fmt.Println(scanKJV.Text(), )
		}
	}


	zTime := time.Since(xStart)

	fmt.Println(zTime)
}
