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

type CategoryOutGetRequest struct {
	CategoryID       int8 `json:"id"`
	CategoryName     string    `json:"name"`
}

type EnrollmentInPostRequest struct {
	CourseID   int16 `json:"course_id"`
	EmployeeID uuid.UUID `json:"employee_id"`
}

type EnrollmentOutGeneric struct {
	EnrollmentID int64 `json:"id"`
}

type CourseByCategoryIDOutGetRequest struct {
	CourseID     int64 `json:"id"`
	CourseDescription string    `json:"description"`
	CourseID         int16 `json:"course_id"`
	CreatedAt        time.Time `json:"created_at"`
	EnrollmentStatus string    `json:"status"`
}

type EnrollmentInPutRequest struct {
	EnrollmentStatus string `json:"status"`
}

type CoursesMicroservice struct {
	PostgresqlPool *pgxpool.Pool
	Router         *mux.Router
}

func (cm *CoursesMicroservice) Initialize() {
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
	cm.PostgresqlPool = postgresqlPool
	cm.Router = mux.NewRouter()
	cm.initializeRoutes()
}

func (cm *CoursesMicroservice) initializeRoutes() {
	cm.Router.HandleFunc("/categories", cm.getCategories).Methods("GET")
	cm.Router.HandleFunc("/courses/{category_id}", cm.getCoursesByCategoryID).Methods("GET")
	cm.Router.HandleFunc("/quizes/course_id", cm.getQuizesByCourseID).Methods("GET")
	cm.Router.HandleFunc("/questions/{quiz_id}", cm.getQuestionsByQuizID).Methods("GET")
	cm.Router.HandleFunc("/answers/{question_id}", cm.getAnswersByQuestionID).Methods("GET")
}

func (cm *CoursesMicroservice) getCategories(responseWriter http.ResponseWriter, request *http.Request) {
	categories, categoriesSelectError := cm.PostgresqlPool.Query(context.Background(), "SELECT id, name FROM courses.categories_m_v;")
	if (categoriesSelectError != nil) {
		respondWithError(responseWriter, http.StatusInternalServerError, "Unable to retrieve categories")
		return
	}
	defer categories.Close()
	var categoriesOutsGetRequest []CategoryOutGetRequest
	for (categories.Next()) {
		var categoryOutGetRequest CategoryOutGetRequest
		categoriesScanError := categories.Scan(&categoryOutGetRequest.CategoryID, &categoryOutGetRequest.CategoryName)
		if (categoriesScanError != nil) {
			respondWithError(responseWriter, http.StatusInternalServerError, "Unable to scan categories")
			return
		}
		categoriesOutsGetRequest = append(categoriesOutsGetRequest, categoryOutGetRequest)
	}
	respondWithJSON(responseWriter, http.StatusOK, categoriesOutsGetRequest)
}

func (cm *CoursesMicroservice) getCoursesByCategoryID(responseWriter http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	coursesByCategoryID, coursesByCategoryIDSelectError := cm.PostgresqlPool.Query(context.Background(), "SELECT id, description, filename, language, length, name FROM courses.courses WHERE category_id = $1;", variables["category_id"])
	if (coursesByCategoryIDSelectError != nil) {
		respondWithError(responseWriter, http.StatusInternalServerError, "Unable to retrieve courses by a category ID")
		return
	}
	defer coursesByCategoryID.Close()
	var coursesByCategoryIDOutsGetRequest []EnrollmentByEmployeeIDOutGetRequest
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
