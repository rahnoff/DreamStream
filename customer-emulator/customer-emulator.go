package main

import (
    "fmt"
    "sync"
    // "time"
	"os"
	"encoding/csv"
	"math/rand"
	"io"
	"log"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
)

// Define a struct to model the data
type User struct {
	Name string `json:"course_id"`
	Job  string `json:"employee_id"`
}

func readIDs(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		if err == io.EOF {
			return []string{}, nil
		}
		return nil, err
	}
	var firstColumn []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) > 0 {
			firstColumn = append(firstColumn, record[0])
		}
	}
	return firstColumn, nil
}

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
	coursesIDs, err := readIDs("courses.csv")
	if err != nil {
		log.Fatal(err)
	}
	employeesIDs, err := readIDs("employees.csv")
	if err != nil {
		log.Fatal(err)
	}
    for {
        // fmt.Printf("Worker %d is working...\n", id)
	// rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(coursesIDs))
	randomIndex_1 := rand.Intn(len(employeesIDs))
	chosen := coursesIDs[randomIndex]
	chosen_1 := employeesIDs[randomIndex_1]
	// fmt.Println(chosen, chosen_1)
	user := User{
		Name: chosen,
		Job:  chosen_1,
	}
	fmt.Println(user)

	// Marshal the struct into JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// Create a new HTTP request
	url := "http://127.0.0.1:2000/enrollments"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read and print the response
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
    }
}

func main() {
    var wg sync.WaitGroup

    // Start 3 workers
    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    // This won't wait forever since workers never finish
    // So we block differently
    select {} // Keep main alive
}