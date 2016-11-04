package main

import (
	"sync"
	"net/http"
	"fmt"
	"log"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/counter", myCounter)
	log.Fatal(http.ListenAndServe("prog:8000", nil))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func myCounter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(w, "Count %d\n", count)
}