package services

import (
	"time"

	"hospital/models"
	"hospital/repo"
	"hospital/utils"
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

	err := repo.CreateDoctor(&doctor)
	if err != nil {
		return nil, err
	}

	return &doctor, nil
}

// GetDoctorByID retrieves a doctor by ID from the database
func (ds *DoctorService) GetDoctorByID(doctorID string) (*models.Doctor, error) {
	return repo.GetDoctorByID(doctorID)
}

// UpdateDoctorContact updates a doctor's contact number
func (ds *DoctorService) UpdateDoctorContact(doctorID string, payload models.UpdateDoctorPayload) (*models.Doctor, error) {
	// First, check if the doctor exists
	_, err := repo.GetDoctorByID(doctorID)
	if err != nil {
		return nil, err
	}

	// Perform the update
	updates := map[string]interface{}{
		"contact_no": payload.ContactNo,
		"updated_at": time.Now().Unix(),
	}

	err = repo.UpdateDoctor(doctorID, updates)
	if err != nil {
		return nil, err
	}

	// Return the updated doctor
	return repo.GetDoctorByID(doctorID)
}
