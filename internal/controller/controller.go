package controller

import (
	"GeoAPI/internal/controller/responder"
	"GeoAPI/internal/service"
	"GeoAPI/internal/service/models"
	"encoding/json"
	"net/http"
)

type Handler struct {
	s service.Service
	r responder.Responder
}

func New(serv service.Service, respond responder.Responder) Handler {
	return Handler{
		s: serv,
		r: respond,
	}
}

// @Summary Search for an address
// @ID searchAddress
// @Tags search
// @Accept  json
// @Produce  json
// @Param request body models.SearchRequest true "Search request"
// @Success 200 {object} models.AddressSearch
// @Failure 500
// @Router /address/search [post]
func (h *Handler) SearchAddressHandler(w http.ResponseWriter, r *http.Request) {
	var searchRequest models.SearchRequest
	err := json.NewDecoder(r.Body).Decode(&searchRequest)
	if err != nil {
		h.r.RespondWithErr(w, err)
		return
	}
	address, err := h.s.Address(searchRequest)
	if err != nil {
		h.r.RespondWithErr(w, err)
		return
	}
	h.r.RespondWithOk(w, address)
}

// @Summary Geocode an address
// @ID geocodeAddress
// @Tags geocode
// @Accept  json
// @Produce  json
// @Param request body models.GeocodeRequest true "Geocode request"
// @Success 200 {object} models.AddressGeo
// @Failure 500
// @Router /address/geocode [post]
func (h *Handler) GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var geocodeRequest models.GeocodeRequest
	err := json.NewDecoder(r.Body).Decode(&geocodeRequest)
	if err != nil {
		h.r.RespondWithErr(w, err)
		return
	}
	geocode, err := h.s.Geocode(geocodeRequest)
	if err != nil {
		h.r.RespondWithErr(w, err)
		return
	}
	h.r.RespondWithOk(w, geocode)
}
