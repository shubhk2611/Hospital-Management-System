package handlers

import (
	"net/http"

	"hospital/models"
	"hospital/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DoctorHandler struct {
	doctorService *services.DoctorService
}

func NewDoctorHandler() *DoctorHandler {
	return &DoctorHandler{
		doctorService: services.NewDoctorService(),
	}
}

// CreateDoctor handles POST requests to create a new doctor
func (dh *DoctorHandler) CreateDoctor(c *gin.Context) {
	var payload models.DoctorPayload

	// Bind and validate the JSON request body
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor, err := dh.doctorService.CreateDoctor(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create doctor in database"})
		return
	}

	c.JSON(http.StatusCreated, doctor)
}

// GetDoctorByID handles GET requests to retrieve a doctor by ID
func (dh *DoctorHandler) GetDoctorByID(c *gin.Context) {
	doctorID := c.Param("id")

	doctor, err := dh.doctorService.GetDoctorByID(doctorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve doctor from database"})
		}
		return
	}

	c.JSON(http.StatusOK, doctor)
}

// UpdateDoctorContact handles PATCH requests to update a doctor's contact number
func (dh *DoctorHandler) UpdateDoctorContact(c *gin.Context) {
	doctorID := c.Param("id")

	var payload models.UpdateDoctorPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor, err := dh.doctorService.UpdateDoctorContact(doctorID, payload)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update doctor in database"})
		}
		return
	}

	c.JSON(http.StatusOK, doctor)
}
