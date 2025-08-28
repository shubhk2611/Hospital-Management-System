package router

import (
	"hospital/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the Gin router with all API routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Initialize handlers
	doctorHandler := handlers.NewDoctorHandler()
	patientHandler := handlers.NewPatientHandler()

	// Doctor API Endpoints
	router.POST("/doctor", doctorHandler.CreateDoctor)
	router.GET("/doctor/:id", doctorHandler.GetDoctorByID)
	router.PATCH("/doctor/:id", doctorHandler.UpdateDoctorContact)

	// Patient API Endpoints
	router.POST("/patient", patientHandler.CreatePatient)
	router.GET("/patient/:id", patientHandler.GetPatientByID)
	router.PATCH("/patient/:id", patientHandler.UpdatePatient)
	router.GET("/fetchPatientByDoctorId/:doctorId", patientHandler.GetPatientsByDoctorID)

	return router
}
