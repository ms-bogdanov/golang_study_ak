package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {

	r := chi.NewRouter()

	r.Get("/1", Hw)
	r.Get("/2", Hw2)
	r.Get("/3", Hw3)
	r.Get("/", Nf)

	http.ListenAndServe(":8080", r)
}

func Hw(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}

func Hw2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world 2"))
}

func Hw3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world 3"))
}

func Nf(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}
