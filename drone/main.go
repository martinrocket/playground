package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, drone!")
	createPage1()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/static", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func createPage1() {
	myTmpl, _ := ioutil.ReadFile("template.html")
	newfile, _ := os.Create("/static/output.html")
	newfile.Close()
	newfile, _ = os.OpenFile("/static/output.html", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Println(string(myTmpl))
	t, _ := template.ParseFiles("template.html")
	t.Execute(newfile, "my data gained")
	newfile.Close()

}
