package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeLink(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response,"Welcome")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",homeLink)
	log.Fatal(http.ListenAndServe(":8080",router))

}
