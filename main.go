package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print(err.Error())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET", http.StatusNotFound)
	}

	fmt.Fprintf(w, "HelloWorld!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Could not parse the form!", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Received the form!\n")

	responseMessage := "Name: " + r.FormValue("name") + " | Address: " + r.FormValue("address")
	fmt.Fprint(w, responseMessage)
}
