package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Printf("Could not create file, error: %q\n", err)
	}

	defer f.Close()

	r := rand.New(rand.NewSource(99))
	for i := 0; i < 100; i++ {
		aFl := r.Float64()
		aInt := r.Int63n(10)
		bFl := r.Float64()
		bInt := r.Int63n(10)
		aStr := strconv.FormatFloat(aFl+float64(aInt), 'f', -1, 64)
		bStr := strconv.FormatFloat(bFl+float64(bInt), 'f', -1, 64)
		cStr := aStr + "," + bStr + "\n"
		f.WriteString(cStr)
	}

}
