package responder

import (
	"GeoAPI/internal/service/models"
	"encoding/json"
	"net/http"
)

type Responder interface {
	RespondWithErr(w http.ResponseWriter, err error)
	RespondWithOk(w http.ResponseWriter, data interface{})
}

type Respond struct {
	AddressSearch []*models.AddressSearch
	AddressGeo    *models.AddressGeo
}

func (r *Respond) RespondWithErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func (r *Respond) RespondWithOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func NewRespond() Respond {
	return Respond{}
}
