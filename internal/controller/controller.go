package controller

import (
	"GeoAPI/internal/controller/Responder"
	"GeoAPI/internal/service"
	"GeoAPI/internal/service/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Call interface {
	callDataSearch(w http.ResponseWriter)
	callDataGeo(w http.ResponseWriter)
}

type Request struct {
	models.GeocodeRequest
	models.SearchRequest
}

func (sr *Request) callDataSearch(w http.ResponseWriter) {
	url := "https://cleaner.dadata.ru/api/v1/clean/address"
	requestData := []string{sr.Query}

	response := responder.NewRespond()

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		response.RespondWithErr(w, err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		response.RespondWithErr(w, err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Token 9a84b6e525fb548e7170b77175e9e15af84a30ac")
	req.Header.Add("X-Secret", "6ecfe8510311d14daf5de31de9a5af4ceeb5b0d5")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		response.RespondWithErr(w, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.RespondWithErr(w, err)
		return
	}

	convertor := service.NewServiceConvertor(body)
	response.AddressSearch, err = convertor.ConversionFromSourceToSearch()
	if err != nil {
		response.RespondWithErr(w, err)
		return
	}

	response.RespondWithOk(w, response.AddressSearch)
}

func (sr *Request) callDataGeo(w http.ResponseWriter) {
	url := "http://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address"

	response := responder.NewRespond()
	requestBody, err := json.Marshal(sr.GeocodeRequest)
	if err != nil {
		response.RespondWithErr(w, err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		response.RespondWithErr(w, err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Token 9a84b6e525fb548e7170b77175e9e15af84a30ac")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		response.RespondWithErr(w, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.RespondWithErr(w, err)
		return
	}

	convertor := service.NewServiceConvertor(body)

	response.AddressGeo, err = convertor.ConversionFromSourceToGeo()
	if err != nil {
		response.RespondWithErr(w, err)
		return
	}

	response.RespondWithOk(w, response.AddressGeo)
}

// @Summary Search for an address
// @ID searchAddress
// @Tags search
// @Accept  json
// @Produce  json
// @Param request body models.SearchRequest true "Search Request"
// @Success 200 {object} models.AddressSearch
// @Failure 500
// @Security ApiKeyAuth
// @Router /address/search [post]
func SearchAddressHandler(w http.ResponseWriter, r *http.Request) {
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request.SearchRequest)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	request.callDataSearch(w)
}

// @Summary Geocode an address
// @ID geocodeAddress
// @Tags geocode
// @Accept  json
// @Produce  json
// @Param request body models.GeocodeRequest true "Geocode Request"
// @Success 200 {object} models.AddressGeo
// @Failure 500
// @Security ApiKeyAuth
// @Router /address/geocode [post]
func GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request.GeocodeRequest)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	request.callDataGeo(w)
}
