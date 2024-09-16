package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/ptflp/godecoder"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	c := NewController(NewResponder(godecoder.NewDecoder(), nil))

	r.Post("/api/address/search", c.SearchHandler)
	r.Post("/api/address/geocode", c.GeocodeHandler)

	http.ListenAndServe(":8080", r)
}

type Controller struct {
	responder Responder
}

func NewController(responder Responder) *Controller {
	return &Controller{
		responder: responder,
	}
}

func (c Controller) SearchHandler(w http.ResponseWriter, r *http.Request) {
	var s RequestAddressSearch
	json.NewDecoder(r.Body).Decode(&s)
	url := "https://nominatim.openstreetmap.org/search?q=" + s.Query + "&format=json"

	resp, err := http.Get(url)
	if err != nil {
		c.responder.ErrorBadRequest(w, err)
	}
	defer resp.Body.Close()

	tmp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	c.responder.OutputJSON(w, tmp)
}

func (c Controller) GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var s RequestAddressGeocode
	json.NewDecoder(r.Body).Decode(&s)
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", s.Lat, s.Lng)

	resp, err := http.Get(url)
	if err != nil {
		c.responder.ErrorBadRequest(w, err)
	}
	defer resp.Body.Close()

	tmp, err := io.ReadAll(resp.Body)
	if err != nil {
		c.responder.ErrorBadRequest(w, err)
	}

	c.responder.OutputJSON(w, tmp)

}

type RequestAddressSearch struct {
	Query string `json:"query"`
}

type RequestAddressGeocode struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type MyController struct {
	responder Responder
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})

	ErrorUnauthorized(w http.ResponseWriter, err error)
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorForbidden(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type Respond struct {
	log *zap.Logger
	godecoder.Decoder
}

func NewResponder(decoder godecoder.Decoder, logger *zap.Logger) Responder {
	return &Respond{log: logger, Decoder: decoder}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type",
		"application/json;charset=utf-8")
	if err := r.Encode(w, responseData); err != nil {
		r.log.Error("responder json encode error", zap.Error(err))
	}
}

func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	r.log.Info("http response bad request status code", zap.Error(err))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := r.Encode(w, Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		r.log.Info("response writer error on write", zap.Error(err))
	}
}

func (r *Respond) ErrorUnauthorized(w http.ResponseWriter, err error) {
	r.log.Info("http response unauthorized status code", zap.Error(err))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	if err := r.Encode(w, Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		r.log.Error("response writer error on write", zap.Error(err))
	}
}

func (r *Respond) ErrorForbidden(w http.ResponseWriter, err error) {
	r.log.Warn("http resposne forbidden", zap.Error(err))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusForbidden)
	if err := r.Encode(w, Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		r.log.Error("response writer error on write", zap.Error(err))
	}
}

func (r *Respond) ErrorInternal(w http.ResponseWriter, err error) {
	if errors.Is(err, context.Canceled) {
		return
	}
	r.log.Error("http response internal error", zap.Error(err))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := r.Encode(w, Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		r.log.Error("response writer error on write", zap.Error(err))
	}
}
