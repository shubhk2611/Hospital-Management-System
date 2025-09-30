package repo

import (
	"hospital/database"
	"hospital/models"
)

// Post doctor
func CreateDoctor(doctor *models.Doctor) error {
	result := database.GetDB().Create(doctor)
	return result.Error
}

// Patch doctor
func UpdateDoctor(doctorID string, updateData map[string]interface{}) error {
	result := database.GetDB().Model(&models.Doctor{}).Where("ID = ?", doctorID).Updates(updateData)
	return result.Error
}

// Get doctor by ID
func GetDoctorByID(doctorID string) (*models.Doctor, error) {
	var doctor models.Doctor
	result := database.GetDB().First(&doctor, "ID = ?", doctorID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &doctor, nil
}
