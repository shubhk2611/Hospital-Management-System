package models

// PatientPayload defines the expected JSON for creating a patient.
type PatientPayload struct {
	Name      string `json:"name" binding:"required"`
	ContactNo string `json:"contact_no" binding:"required"`
	Address   string `json:"address" binding:"required"`
	DoctorID  string `json:"doctor_id" binding:"required"`
}

// UpdatePatientPayload defines the optional fields for a patient update.
type UpdatePatientPayload struct {
	ContactNo *string `json:"contact_no"`
	Address   *string `json:"address"`
	DoctorID  *string `json:"doctor_id"`
}

// Patient represents the "patients" table in the database.
type Patient struct {
	ID        string `gorm:"primaryKey;size:5"`
	Name      string `gorm:"varchar(255)"`
	ContactNo string `gorm:"char(10)"`
	Address   string `gorm:"varchar(255)"`
	DoctorID  string `gorm:"char(5)"`
	CreatedAt int64
	UpdatedAt int64
}
