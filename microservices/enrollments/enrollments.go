package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EnrollmentOutGetRequest struct {
	EmployeeID       uuid.UUID `json:"employee_id"`
	EmployeeName     string    `json:"employee_name"`
	CourseName       string    `json:"course_name"`
	EnrolledAt       time.Time `json:"enrolled_at"`
	EnrollmentStatus string    `json:"status"`
}

type EnrollmentInPostRequest struct {
	CourseID   int16 `json:"course_id"`
	EmployeeID uuid.UUID `json:"employee_id"`
}

type EnrollmentOutGeneric struct {
	EnrollmentID int64 `json:"id"`
}

type EnrollmentByEmployeeIDOutGetRequest struct {
	EnrollmentID     int64 `json:"id"`
	CourseID         int16 `json:"course_id"`
	CreatedAt        time.Time `json:"created_at"`
	EnrollmentStatus string    `json:"status"`
}

type EnrollmentInPutRequest struct {
	EnrollmentStatus string `json:"status"`
}

type EnrollmentsMicroservice struct {
	PostgresqlPool *pgxpool.Pool
	Router         *mux.Router
}

func (em *EnrollmentsMicroservice) Initialize() {
	postgresqlURL := os.Getenv("POSTGRESQL_URL")
	if (postgresqlURL == "") {
		postgresqlURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}
	postgresqlPool, postgresqlPoolCreationError := pgxpool.New(context.Background(), postgresqlURL)
	if (postgresqlPoolCreationError != nil) {
		log.Fatal("Unable to create a PostgreSQL connection pool:\n", postgresqlPoolCreationError.Error())
	}
	postgresqlPoolPingError := postgresqlPool.Ping(context.Background())
	if (postgresqlPoolPingError != nil) {
		log.Fatal("Unable to ping a PostgreSQL server:\n", postgresqlPoolPingError.Error())
	}
	log.Println("Connected to a PostgreSQL server")
	em.PostgresqlPool = postgresqlPool
	em.Router = mux.NewRouter()
	em.initializeRoutes()
}

func (em *EnrollmentsMicroservice) initializeRoutes() {
	em.Router.HandleFunc("/enrollments", em.getEnrollments).Methods("GET")
	em.Router.HandleFunc("/enrollments/{employee_id}", em.getEnrollmentsByEmployeeID).Methods("GET")
	em.Router.HandleFunc("/enrollments", em.createEnrollment).Methods("POST")
	em.Router.HandleFunc("/enrollments/{id}", em.updateEnrollment).Methods("PUT")
	// a.Router.HandleFunc("/enrollments/{id}", a.deleteProduct).Methods("DELETE")
}

func (em *EnrollmentsMicroservice) getEnrollments(responseWriter http.ResponseWriter, request *http.Request) {
	enrollments, enrollmentsSelectError := em.PostgresqlPool.Query(context.Background(), "SELECT employee_id, employee_name, course_name, enrolled_at, enrollment_status FROM enrollments.enrollments_m_v;")
	if (enrollmentsSelectError != nil) {
		respondWithError(responseWriter, http.StatusInternalServerError, "Unable to retrieve enrollments")
		return
	}
	defer enrollments.Close()
	var enrollmentsOutsGetRequest []EnrollmentOutGetRequest
	for (enrollments.Next()) {
		var enrollmentOutGetRequest EnrollmentOutGetRequest
		enrollmentsScanError := enrollments.Scan(&enrollmentOutGetRequest.EmployeeID, &enrollmentOutGetRequest.EmployeeName, &enrollmentOutGetRequest.CourseName, &enrollmentOutGetRequest.EnrolledAt, &enrollmentOutGetRequest.EnrollmentStatus)
		if (enrollmentsScanError != nil) {
			respondWithError(responseWriter, http.StatusInternalServerError, "Unable to scan enrollments")
			return
		}
		enrollmentsOutsGetRequest = append(enrollmentsOutsGetRequest, enrollmentOutGetRequest)
	}
	respondWithJSON(responseWriter, http.StatusOK, enrollmentsOutsGetRequest)
}

