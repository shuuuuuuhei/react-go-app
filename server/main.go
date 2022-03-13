package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", echoHello)
	http.ListenAndServe(":4000", nil)
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}