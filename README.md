Hospital Management System API

This is a RESTful API for a hospital management system built with Go. The API provides endpoints to manage doctors and patients, with all data being persisted in a MySQL database.

Tech Stack :

Go (Golang): The core programming language.
Gin Gonic: A high-performance HTTP web framework for routing.
GORM: An ORM (Object-Relational Mapper) for interacting with the MySQL database.
MySQL: The relational database used for data storage.


Features:

Doctor Management: Endpoints to create, retrieve, and update doctor records.
Patient Management: Endpoints to create, retrieve, and update patient records.
Doctor-Patient Association: An endpoint to retrieve a list of all patients assigned to a specific doctor.
Modular Architecture: The project is structured into separate packages (doctor, patient, database) for better organization and scalability.


Setup Instructions :

Prerequisites
Go (version 1.18 or higher)
A running MySQL database instance
A tool for API testing (e.g., Postman, curl)

1. Project Initialization
Clone this repository or create a new project directory.
Initialize the Go module in the root directory:
go mod init hospital-management-modular

Install the required dependencies:
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql

Run go mod tidy to clean up and synchronize dependencies.
go mod tidy

2. Database Configuration :

Ensure your MySQL server is running.
Create a database named hospital. You can do this from the MySQL command line client:

CREATE DATABASE hospital;

Update the database credentials in the database/db.go file. Find the dsn variable and replace the placeholder password with your own:

// In database/db.go
dsn := "root:your-password-here@tcp(127.0.0.1:3306)/hospital?charset=utf8mb4&parseTime=True&loc=Local"

3. Running the Application
From the project's root directory, run the application:

go run main.go

The server will start and listen on http://localhost:8080. It will also automatically create the doctors and patients tables in your database.

API Endpoints
All API endpoints return JSON responses.

Doctor Endpoints
POST /doctor

Description: Creates a new doctor.

Request Body:

{
  "name": "Dr. Smith",
  "contact_no": "1234567890"
}

Response: 201 Created with the new doctor record.

GET /doctor/:id

Description: Retrieves a doctor by their unique ID.

Response: 200 OK with the doctor record, or 404 Not Found if the ID does not exist.

PATCH /doctor/:id

Description: Updates a doctor's contact number.

Request Body:

{
  "contact_no": "9876543210"
}

Response: 200 OK with the updated record, or 404 Not Found.

Patient Endpoints
POST /patient

Description: Creates a new patient. The assigned doctor_id must be for an existing doctor.

Request Body:

{
  "name": "John Doe",
  "contact_no": "1112223333",
  "address": "123 Main St",
  "doctor_id": "your-doctor-id"
}

Response: 201 Created with the new patient record.

GET /patient/:id

Description: Retrieves a patient by their unique ID.

Response: 200 OK with the patient record, or 404 Not Found.

PATCH /patient/:id

Description: Updates one or more patient fields (contact_no, address, doctor_id).

Request Body (example):

{
  "address": "456 Oak St",
  "doctor_id": "another-doctor-id"
}

Response: 200 OK with the updated record, or 404 Not Found.

GET /fetchPatientByDoctorId/:doctorId

Description: Retrieves a list of all patients associated with a specific doctor.

Response: 200 OK with a JSON array of patient records, or 404 Not Found if the doctor does not exist.