func (em *EnrollmentsMicroservice) getEnrollmentsByEmployeeID(responseWriter http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	enrollmentsByEmployeeID, enrollmentsByEmployeeIDSelectError := em.PostgresqlPool.Query(context.Background(), "SELECT id, course_id, created_at, status FROM enrollments.enrollments WHERE employee_id = $1;", variables["employee_id"])
	if (enrollmentsByEmployeeIDSelectError != nil) {
		respondWithError(responseWriter, http.StatusInternalServerError, "Unable to retrieve enrollments by an employee ID")
		return
	}
	defer enrollmentsByEmployeeID.Close()
	var enrollmentsByEmployeeIDOutsGetRequest []EnrollmentByEmployeeIDOutGetRequest
	for (enrollmentsByEmployeeID.Next()) {
		var EnrollmentByEmployeeIDOutGetRequest EnrollmentByEmployeeIDOutGetRequest
		enrollmentsByEmployeeIDScanError := enrollmentsByEmployeeID.Scan(&EnrollmentByEmployeeIDOutGetRequest.EnrollmentID, &EnrollmentByEmployeeIDOutGetRequest.CourseID, &EnrollmentByEmployeeIDOutGetRequest.CreatedAt, &EnrollmentByEmployeeIDOutGetRequest.EnrollmentStatus)
		if (enrollmentsByEmployeeIDScanError != nil) {
			respondWithError(responseWriter, http.StatusInternalServerError, "Unable to scan enrollments by an employee ID")
			return
		}
		enrollmentsByEmployeeIDOutsGetRequest = append(enrollmentsByEmployeeIDOutsGetRequest, EnrollmentByEmployeeIDOutGetRequest)
	}
	respondWithJSON(responseWriter, http.StatusOK, enrollmentsByEmployeeIDOutsGetRequest)
}

func (em *EnrollmentsMicroservice) createEnrollment(responseWriter http.ResponseWriter, request *http.Request) {
	var (
		enrollmentInPostRequest EnrollmentInPostRequest
		enrollmentOutGeneric    EnrollmentOutGeneric
	)
	decoder := json.NewDecoder(request.Body)
	decodeRequestBodyToEnrollmentInPostRequestError := decoder.Decode(&enrollmentInPostRequest)
	if (decodeRequestBodyToEnrollmentInPostRequestError != nil) {
		log.Println("Unable to decode a request body into enrollmentInPostRequest:\n", decodeRequestBodyToEnrollmentInPostRequestError.Error())
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()
	enrollError := em.PostgresqlPool.QueryRow(context.Background(), "CALL enrollments.enroll($1, $2, NULL)", enrollmentInPostRequest.CourseID, enrollmentInPostRequest.EmployeeID).Scan(&enrollmentOutGeneric.EnrollmentID)
	if (enrollError != nil) {
		log.Println("Unable to enroll:\n", enrollError.Error())
		respondWithError(responseWriter, http.StatusInternalServerError, "Unable to create an enrollment")
		return
	}
	respondWithJSON(responseWriter, http.StatusCreated, enrollmentOutGeneric)
}

func (em *EnrollmentsMicroservice) updateEnrollment(responseWriter http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	var (
		enrollmentInPutRequest EnrollmentInPutRequest
		enrollmentOutGeneric   EnrollmentOutGeneric
	)
	decoder := json.NewDecoder(request.Body)
	decodeRequestBodyToEnrollmentInPutRequestError := decoder.Decode(&enrollmentInPutRequest)
	if (decodeRequestBodyToEnrollmentInPutRequestError != nil) {
		log.Println("Unable to decode a request body into enrollmentInPutRequest:\n", decodeRequestBodyToEnrollmentInPutRequestError.Error())
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()
	updateEnrollmentError := em.PostgresqlPool.QueryRow(context.Background(), "CALL enrollments.update_enrollment_status($1, $2, NULL)", variables["id"], enrollmentInPutRequest.EnrollmentStatus).Scan(&enrollmentOutGeneric.EnrollmentID)
	if (updateEnrollmentError != nil) {
		log.Println("Unable to update an enrollment status:\n", updateEnrollmentError.Error())
		respondWithError(responseWriter, http.StatusInternalServerError, "Unable to update an enrollment")
		return
	}
	respondWithJSON(responseWriter, http.StatusNoContent, enrollmentOutGeneric)
}

func respondWithJSON(responseWriter http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(code)
	responseWriter.Write(response)
}

func respondWithError(responseWriter http.ResponseWriter, code int, message string) {
	respondWithJSON(responseWriter, code, map[string]string{"error": message})
}

func (em *EnrollmentsMicroservice) Run() {
	enrollmentsURL := os.Getenv("ENROLLMENTS_URL")
	if (enrollmentsURL == "") {
		enrollmentsURL = "127.0.0.1:2000"
	}
	log.Printf("Server started on %s", enrollmentsURL)
	log.Fatal(http.ListenAndServe(enrollmentsURL, em.Router))
}

func main() {
	enrollmentsMicroservice := EnrollmentsMicroservice{}
	enrollmentsMicroservice.Initialize()
	enrollmentsMicroservice.Run()
}
