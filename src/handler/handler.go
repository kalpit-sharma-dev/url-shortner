package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	fmt.Println("HandleGetUrl", r)

	products, err := h.service.GetUrl(r.Context(), "")
	if err != nil {
		log.Fatal("error", err)
		return
	}
	json.NewEncoder(w).Encode(products)

}

func (h *Handler) HandleCreateUrl(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var req model.Request
	err := decoder.Decode(&req)
	if err != nil {
		log.Fatal("unable to decode", err)
	}
	fmt.Println("HandleCreateUrl", r)

	fmt.Println("1111111111111111111111111111111111111111111111111111111")
	products, err := h.service.CreateUrl(r.Context(), req)
	if err != nil {
		log.Fatal("error", err)
		return
	}
	json.NewEncoder(w).Encode(products)

}
