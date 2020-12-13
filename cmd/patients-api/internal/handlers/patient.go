package handlers

import (
	"log"
	"net/http"

	"github.com/brudnak/myndshft/internal/patient"
	"github.com/brudnak/myndshft/internal/platform/web"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// PatientService has handler methods for dealing with Patients
type Patient struct {
	DB  *sqlx.DB
	Log *log.Logger
}

// List returns all patients as a list
func (p *Patient) List(w http.ResponseWriter, r *http.Request) error {

	list, err := patient.List(r.Context(), p.DB)

	if err != nil {
		return err
	}

	return web.Respond(w, list, http.StatusOK)
}

// Retrieve gives a single Patient.
func (p *Patient) Retrieve(w http.ResponseWriter, r *http.Request) error {

	id := chi.URLParam(r, "id")

	pat, err := patient.Retrieve(r.Context(), p.DB, id)

	if err != nil {
		switch err {
		case patient.ErrNotFound:
			return web.NewRequestError(err, http.StatusNotFound)
		case patient.ErrInvalidID:
			return web.NewRequestError(err, http.StatusBadRequest)
		default:
			return errors.Wrapf(err, "looking for patient %q", id)
		}
	}

	return web.Respond(w, pat, http.StatusOK)
}

// Update
func (p *Patient) Update(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	var update patient.UpdatePatient
	if err := web.Decode(r, &update); err != nil {
		return errors.Wrap(err, "decoding patient update")
	}

	if err := patient.Update(r.Context(), p.DB, id, update); err != nil {
		switch err {
		case patient.ErrNotFound:
			return web.NewRequestError(err, http.StatusNotFound)
		case patient.ErrInvalidID:
			return web.NewRequestError(err, http.StatusBadRequest)
		default:
			return errors.Wrapf(err, "updating patient %q", id)
		}
	}

	return web.Respond(w, nil, http.StatusNoContent)
}

// Create decode a JSON document from a POST request and create a new Patient.
func (p *Patient) Create(w http.ResponseWriter, r *http.Request) error {

	var np patient.NewPatient
	if err := web.Decode(r, &np); err != nil {
		return err
	}

	pat, err := patient.Create(r.Context(), p.DB, np)
	if err != nil {
		return err
	}

	return web.Respond(w, pat, http.StatusCreated)
}

func (p *Patient) Delete(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	if err := patient.Delete(r.Context(), p.DB, id); err != nil {
		switch err {
		case patient.ErrInvalidID:
			return web.NewRequestError(err, http.StatusBadRequest)
		default:
			return errors.Wrapf(err, "deleting patient %q", id)
		}
	}

	return web.Respond(w, nil, http.StatusNoContent)
}