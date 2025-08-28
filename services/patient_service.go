package services

import (
	"time"

	"hospital/database"
	"hospital/models"
	"hospital/utils"

	"gorm.io/gorm"
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
	_, err := ps.doctorService.GetDoctorByID(payload.DoctorID)
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

	result := database.GetDB().Create(&patient)
	if result.Error != nil {
		return nil, result.Error
	}

	return &patient, nil
}

// GetPatientByID retrieves a patient by ID from the database
func (ps *PatientService) GetPatientByID(patientID string) (*models.Patient, error) {
	var patient models.Patient
	result := database.GetDB().First(&patient, "ID = ?", patientID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &patient, nil
}

// UpdatePatient updates a patient's details
func (ps *PatientService) UpdatePatient(patientID string, payload models.UpdatePatientPayload) (*models.Patient, error) {
	var patient models.Patient

	// Check if the patient exists
	result := database.GetDB().First(&patient, "ID = ?", patientID)
	if result.Error != nil {
		return nil, result.Error
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
	result = database.GetDB().Model(&patient).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// Re-fetch the updated record
	database.GetDB().First(&patient, "ID = ?", patientID)
	return &patient, nil
}

// GetPatientsByDoctorID retrieves all patients associated with a doctor
func (ps *PatientService) GetPatientsByDoctorID(doctorID string) ([]models.Patient, error) {
	// First, verify that the doctor exists
	_, err := ps.doctorService.GetDoctorByID(doctorID)
	if err != nil {
		return nil, err
	}

	var patients []models.Patient
	result := database.GetDB().Where("doctor_id = ?", doctorID).Find(&patients)
	if result.Error != nil {
		return nil, result.Error
	}

	return patients, nil
}
