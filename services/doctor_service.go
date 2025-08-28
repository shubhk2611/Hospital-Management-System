package services

import (
	"time"

	"hospital/database"
	"hospital/models"
	"hospital/utils"

	"gorm.io/gorm"
)

type DoctorService struct{}

func NewDoctorService() *DoctorService {
	return &DoctorService{}
}

// CreateDoctor creates a new doctor in the database
func (ds *DoctorService) CreateDoctor(payload models.DoctorPayload) (*models.Doctor, error) {
	doctor := models.Doctor{
		ID:        utils.GenerateID(),
		Name:      payload.Name,
		ContactNo: payload.ContactNo,
	}

	result := database.GetDB().Create(&doctor)
	if result.Error != nil {
		return nil, result.Error
	}

	return &doctor, nil
}

// GetDoctorByID retrieves a doctor by ID from the database
func (ds *DoctorService) GetDoctorByID(doctorID string) (*models.Doctor, error) {
	var doctor models.Doctor
	result := database.GetDB().First(&doctor, "ID = ?", doctorID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &doctor, nil
}

// UpdateDoctorContact updates a doctor's contact number
func (ds *DoctorService) UpdateDoctorContact(doctorID string, payload models.UpdateDoctorPayload) (*models.Doctor, error) {
	var doctor models.Doctor

	// First, check if the doctor exists
	result := database.GetDB().First(&doctor, "ID = ?", doctorID)
	if result.Error != nil {
		return nil, result.Error
	}

	// Perform the update
	updates := map[string]interface{}{
		"contact_no": payload.ContactNo,
		"updated_at": time.Now().Unix(),
	}

	result = database.GetDB().Model(&doctor).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// Re-fetch the updated record
	database.GetDB().First(&doctor, "ID = ?", doctorID)
	return &doctor, nil
}
