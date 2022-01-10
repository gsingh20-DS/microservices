package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gsingh20-DS/microservices/handlers"
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
	sm.Handle("/goodbye", gb)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s.ListenAndServe()
}
