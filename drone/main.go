package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, drone!")

	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}

func loadPage1() (content []byte) {

	t := template.New("homePage")
	t, err := t.ParseFiles("template.html")
	if err != nil {
		log.Println("Error parsing home p")
	}

}

func handler1(w http.ResponseWriter, r *http.Request) {
	content := loadPage1()
	//fmt.Fprintf(w, "Drone App\n\n%s\n\n", r.URL.Path[1:])
	fmt.Fprintf(w, string(content))
}
