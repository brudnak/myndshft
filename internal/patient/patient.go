package patient

import "github.com/jmoiron/sqlx"

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
