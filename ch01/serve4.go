package main

import (
	"log"
	"net/http"
	"golang/ch01/main"
)

func main() {
	handle := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
