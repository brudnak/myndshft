package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brudnak/myndshft/internal/patient"
	"github.com/jmoiron/sqlx"
)

// PatientService has handler methods for dealing with Patients
type Patient struct {
	DB *sqlx.DB
}

// List returns all patients as a list
func (p *Patient) List(w http.ResponseWriter, r *http.Request) {

	list, err := patient.List(p.DB)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error querying db", err)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error marshalling", err)
		return
	}

	w.Header().Set("content-type", "application/json: charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Println("error writing")
	}
}
