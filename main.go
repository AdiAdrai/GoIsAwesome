package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Hello From Go!<br />I'm Alive!!!</h1>")
}

func HelloNameHandler(w http.ResponseWriter, req *http.Request) {

	param := mux.Vars(req)

	name := "World"

	if len(param["name"]) > 0 {
		name = param["name"]
	}
	fmt.Fprintf(w, "<h1>Hello %s!!!</h1>", name)
}

func HelloWorldHandler(w http.ResponseWriter, req *http.Request) {

	name := "World"

	fmt.Fprintf(w, "<h1>Hello %s!!!</h1>", name)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/hello/{name}", HelloNameHandler).Methods("GET")
	router.HandleFunc("/hello", HelloWorldHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
