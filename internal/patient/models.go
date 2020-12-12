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
