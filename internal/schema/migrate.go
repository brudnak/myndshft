package schema

import (
	"github.com/GuiaBolso/darwin"
	"github.com/jmoiron/sqlx"
)

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Add patients",
		Script: `

CREATE TYPE add_info AS ( notes text, new_patient boolean, race text, ssn text); 

CREATE TABLE patients 
  ( 
     patient_id             UUID, 
     first_name             TEXT, 
     last_name              TEXT, 
     gender                 TEXT, 
     phone                  TEXT, 
     email_address          TEXT, 
     address                TEXT, 
     visit_date             TIMESTAMP, 
     diagnosis              TEXT, 
     drug_code              TEXT, 
     notes                  TEXT,
     new_patient            BOOLEAN,
     race                   TEXT,
     ssn                    TEXT,
     PRIMARY KEY (patient_id) 
  );`,
	},
}

func Migrate(db *sqlx.DB) error {

	driver := darwin.NewGenericDriver(db.DB, darwin.PostgresDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}
