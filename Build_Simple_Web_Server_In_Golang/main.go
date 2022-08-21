package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "ParseForm(), err: %v", err)
		return
	}
	fmt.Fprintf(response, "POST request successful")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(response, "Name = %s\n", name)
	fmt.Fprintf(response, "Address = %s\n", address)
}

func helloHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.RawPath != "/hello" { // I changed to .url.RawPath - it works now
		http.Error(response, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(response, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Printf("hello!") // i removed the response - just the hello string - it works
}

func main() {
	fileServer := http.FileServer(http.Dir("./Static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
