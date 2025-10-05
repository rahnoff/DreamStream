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
	// a.Router.HandleFunc("/enrollments/{id}", a.getEnrollments).Methods("GET")
	// a.Router.HandleFunc("/enrollments", a.createProduct).Methods("POST")
	// a.Router.HandleFunc("/enrollments/{id}", a.updateProduct).Methods("PUT")
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
			// log.Fatal(enrollmentsScanError.Error())
			return
		}
		enrollmentsOuts = append(enrollmentsOuts, enrollmentOut)
	}
	respondWithJSON(responseWriter, http.StatusOK, enrollmentsOuts)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError writes error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
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
