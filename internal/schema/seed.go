package schema

import (
	"github.com/jmoiron/sqlx"
)

// seeds is a string constant containing all of the queries needed to get the
// db seeded to a useful state for development.
//
// Using a constant in a .go file is an easy way to ensure the queries are part
// of the compiled executable and avoids pathing issues with the working
// directory. It has the downside that it lacks syntax highlighting and may be
// harder to read for some cases compared to using .sql files. You may also
// consider a combined approach using a tool like packr or go-bindata.
//
// Note that database servers besides PostgreSQL may not support running
// multiple queries as part of the same execution so this single large constant
// may need to be broken up.

const seeds = `
INSERT INTO patients (patient_id, first_name, last_name, gender, phone, email_address, address, visit_date, diagnosis, drug_code, notes, new_patient, race, ssn) VALUES
	('5677ec19-c833-4007-a181-fc0d345d3f9a', 'Andrew', 'Brudnak', 'Male', '480-881-2889', 'brudnak@protonmail.com', '3925 S Garrison', '2020-01-01 00:00:01.000001+00', 'GW', '1942', 'Patient is vegan', false, 'white', '602-87-9563'),
	('1f91efdb-c096-4c2b-b052-c496592ba30d', 'Russell', 'Brudnak', 'Male', '480-881-2889', 'brudnak@protonmail.com', '3925 S Garrison', '2020-01-01 00:00:01.000001+00', 'GW', '1942', 'Patient is vegan', false, 'white', '602-87-9563'),
	('335640fc-4787-4130-b0c9-6adb167b68e6', 'Ruby', 'Brudnak', 'Female', '480-881-2889', 'brudnak@protonmail.com', '3925 S Garrison', '2020-01-01 00:00:01.000001+00', 'GW', '1942', 'Patient is vegan', false, 'white', '602-87-9563'),
	('7dd87403-bf72-47bc-b434-7802bc0b8ecf', 'Jessica', 'Brudnak', 'Female', '480-881-2889', 'brudnak@protonmail.com', '3925 S Garrison', '2020-01-01 00:00:01.000001+00', 'GW', '1942', 'Patient is vegan', false, 'white', '602-87-9563')
	ON CONFLICT DO NOTHING;
`

// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func Seed(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(seeds); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
