package handlers

import (
	"net/http"

	"hospital/models"
	"hospital/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PatientHandler struct {
	patientService *services.PatientService
}

func NewPatientHandler() *PatientHandler {
	return &PatientHandler{
		patientService: services.NewPatientService(),
	}
}

// CreatePatient handles POST requests to create a new patient
func (ph *PatientHandler) CreatePatient(c *gin.Context) {
	var payload models.PatientPayload

	// Bind and validate the JSON request body
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient, err := ph.patientService.CreatePatient(payload)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor ID not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient in database"})
		}
		return
	}

	c.JSON(http.StatusCreated, patient)
}

// GetPatientByID handles GET requests to retrieve a patient by ID
func (ph *PatientHandler) GetPatientByID(c *gin.Context) {
	patientID := c.Param("id")

	patient, err := ph.patientService.GetPatientByID(patientID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patient from database"})
		}
		return
	}

	c.JSON(http.StatusOK, patient)
}

// UpdatePatient handles PATCH requests to update a patient's details
func (ph *PatientHandler) UpdatePatient(c *gin.Context) {
	patientID := c.Param("id")

	var payload models.UpdatePatientPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient, err := ph.patientService.UpdatePatient(patientID, payload)
	if err != nil {
		if err == gorm.ErrInvalidField {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
		} else if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found or no changes were made"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient in database"})
		}
		return
	}

	c.JSON(http.StatusOK, patient)
}

// GetPatientsByDoctorID handles GET requests to fetch all patients by doctor ID
func (ph *PatientHandler) GetPatientsByDoctorID(c *gin.Context) {
	doctorID := c.Param("doctorId")

	patients, err := ph.patientService.GetPatientsByDoctorID(doctorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patients from database"})
		}
		return
	}

	c.JSON(http.StatusOK, patients)
}
