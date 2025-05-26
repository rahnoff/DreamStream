package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    filename := "output.csv"
    writer, file, err := createCSVWriter(filename)
    if err != nil {
        fmt.Println("Error creating CSV writer:", err)
        return
    }
    defer file.Close()
    header := []string{"id", "title", "authors"}
    writeCSVRecord(writer, header)
    records := [][]string{
        {"John", "30", "New York"},
        {"Alice", "25", "Tokyo"},
        {"Bob", "35", "London"},
    }
    for _, record := range records {
        writeCSVRecord(writer, record)
    }
    // Flush the writer and check for any errors
    writer.Flush()
    if err := writer.Error(); err != nil {
        fmt.Println("Error flushing CSV writer:", err)
    }
}

func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {
    f, err := os.Create(filename)
    if err != nil {
        return nil, nil, err
    }
    writer := csv.NewWriter(f)
    return writer, f, nil
}

func writeCSVRecord(writer *csv.Writer, record []string) {
    err := writer.Write(record)
    if err != nil {
        fmt.Println("Error writing record to CSV:", err)
    }
}