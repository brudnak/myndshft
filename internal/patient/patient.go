package patient

import "github.com/jmoiron/sqlx"

// List returns all Patients.
func List(db *sqlx.DB) ([]Patient, error) {
	list := []Patient{}

	const q = `SELECT * FROM patients`

	if err := db.Select(&list, q); err != nil {
		return nil, err
	}

	return list, nil
}
