package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, drone!")

	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}

func loadPage1() (content []byte) {
	templ, err := ioutil.ReadFile("template.html")
	//str1 := "here is your string"
	if err != nil {
		log.Println("error")
	}
	return templ
}

func handler1(w http.ResponseWriter, r *http.Request) {
	content := loadPage1()
	//fmt.Fprintf(w, "Drone App\n\n%s\n\n", r.URL.Path[1:])
	fmt.Fprintf(w, string(content))
}
