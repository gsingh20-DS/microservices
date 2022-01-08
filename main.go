package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gsingh20-DS/microservices/tree/main/handlers"
)

func main() {

	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("This is my good bye microservices")
	// })

	l := log.New(os.Stdout, "This is from logger", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("goodbye", gb)
	http.ListenAndServe(":9090", nil)
}
