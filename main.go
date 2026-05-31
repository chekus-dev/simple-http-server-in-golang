package main 

import (
	"fmt"
	"net/http"
	"log"
)

func Home(w http.ResponseWriter, res *http.Request) {
	fmt.Fprint(w,"<h1> welcome to simple http server in golang <h1>")
}

func main(){
	http.HandleFunc("/", Home)
	fmt.Println("starting simple http server in golang on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
