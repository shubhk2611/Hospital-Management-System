package models

// DoctorPayload defines the expected JSON structure for creating a doctor.
type DoctorPayload struct {
	Name      string `json:"name" binding:"required"`
	ContactNo string `json:"contact_no" binding:"required,len=10"`
}

// Doctor represents the "doctors" table in the database.
type Doctor struct {
	ID        string `gorm:"primaryKey;size:5"`
	Name      string `gorm:"varchar(255)"`
	ContactNo string `gorm:"char(10);unique"`
	CreatedAt int64
	UpdatedAt int64
}

// UpdateDoctorPayload defines the fields allowed for a doctor update.
type UpdateDoctorPayload struct {
	ContactNo string `json:"contact_no" binding:"required"`
}
