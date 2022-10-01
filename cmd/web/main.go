package main

import (
	"fmt"
	"net/http"

	"github.com/marceesty/goWeb/handlers"
)

const portNumber = ":8080"


func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting app on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}