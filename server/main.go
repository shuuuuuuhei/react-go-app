package main

import (
	"github.com/react-go-app/backend/infrastructure"
	"fmt"
	"net/http"
)

func main() {
	infrastructure.Init()
}

func echoHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}
