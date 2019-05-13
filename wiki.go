package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var args string = r.URL.Path[1:]
	if len(args) == 0 {
		args = "Cloud Run"
	}
	fmt.Fprintf(w, "Hi there, I love %s!", args)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
