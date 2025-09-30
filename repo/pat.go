package repo

import (
	"hospital/database"
	"hospital/models"
)

// Post patient
func CreatePatient(patient *models.Patient) error {
	result := database.GetDB().Create(patient)
	return result.Error
}

// Patch patient
func UpdatePatient(patientID string, updateData map[string]interface{}) error {
	result := database.GetDB().Model(&models.Patient{}).Where("ID = ?", patientID).Updates(updateData)
	return result.Error
}

// Get patient by ID
func GetPatientByID(patientID string) (*models.Patient, error) {
	var patient models.Patient
	result := database.GetDB().First(&patient, "ID = ?", patientID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &patient, nil
}

// Get all patients associated with doctor ID
func GetPatientsByDoctorID(doctorID string) ([]models.Patient, error) {
	var patients []models.Patient
	result := database.GetDB().Where("doctor_id = ?", doctorID).Find(&patients)
	if result.Error != nil {
		return nil, result.Error
	}
	return patients, nil
}