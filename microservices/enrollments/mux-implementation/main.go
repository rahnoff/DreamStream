package main

import (
	"database/sql"
    // "fmt"
    "log"
	"time"
	"os"

    "net/http"
    // strconv"
    "encoding/json"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
	"github.com/google/uuid"
)


type App struct {
	Router *mux.Router
	DB *sql.DB
}


func (a *App) Initialize(user string, password string, databaseName string) {
	// connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, databaseName)
    var err error
    // a.DB, err = sql.Open("postgres", connectionString)
	a.DB, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/dream_stream?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    a.Router = mux.NewRouter()
    a.initializeRoutes()
}


func (a *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a.Router))
}


type enrollmentView struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}


type enrollment struct {
	ID uuid.UUID
	CourseID uuid.UUID
	CreatedAt time.Time
	Status string
}


func (e *enrollment) getEnrollment(db *sql.DB, employeeID uuid.UUID) error {
    // return db.QueryRow(context.Background(), "SELECT id, course_id, created_at, status FROM enrollments.enrollments WHERE employee_id = $1", employeeID).Scan(&e.ID, &e.CourseID, &e.CreatedAt, &e.Status)
	return db.QueryRow("SELECT id, course_id, created_at, status FROM enrollments.enrollments WHERE employee_id = $1", employeeID).Scan(&e.ID, &e.CourseID, &e.CreatedAt, &e.Status)

}

// func (p *product) updateProduct(db *sql.DB) error {
//     _, err :=
//         db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
//             p.Name, p.Price, p.ID)
//     return err
// }

// func (p *product) deleteProduct(db *sql.DB) error {
//     _, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)
//     return err
// }

// func (p *product) createProduct(db *sql.DB) error {
//     err := db.QueryRow(
//         "INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
//         p.Name, p.Price).Scan(&p.ID)
//     if err != nil {
//         return err
//     }
//     return nil
// }

// func getProducts(db *sql.DB, start, count int) ([]product, error) {
//     rows, err := db.Query(
//         "SELECT id, name,  price FROM products LIMIT $1 OFFSET $2",
//         count, start)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()
//     products := []product{}
//     for rows.Next() {
//         var p product
//         if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
//             return nil, err
//         }
//         products = append(products, p)
//     }
//     return products, nil
// }


func (a *App) getEnrollment(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    // employeeID, err := strconv.Atoi(vars["employee_id"])
	employeeID := vars["employee_id"]
	employeeIDUUID, err := uuid.Parse(employeeID)
    if err != nil {
        return
    }
	// employeeIDUUID := uuid.MustParse(fmt.Sprintf("%d", vars["employee_id"]))
    // if err != nil {
    //     respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
    //     return
    // }
    e := enrollment{}
    if err := e.getEnrollment(a.DB, employeeIDUUID); err != nil {
        switch err {
        case sql.ErrNoRows:
            respondWithError(w, http.StatusNotFound, "Enrollment not found")
        default:
            respondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }
    respondWithJSON(w, http.StatusOK, e)
}


func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}


func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}


func (a *App) initializeRoutes() {
    // a.Router.HandleFunc("/enrollments", a.getProducts).Methods("GET")
    // a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
    // a.Router.HandleFunc("/enrollments/{employee_id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$}", a.getEnrollment).Methods("GET")
    a.Router.HandleFunc("/enrollments/{employee_id}", a.getEnrollment).Methods("GET")
	// a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
    // a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
}


func main() {
	a := App{}
	a.Initialize(os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))
	a.Run(":8000")
}
