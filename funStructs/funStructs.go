package main

import (
	"encoding/json"
	"fmt"

)

type myStruct struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AltName string `json:"alt"`
	MyItems string `json:"items"`
}

func main() {

	a := myStruct{1, "First", "1st", "hello"}
	fmt.Println(a)
	s, err := a.makeJson()
	fmt.Println(s)
	fmt.Printf("%T %T \n", s, err)
	unJson(s)


}

func (m myStruct) makeJson() (string, error) {

	c, err := json.MarshalIndent(m, " ", "   ")
	return string(c), err

}

func  unJson(s string)  {

	fmt.Println(s)
	st := json.Unmarshal(s, myStruct{})
	fmt.Println()
	return



}
