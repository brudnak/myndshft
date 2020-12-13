package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brudnak/myndshft/internal/patient"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

// PatientService has handler methods for dealing with Patients
type Patient struct {
	DB  *sqlx.DB
	Log *log.Logger
}

// List returns all patients as a list
func (p *Patient) List(w http.ResponseWriter, r *http.Request) {

	list, err := patient.List(p.DB)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Println("error querying db", err)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Println("error marshalling", err)
		return
	}

	w.Header().Set("content-type", "application/json: charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		p.Log.Println("error writing")
	}
}

// Retrieve gives a single Patient.
func (p *Patient) Retrieve(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	pat, err := patient.Retrieve(p.DB, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Println("error querying db", err)
		return
	}

	data, err := json.Marshal(pat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Println("error marshalling", err)
		return
	}

	w.Header().Set("content-type", "application/json: charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		p.Log.Println("error writing")
	}
}

// Create decode a JSON document from a POST request and create a new Patient.
func (p *Patient) Create(w http.ResponseWriter, r *http.Request) {

	var np patient.NewPatient
	if err := json.NewDecoder(r.Body).Decode(&np); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		p.Log.Println(err)
		return
	}

	pat, err := patient.Create(p.DB, np)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Println("error querying db", err)
		return
	}

	data, err := json.Marshal(pat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.Log.Println("error marshalling", err)
		return
	}

	w.Header().Set("content-type", "application/json: charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(data); err != nil {
		p.Log.Println("error writing")
	}
}
