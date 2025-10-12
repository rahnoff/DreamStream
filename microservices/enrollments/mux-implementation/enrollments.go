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

type EnrollmentByEmployeeIDOut struct {
	EnrollmentID     uuid.UUID `json:"enrollment_id"`
	CourseID         uuid.UUID `json:"course_id"`
	CreatedAt        time.Time `json:"created_at"`
	EnrollmentStatus string    `json:"enrollment_status"`
}

type EnrollmentsMicroservice struct {
	PostgresqlPool *pgxpool.Pool
	Router         *mux.Router
}

func (em *EnrollmentsMicroservice) Initialize() {
	postgresqlPool, postgresqlPoolCreationError := pgxpool.New(context.Background(), os.Getenv("POSTGRESQL_URL"))
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
	vars := mux.Vars(request)
	enrollmentsByEmployeeID, enrollmentsByEmployeeIDSelectError := em.PostgresqlPool.Query(context.Background(), "SELECT id, course_id, created_at, status FROM enrollments.enrollments WHERE employee_id = $1;", vars["employee_id"])
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

func respondWithJSON(responseWriter http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(code)
	responseWriter.Write(response)
}

func respondWithError(responseWriter http.ResponseWriter, code int, message string) {
	respondWithJSON(responseWriter, code, map[string]string{"error": message})
}

func (em *EnrollmentsMicroservice) Run(address string) {
	log.Printf("Server started on %s", address)
	log.Fatal(http.ListenAndServe(address, em.Router))
}

func main() {
	enrollmentsMicroservice := EnrollmentsMicroservice{}
	postgresqlURL := os.Getenv("POSTGRESQL_URL")
	if postgresqlURL == "" {
		postgresqlURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}
	enrollmentsMicroservice.Initialize()
	enrollmentsMicroservice.Run(":" + os.Getenv("ENROLLMENTS_PORT"))
}
