package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"os"
	"strconv"
)

const (
	DEFAULT_PORT = 8000  
	DEFAULT_HOST = "localhost"
)

var mu sync.Mutex
var count int

func main() {
	var port int
	var err error

	if len(os.Args) > 1 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Not a valid number for port %s. Using default port %d instead.\n", os.Args[1], DEFAULT_PORT)
			port = DEFAULT_PORT
		}
	} else {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", DEFAULT_HOST, port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL Path: %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}
