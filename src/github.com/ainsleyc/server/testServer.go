package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
  Greetings string
  Punct string
  Who string
}

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s) 
}

func (s Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s) 
}

func main() {
  http.Handle("/string", String("This is a string"))
  http.Handle("/struct", &Struct{"Hello", ":", "Go"})

	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
