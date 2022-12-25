package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kalpit-sharma-dev/url-shortner/src/errors"
	"github.com/kalpit-sharma-dev/url-shortner/src/model"
	"github.com/kalpit-sharma-dev/url-shortner/src/service"
)

type Handler struct {
	service service.UrlService
}

// NewHandler creates new handler  handler.
func NewHandler(service service.UrlService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HandleGetUrl(w http.ResponseWriter, r *http.Request) {

	fmt.Println("INFO : HandleGetUrl", r)
	v := mux.Vars(r)
	uri := v["url"]

	url, err := h.service.GetUrl(r.Context(), uri)
	if err != nil {
		log.Fatal("error", err)
		errors.HandleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(url)

}

func (h *Handler) HandleCreateUrl(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var req model.Request
	err := decoder.Decode(&req)
	if err != nil {
		log.Fatal("unable to decode", err)
		return
	}
	log.Println("INFO : HandleCreateUrl", r)

	url, err := h.service.CreateUrl(r.Context(), req)
	if err != nil {
		log.Fatal("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		errors.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(url)

}
