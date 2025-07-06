package main

import (
	// "time"
	"context"
	"fmt"
	"os"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// type enrollmentFromView struct {
	// ID    int
	// Name  string
	// Price float64
// }

// type enrollmentByEmployeeID struct {
	// ID        uuid.UUID
	// CourseID  uuid.UUID
	// CreatedAt time.Time
	// Status    string
// }

type enrollmentToCreate struct {
	CourseID   uuid.UUID `json:"course_id"`
	EmployeeID uuid.UUID `json:"employee_id"`
}

func main() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)

		os.Exit(1)
	}

	defer dbpool.Close()

	router := gin.Default()

	router.POST("/enrollments", func(c *gin.Context) {
		var enrollmentJSON enrollmentToCreate

		if err := c.ShouldBindJSON(&enrollmentJSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		
			return
		}

		_, err := dbpool.Exec(context.Background(), "CALL enrollments.enroll(@course_id, @employee_id)", pgx.NamedArgs{"course_id": enrollmentJSON.CourseID, "employee_id": enrollmentJSON.EmployeeID,})
		if err != nil {
			fmt.Println(err)
			// os.Exit(1)
			// return fmt.Errorf("unable to create enrollment: %w", err)
		}

		if _, err := dbpool.Exec(context.Background(), "SELECT id FROM enrollments.enrollments WHERE course_id = @course_id AND employee_id = @employee_id", pgx.NamedArgs{"course_id": enrollmentJSON.CourseID, "employee_id": enrollmentJSON.EmployeeID,}); err != nil {
			os.Exit(1)
		}

		c.JSON(http.StatusOK, gin.H{"status": "enrollment created"})
	})
	router.Run(":8080")
}
