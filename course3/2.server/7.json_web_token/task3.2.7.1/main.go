package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"time"
)

var users = make(map[string]string)

func main() {
	r := chi.NewRouter()

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/api/address/search", SearchHandler)
		r.Post("/api/address/geocode", GeocodeHandler)
	})

	r.Post("/api/register", Registration)
	r.Post("/api/login", Login)

	http.ListenAndServe(":8080", r)
}

func Registration(w http.ResponseWriter, r *http.Request) {
	var newUser User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}
	defer r.Body.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	}

	users[newUser.Username] = string(hash)
	w.WriteHeader(http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	var newUser User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}
	defer r.Body.Close()

	hashedPass, ok := users[newUser.Username]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(newUser.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"username": newUser.Username,
		"exp":      jwtauth.ExpireIn(time.Hour),
	}
	_, tokenString, _ := tokenAuth.Encode(claims)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
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

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
