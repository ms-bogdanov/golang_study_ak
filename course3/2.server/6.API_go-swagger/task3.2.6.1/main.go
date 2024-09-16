package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"log"
	"net/http"
)

// @title GeoService
// @version 1.0
// @description Simple GeoService.

// @host localhost:8080
// @BasePath /
func main() {
	r := chi.NewRouter()

	r.Post("/api/address/search", SearchHandler)
	r.Post("/api/address/geocode", GeocodeHandler)

	http.ListenAndServe(":8080", r)
}

// SearchHandler handle POST-requests on api/address/search
// @Summary SearchCity
// @Tags Search
// @Description Search city Name by coords
// @Accept  json
// @Produce  json
// @Param  coordinates  body  RequestAddressSearch true  "Lattitude and Longitude"
// @Success 200 {object} string
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/address/search [post]
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

// GeocodeHandler handle POST-requests on api/address/geocode
// @Summary SearchCity
// @Tags Search
// @Description Search city Name by coords
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/address/geocode [post]
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

type errorResponse struct {
	Message string `json:"message"`
}
