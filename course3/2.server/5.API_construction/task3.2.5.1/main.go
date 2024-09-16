package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Post("/api/address/search", SearchHandler)
	r.Post("/api/address/geocode", GeocodeHandler)

	http.ListenAndServe(":8080", r)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var s RequestAddressSearch
	json.NewDecoder(r.Body).Decode(&s)
	url := "https://nominatim.openstreetmap.org/search?q=" + s.Query + "&format=json"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	tmp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	w.Write(tmp)
}

func GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var c RequestAddressGeocode
	json.NewDecoder(r.Body).Decode(&c)
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", c.Lat, c.Lng)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	tmp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	w.Write(tmp)
}

type RequestAddressSearch struct {
	Query string `json:"query"`
}

type RequestAddressGeocode struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
