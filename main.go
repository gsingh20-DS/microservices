package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("This is my first microservices")
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("This is my good bye microservices")
	})

	http.ListenAndServe(":9090", nil)
}
