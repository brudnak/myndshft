package patient

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/google/uuid"
)

// List returns all Patients.
func List(db *sqlx.DB) ([]Patient, error) {
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

	if err := db.Select(&list, q); err != nil {
		return nil, err
	}

	return list, nil
}

// Retrieve returns a single Patient
func Retrieve(db *sqlx.DB, id string) (*Patient, error) {
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

	if err := db.Get(&p, q, id); err != nil {
		return nil, err
	}

	return &p, nil
}

// Create makes a new Patient.
func Create(db *sqlx.DB, np NewPatient) (*Patient, error) {
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

	if p.VisitDate == "" {
		if _, err := db.Exec(q, p.ID, p.FirstName, p.LastName, p.Gender, p.Phone, p.EmailAddress, p.Address, nil, p.Diagnosis, p.DrugCode, p.Notes, p.NewPatient, p.Race, p.Ssn); err != nil {
			return nil, errors.Wrapf(err, "inserting product: %v", np)
		}
		return &p, nil
	} else {
		if _, err := db.Exec(q, p.ID, p.FirstName, p.LastName, p.Gender, p.Phone, p.EmailAddress, p.Address, p.VisitDate, p.Diagnosis, p.DrugCode, p.Notes, p.NewPatient, p.Race, p.Ssn); err != nil {
			return nil, errors.Wrapf(err, "inserting product: %v", np)
		}
		return &p, nil
	}
}
