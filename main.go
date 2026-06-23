package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", homeHandler)
	mux.HandleFunc("/customer", CustomerHandler)
	mux.HandleFunc("/waiter", waiterHandler)
	fmt.Println("starting the server on port :3000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Customer")
}

func waiterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "What can we get you")
}
