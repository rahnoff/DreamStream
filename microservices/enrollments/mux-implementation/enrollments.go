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

type EnrollmentOut struct {
	EmployeeID       uuid.UUID `json:"employee_id"`
	EmployeeName     string    `json:"employee_name"`
	CourseName       string    `json:"course_name"`
	EnrolledAt       time.Time `json:"enrolled_at"`
	EnrollmentStatus string    `json:"enrollment_status"`
}

type EnrollmentInCreatePostRequest struct {
	CourseID   uuid.UUID `json:"course_id"`
	EmployeeID uuid.UUID `json:"employee_id"`
}

type EnrollmentOutPostRequest struct {
	EnrollmentID uuid.UUID `json:"enrollment_id"`
}

type EnrollmentByEmployeeIDOut struct {
	EnrollmentID     uuid.UUID `json:"enrollment_id"`
	CourseID         uuid.UUID `json:"course_id"`
	CreatedAt        time.Time `json:"created_at"`
	EnrollmentStatus string    `json:"enrollment_status"`
}

type EnrollmentInUpdatePutRequest struct {
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
	var enrollmentsOuts []EnrollmentOut
	for (enrollments.Next()) {
		var enrollmentOut EnrollmentOut
		enrollmentsScanError := enrollments.Scan(&enrollmentOut.EmployeeID, &enrollmentOut.EmployeeName, &enrollmentOut.CourseName, &enrollmentOut.EnrolledAt, &enrollmentOut.EnrollmentStatus)
		if (enrollmentsScanError != nil) {
			respondWithError(responseWriter, http.StatusInternalServerError, "Unable to scan enrollments")
			return
		}
		enrollmentsOuts = append(enrollmentsOuts, enrollmentOut)
	}
	respondWithJSON(responseWriter, http.StatusOK, enrollmentsOuts)
}

func (em *EnrollmentsMicroservice) getEnrollmentsByEmployeeID(responseWriter http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	enrollmentsByEmployeeID, enrollmentsByEmployeeIDSelectError := em.PostgresqlPool.Query(context.Background(), "SELECT id, course_id, created_at, status FROM enrollments.enrollments WHERE employee_id = $1;", variables["employee_id"])
	if (enrollmentsByEmployeeIDSelectError != nil) {
		respondWithError(responseWriter, http.StatusInternalServerError, "Unable to retrieve enrollments by an employee ID")
		return
	}
	defer enrollmentsByEmployeeID.Close()
	var enrollmentsByEmployeeIDOuts []EnrollmentByEmployeeIDOut
	for (enrollmentsByEmployeeID.Next()) {
		var enrollmentByEmployeeIDOut EnrollmentByEmployeeIDOut
		enrollmentsByEmployeeIDScanError := enrollmentsByEmployeeID.Scan(&enrollmentByEmployeeIDOut.EnrollmentID, &enrollmentByEmployeeIDOut.CourseID, &enrollmentByEmployeeIDOut.CreatedAt, &enrollmentByEmployeeIDOut.EnrollmentStatus)
		if (enrollmentsByEmployeeIDScanError != nil) {
			respondWithError(responseWriter, http.StatusInternalServerError, "Unable to scan enrollments by an employee ID")
			return
		}
		enrollmentsByEmployeeIDOuts = append(enrollmentsByEmployeeIDOuts, enrollmentByEmployeeIDOut)
	}
	respondWithJSON(responseWriter, http.StatusOK, enrollmentsByEmployeeIDOuts)
}

func (em *EnrollmentsMicroservice) createEnrollment(responseWriter http.ResponseWriter, request *http.Request) {
	var (
		enrollmentInCreatePostRequest EnrollmentInCreatePostRequest
		enrollmentOutPostRequest      EnrollmentOutPostRequest
	)
	decoder := json.NewDecoder(request.Body)
	decodeRequestBodyToEnrollmentInCreatePostRequestError := decoder.Decode(&enrollmentInCreatePostRequest)
	if (decodeRequestBodyToEnrollmentInCreatePostRequestError != nil) {
		log.Println("Unable to decode a request body into enrollmentInCreatePostRequest:\n", decodeRequestBodyToEnrollmentInCreatePostRequestError.Error())
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()
	enrollError := em.PostgresqlPool.QueryRow(context.Background(), "CALL enrollments.enroll($1, $2, NULL)", enrollmentInCreatePostRequest.CourseID, enrollmentInCreatePostRequest.EmployeeID).Scan(&enrollmentOutPostRequest.EnrollmentID)
	if (enrollError != nil) {
		log.Println("Unable to enroll:\n", enrollError.Error())
		respondWithError(responseWriter, http.StatusInternalServerError, "Unable to create an enrollment")
		return
	}
	respondWithJSON(responseWriter, http.StatusCreated, enrollmentOutPostRequest)
}

func (em *EnrollmentsMicroservice) updateEnrollment(responseWriter http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	var (
		enrollmentInUpdatePutRequest EnrollmentInUpdatePutRequest
		enrollmentOutPostRequest      EnrollmentOutPostRequest
	)
	decoder := json.NewDecoder(request.Body)
	decodeRequestBodyToEnrollmentInUpdatePutRequestError := decoder.Decode(&enrollmentInUpdatePutRequest)
	if (decodeRequestBodyToEnrollmentInUpdatePutRequestError != nil) {
		log.Println("Unable to decode a request body into enrollmentInUpdatePutRequest:\n", decodeRequestBodyToEnrollmentInUpdatePutRequestError.Error())
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()
	updateEnrollmentError := em.PostgresqlPool.QueryRow(context.Background(), "CALL enrollments.update_enrollment_status($1, $2, NULL)", variables["id"], enrollmentInUpdatePutRequest.EnrollmentStatus).Scan(&enrollmentOutPostRequest.EnrollmentID)
	if (updateEnrollmentError != nil) {
		log.Println("Unable to update an enrollment status:\n", updateEnrollmentError.Error())
		respondWithError(responseWriter, http.StatusInternalServerError, "Unable to update an enrollment")
		return
	}
	respondWithJSON(responseWriter, http.StatusCreated, enrollmentOutPostRequest)
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
