package main

import (
	"fmt"
	"log"
	"net/http"
)

// formhandler function
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s \n", name)
	fmt.Fprintf(w, "Address = %s \n", address)
}

// hellohandler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Println(w, " hello")
}

func main() {
	// initialize a web server with the http protocol
	fileServer := http.FileServer(http.Dir("./static"))

	// this will serve the index.html file
	http.Handle("/", fileServer)

	// show us the form.html
	http.HandleFunc("/form", formHandler)

	// show the print hello on the webpage
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port : 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
