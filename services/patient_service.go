package services

import (
	"gorm.io/gorm"
	"hospital/models"
	"hospital/repo"
	"hospital/utils"
	"time"
)

type PatientService struct {
	doctorService *DoctorService
}

func NewPatientService() *PatientService {
	return &PatientService{
		doctorService: NewDoctorService(),
	}
}

// CreatePatient creates a new patient in the database
func (ps *PatientService) CreatePatient(payload models.PatientPayload) (*models.Patient, error) {
	// Verify that the assigned doctor exists
	_, err := repo.GetDoctorByID(payload.DoctorID)
	if err != nil {
		return nil, err
	}

	patient := models.Patient{
		ID:        utils.GenerateID(),
		Name:      payload.Name,
		ContactNo: payload.ContactNo,
		Address:   payload.Address,
		DoctorID:  payload.DoctorID,
	}

	err = repo.CreatePatient(&patient)
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

// GetPatientByID retrieves a patient by ID from the database
func (ps *PatientService) GetPatientByID(patientID string) (*models.Patient, error) {
	return repo.GetPatientByID(patientID)
}

// UpdatePatient updates a patient's details
func (ps *PatientService) UpdatePatient(patientID string, payload models.UpdatePatientPayload) (*models.Patient, error) {
	// Check if the patient exists
	_, err := repo.GetPatientByID(patientID)
	if err != nil {
		return nil, err
	}

	// Prepare updates map
	updates := make(map[string]interface{})
	if payload.ContactNo != nil {
		updates["contact_no"] = *payload.ContactNo
	}
	if payload.Address != nil {
		updates["address"] = *payload.Address
	}
	if payload.DoctorID != nil {
		updates["doctor_id"] = *payload.DoctorID
	}

	// Return error if no fields to update
	if len(updates) == 0 {
		return nil, gorm.ErrInvalidField
	}

	// Add timestamp
	updates["updated_at"] = time.Now().Unix()

	// Perform the update
	err = repo.UpdatePatient(patientID, updates)
	if err != nil {
		return nil, err
	}

	// Return the updated patient
	return repo.GetPatientByID(patientID)
}

// GetPatientsByDoctorID retrieves all patients associated with a doctor
func (ps *PatientService) GetPatientsByDoctorID(doctorID string) ([]models.Patient, error) {
	// First, verify that the doctor exists
	_, err := repo.GetDoctorByID(doctorID)
	if err != nil {
		return nil, err
	}

	return repo.GetPatientsByDoctorID(doctorID)
}
