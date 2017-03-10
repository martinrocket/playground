package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"github.com/golang/protobuf/proto"
	"/playground"


)

type MyList struct {
	Id     int            `json:"id"`
	Status string         `json:"status"`
	Table  string         `json:"table"`
	Name   string         `json:"name"`
	Items  map[int]string `json:"items"`
}


func main() {
	fmt.Println("Hello Gophers!")
	myTime := time.Now()
	fmt.Println(myTime.Format("Jan _2 15:04:05.000"))
	fmt.Println(myTime.Format("01_02_15:04:05.000"))
	fmt.Println(myTime.Format("01_02_2006"))
	writeJson()
	writeProto()
	fmt.Print("Enter text: ")

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

func writeJson() {

	s := MyList{
		Id:     1,
		Status: "open",
		Table:  "32",
		Name:   "Jana Miller",
		Items: map[int]string{
			1: "Coke",
			2: "Potatoes",
			3: "Pie",
		},
	}
	jCk, err := json.MarshalIndent(s, " ", "    ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jCk))
}

func writeProto() {
	test := new(playground.CheckProto)
	test.ID = proto.Int32(25)
	test.Status =  proto.String("open")
	test.Table = proto.String("35")
	test.Name =  proto.String("Jana Miller")
	test.Items = proto.String("Potatoes")


	data, err :=proto.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}


