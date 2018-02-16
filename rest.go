package main

import (
	"net/http"
	"log"
)

func main(){

	router := NewRouter()

	server := http.ListenAndServe(":666",router)

	log.Fatal(server)

}
