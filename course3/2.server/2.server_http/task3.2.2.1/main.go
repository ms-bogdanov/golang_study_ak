package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	})

	http.ListenAndServe(":8080", nil)
}
