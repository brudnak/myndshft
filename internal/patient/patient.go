package patient

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/google/uuid"
)

var (
	ErrNotFound  = errors.New("patient not found")
	ErrInvalidID = errors.New("id provided was not a valid UUID")
)

// List returns all Patients.
func List(ctx context.Context, db *sqlx.DB) ([]Patient, error) {
	list := []Patient{}

	const q = `
SELECT patient_id, 
       first_name, 
       last_name, 
       gender, 
       phone, 
       email_address, 
       address, 
       visit_date, 
       diagnosis, 
       drug_code, 
       notes, 
       new_patient, 
       race, 
       ssn 
FROM   patients;`

	if err := db.SelectContext(ctx, &list, q); err != nil {
		return nil, err
	}

	return list, nil
}

// Retrieve returns a single Patient
func Retrieve(ctx context.Context, db *sqlx.DB, id string) (*Patient, error) {

	if _, err := uuid.Parse(id); err != nil {
		return nil, ErrInvalidID
	}

	var p Patient

	const q = `
SELECT patient_id, 
       first_name, 
       last_name, 
       gender, 
       phone, 
       email_address, 
       address, 
       visit_date, 
       diagnosis, 
       drug_code, 
       notes, 
       new_patient, 
       race, 
       ssn 
FROM   patients
WHERE patient_id = $1;`

	if err := db.GetContext(ctx, &p, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &p, nil
}

// Create makes a new Patient.
func Create(ctx context.Context, db *sqlx.DB, np NewPatient) (*Patient, error) {
	p := Patient{
		ID:           uuid.New().String(),
		FirstName:    np.FirstName,
		LastName:     np.LastName,
		Gender:       np.Gender,
		Phone:        np.Phone,
		EmailAddress: np.EmailAddress,
		Address:      np.Address,
		VisitDate:    np.VisitDate,
		Diagnosis:    np.Diagnosis,
		DrugCode:     np.DrugCode,
		Notes:        np.Notes,
		NewPatient:   np.NewPatient,
		Race:         np.Race,
		Ssn:          np.Ssn,
	}

	const q = `
INSERT INTO patients
(patient_id, first_name, last_name, gender, phone, email_address, address, visit_date, diagnosis, drug_code, notes, new_patient, race, ssn)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`

	if _, err := db.ExecContext(ctx, q, p.ID, p.FirstName, p.LastName, p.Gender, p.Phone, p.EmailAddress, p.Address, p.VisitDate, p.Diagnosis, p.DrugCode, p.Notes, p.NewPatient, p.Race, p.Ssn); err != nil {
		return nil, errors.Wrapf(err, "inserting product: %v", np)
	}
	return &p, nil

}

func Update(ctx context.Context, db *sqlx.DB, id string, update UpdatePatient) error {

	pat, err := Retrieve(ctx, db, id)
	if err != nil {
		return err
	}

	if update.FirstName != nil {
		pat.FirstName = *update.FirstName
	}
	if update.LastName != nil {
		pat.LastName = *update.LastName
	}
	if update.Gender != nil {
		pat.Gender = *update.Gender
	}
	if update.Phone != nil {
		pat.Phone = *update.Phone
	}
	if update.EmailAddress != nil {
		pat.EmailAddress = *update.EmailAddress
	}
	if update.Address != nil {
		pat.Address = *update.Address
	}
	if update.VisitDate != nil {
		pat.VisitDate = *update.VisitDate
	}
	if update.Diagnosis != nil {
		pat.Diagnosis = *update.Diagnosis
	}
	if update.DrugCode != nil {
		pat.DrugCode = *update.DrugCode
	}
	if update.Notes != nil {
		pat.Notes = *update.Notes
	}
	if update.NewPatient != nil {
		pat.NewPatient = *update.NewPatient
	}
	if update.Race != nil {
		pat.Race = *update.Race
	}
	if update.Ssn != nil {
		pat.Ssn = *update.Ssn
	}

	const q = `
UPDATE patients SET
       "first_name" = $2,
       "last_name" = $3,
       "gender" = $4,
       "phone" = $5,
       "email_address" = $6,
       "address" = $7,
       "visit_date" = $8,
       "diagnosis" = $9,
       "drug_code" = $10,
       "notes" = $11,
       "new_patient" = $12,
       "race" = $13,
       "ssn" = $14
WHERE patient_id = $1;`

	_, err = db.ExecContext(ctx, q, id, pat.FirstName, pat.LastName, pat.Gender, pat.Phone, pat.EmailAddress, pat.Address, pat.VisitDate, pat.Diagnosis, pat.DrugCode, pat.Notes, pat.NewPatient, pat.Race, pat.Ssn)
	if err != nil {
		return errors.Wrap(err, "updating patient")
	}
	return nil
}
