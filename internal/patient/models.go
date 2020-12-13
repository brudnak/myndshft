package patient

// Patient is a hospital patient.
type Patient struct {
	ID           string `db:"patient_id" json:"id"`
	FirstName    string `db:"first_name" json:"first_name"`
	LastName     string `db:"last_name" json:"last_name"`
	Gender       string `db:"gender" json:"gender"`
	Phone        string `db:"phone" json:"phone"`
	EmailAddress string `db:"email_address" json:"email_address"`
	Address      string `db:"address" json:"address"`
	VisitDate    string `db:"visit_date" json:"visit_date"`
	Diagnosis    string `db:"diagnosis" json:"diagnosis"`
	DrugCode     string `db:"drug_code" json:"drug_code"`
	Notes        string `db:"notes" json:"notes"`
	NewPatient   bool   `db:"new_patient" json:"new_patient"`
	Race         string `db:"race" json:"race"`
	Ssn          string `db:"ssn" json:"ssn"`
}

// NewPatient is what we require from clients to make a new Patient.
type NewPatient struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Phone        string `json:"phone"`
	EmailAddress string `json:"email_address"`
	Address      string `json:"address"`
	VisitDate    string `json:"visit_date"`
	Diagnosis    string `json:"diagnosis"`
	DrugCode     string `json:"drug_code"`
	Notes        string `json:"notes"`
	NewPatient   bool   `json:"new_patient"`
	Race         string `json:"race"`
	Ssn          string `json:"ssn"`
}
