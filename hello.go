package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Goodbye, World!"))
	})
	http.HandleFunc("/buy", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Buy, World!"))
	})
	http.ListenAndServe(":8088", nil)
}
