package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
	"os"
	"encoding/csv"
	"bytes"
	// "github.com/google/uuid"
)

type Category struct {
	CategoryID   int8   `json:"id"`
	CategoryName string `json:"name"`
}

type Course struct {
	CourseID          int64  `json:"id"`
	CourseDescription string `json:"description"`
	CourseFilename    string `json:"filename"`
	CourseLanguage    string `json:"language"`
	CourseLength      string `json:"length"`
	CourseName        string `json:"name"`
}

type Enrollment struct {
	CourseID int64 `json:"course_id"`
	EmployeeID string `json:"employee_id"`
}

type EnrollmentCreated struct {
	EnrollmentCreatedID int64 `json:"id"`
}

type EnrollmentUpdated struct {
	EnrollmentCreatedID int64 `json:"id"`
}

func main() {
	file, err := os.Open("/var/tmp/employees.csv")
    if err != nil {
        log.Fatal("Error opening file:", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		log.Fatal("Error reading header:", err)
	}
	var firstColumn []string

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal("Error reading CSV:", err)
        }
		if len(record) > 0 {
            firstColumn = append(firstColumn, record[0])
        }
    }

	for {
		rand.Seed(time.Now().UnixNano())
	    employeeIndex := rand.Intn(len(firstColumn))
	    selectedEmployeeID := firstColumn[employeeIndex]
		// fmt.Println(selectedEmployeeID)
		resp, categoriesReadError := http.Get("http://127.0.0.1:2002/categories")
		if categoriesReadError != nil {
			log.Fatal("HTTP request failed:", categoriesReadError)
		}
		time.Sleep(5 * time.Second)
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Failed to read response body:", err)
		}
		var categories []Category
		err = json.Unmarshal(body, &categories)
		if err != nil {
			log.Fatal("JSON unmarshal failed:", err)
		}
		var categoriesIDs []int8
		for _, category := range categories {
			categoriesIDs = append(categoriesIDs, category.CategoryID)
		}
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(categoriesIDs))
		selectedCategoryID := categoriesIDs[index]
		// fmt.Println(selectedCategoryID)
		coursesURL := fmt.Sprintf("http://127.0.0.1:2002/courses/%d", selectedCategoryID)
		// fmt.Println(coursesURL)
		resp_1, coursesReadError := http.Get(coursesURL)
		if coursesReadError != nil {
			log.Fatal("HTTP request failed:", coursesReadError)
		}
		defer resp_1.Body.Close()
		// fmt.Println(resp_1)
		body_1, err := io.ReadAll(resp_1.Body)
		// fmt.Println(body_1)
		if err != nil {
			log.Fatal("Failed to read response body:", err)
		}
		var courses []Course
		err = json.Unmarshal(body_1, &courses)
		if err != nil {
			log.Fatal("JSON unmarshal failed:", err)
		}
		var coursesIDs []int64
		for _, course := range courses {
			coursesIDs = append(coursesIDs, course.CourseID)
		}
		// fmt.Println(coursesIDs)
		rand.Seed(time.Now().UnixNano())
		index_1 := rand.Intn(len(coursesIDs))
		selectedCourseID := coursesIDs[index_1]
		// fmt.Println(selectedCourseID)

		enrollment := Enrollment{
			CourseID: selectedCourseID,
			EmployeeID: selectedEmployeeID,
		}
		// fmt.Println(enrollment)
	
		// Marshal struct to JSON
		jsonData, err := json.Marshal(enrollment)
		if err != nil {
			log.Fatal("JSON marshaling failed:", err)
		}
		// fmt.Println(jsonData)
	
		// Create request body reader
		body_2 := bytes.NewBuffer(jsonData)
		resp_2, err := http.Post("http://127.0.0.1:2003/enrollments", "application/json", body_2)
	if err != nil {
		log.Fatal("Request failed:", err)
	}
	if resp_2.StatusCode == http.StatusCreated {
        var enrollmentCreated EnrollmentCreated
        json.NewDecoder(resp_2.Body).Decode(&enrollmentCreated)
        // fmt.Printf("Created: %+v\n", enrollmentCreated)
		enrollmentUpdated := EnrollmentUpdated{
			EnrollmentCreatedID: enrollmentCreated.EnrollmentCreatedID,
		}
		// fmt.Println(enrollment)
	
		// Marshal struct to JSON
		jsonData_1, err := json.Marshal(enrollmentUpdated)
		if err != nil {
			log.Fatal("JSON marshaling failed:", err)
		}
		// fmt.Println(jsonData)
	
		// Create request body reader
		body_3 := bytes.NewBuffer(jsonData_1)
		resp_3, err := http.Put("http://127.0.0.1:2003/enrollments", "application/json", body_3)
	if err != nil {
		log.Fatal("Request failed:", err)
	}
	if resp_2.StatusCode == http.StatusCreated {
        var enrollmentCreated EnrollmentCreated
        json.NewDecoder(resp_2.Body).Decode(&enrollmentCreated)
        fmt.Printf("Created: %+v\n", enrollmentCreated)
    } else {
        fmt.Printf("Unexpected status: %s\n", resp.Status)
    }
    } else {
        fmt.Printf("Unexpected status: %s\n", resp.Status)
    }
	// defer resp_2.Body.Close()

    // var res map[string]interface{}

    // json.NewDecoder(resp.Body).Decode(&res)

    // fmt.Println(res["json"])
	// defer resp_2.Body.Close()

	// Read and print response
	// responseBody, err := io.ReadAll(resp_2.Body)
	// if err != nil {
		// log.Fatal("Reading response failed:", err)
	// }

	// fmt.Printf("Status: %s\n", resp_2.Status)
	// fmt.Printf("Response: %s\n", responseBody)
	}
}
