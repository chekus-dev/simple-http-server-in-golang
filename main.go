package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// messageResponse is the JSON shape returned by each endpoint.
type messageResponse struct {
	Message string `json:"message"`
}

func main() {
	// Build the router and start the server on port 3000.
	mux := newServerMux()
	fmt.Println("starting the server on port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

// newServerMux registers all routes used by this simple API.
func newServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/user", homeHandler)
	mux.HandleFunc("/customer", customerHandler)
	mux.HandleFunc("/waiter", waiterHandler)
	return mux
}

// rootHandler returns the main welcome message and only accepts GET requests.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messageResponse{Message: "Welcome"}); err != nil {
		http.Error(w, "unable to encode response", http.StatusInternalServerError)
	}
}

// homeHandler handles the user welcome route.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messageResponse{Message: "Welcome"}); err != nil {
		http.Error(w, "unable to encode response", http.StatusInternalServerError)
	}
}

// customerHandler returns a simple greeting for customers.
func customerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messageResponse{Message: "Hello Customer"}); err != nil {
		http.Error(w, "unable to encode response", http.StatusInternalServerError)
	}
}

// waiterHandler returns a quick message for a customer asking for help.
func waiterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messageResponse{Message: "What can we get you"}); err != nil {
		http.Error(w, "unable to encode response", http.StatusInternalServerError)
	}
}
