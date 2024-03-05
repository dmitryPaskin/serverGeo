package responder

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	RespondWithErr(w http.ResponseWriter, err error)
	RespondWithOk(w http.ResponseWriter, data interface{})
}

type respond struct {
}

func New() respond {
	return respond{}
}

func (r respond) RespondWithErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func (r respond) RespondWithOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
