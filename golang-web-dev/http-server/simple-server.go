package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handlerEarth)

	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Universe")
}

func handlerEarth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth")
}
