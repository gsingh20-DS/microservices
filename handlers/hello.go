package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHttp(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("This is my first microservices")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "This is error in the code", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("This is error"))
		return
	}
	log.Printf("This is the data came in body %s", d)

	fmt.Fprintf(rw, "This is back from response writer %s", d)

}
