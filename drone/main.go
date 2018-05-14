package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	
)

func main() {
	fmt.Println("Hello, drone!")
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":8080", nil)
}

func createPage1() {
	fmt.Println("this also happened")
	myTmpl, _ := ioutil.ReadFile("template.html")
	f, _ := os.OpenFile("output.html", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Println(string(myTmpl))
	t, _ := template.ParseFiles("template.html")
	t.Execute(f, "my data gain")
	f.Close()

}

func handler1(w http.ResponseWriter, r *http.Request) {
	createPage1()
	webPage :=
}
