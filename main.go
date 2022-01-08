package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("This is my good bye microservices")
	// })

	l := log.New(os.Stdout, "This is from logger", log.LstdFlags)

	http.ListenAndServe(":9090", nil)
}
