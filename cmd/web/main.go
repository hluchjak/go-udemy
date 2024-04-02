package main

import (
	"net/http"
	"website/pkg/handlers"
)

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	_, _ = fmt.Fprintf(w, "Hello, world!")
	//
	//})

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(":8080", nil)
}
